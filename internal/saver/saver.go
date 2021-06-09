package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/ozoncp/ocp-prize-api/internal/flusher"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
)

// ResultCode for saver state
type ResultCode int

const (
	//OKSaverResultCode is setted if all data saved without errors
	OKSaverResultCode ResultCode = iota
	// ErrorSaverResultCode is setted if there are errors with saving data
	ErrorSaverResultCode
)

// State to write in result channel
type State struct {
	ResultCode    ResultCode
	ErrorOfSaving error
	LostedData    uint64
	IsWorking     bool
}

// ISaver interface for prize Saver
type ISaver interface {
	Save(prizeToSave prize.Prize) error
	Init() error
	Close() error
	GetState() State
}

// Saver for prize buffer
type Saver struct {
	capacity        int
	flusher         flusher.IFlusher
	timeout         time.Duration
	ticker          *time.Ticker
	state           State
	doneChannel     chan struct{}
	bufferMutex     sync.Mutex
	buffer          []prize.Prize
	shiftOverriting int
}

// NewSaver return Saver struct with setted capacity, flusher and timeout
func NewSaver(
	capacity int,
	flusher flusher.IFlusher,
	timeout time.Duration,
) ISaver {
	return &Saver{
		capacity: capacity,
		flusher:  flusher,
		timeout:  timeout,
		state: State{
			ResultCode:    OKSaverResultCode,
			ErrorOfSaving: nil,
			IsWorking:     false,
		},
	}
}

// Init saver
func (originSaver *Saver) Init() error {
	if originSaver.capacity < 1 {
		return errors.New("incorrect capacity for saver")
	}
	if originSaver.timeout < 1 {
		return errors.New("incorrect timeout for saver")
	}
	if originSaver.state.IsWorking {
		return errors.New("saver already inited")
	}
	originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
	originSaver.doneChannel = make(chan struct{})
	originSaver.bufferMutex = sync.Mutex{}
	originSaver.shiftOverriting = 0
	go originSaver.savingLoop()
	return nil
}

// savingLoop for flushing data by timeout
func (originSaver *Saver) savingLoop() {
	originSaver.ticker = time.NewTicker(originSaver.timeout)
	originSaver.state.IsWorking = true

	for {
		select {
		case <-originSaver.ticker.C:

			if len(originSaver.buffer) == 0 {
				continue
			}
			originSaver.bufferMutex.Lock()
			leftPrizes, err := originSaver.flusher.Flush(originSaver.buffer)
			if leftPrizes != nil {
				originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
				originSaver.buffer = append(originSaver.buffer, leftPrizes...)
				originSaver.state.ResultCode = ErrorSaverResultCode
				originSaver.state.ErrorOfSaving = err
			} else {
				originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
				originSaver.state.ResultCode = OKSaverResultCode
				originSaver.state.ErrorOfSaving = nil
			}
			originSaver.shiftOverriting = 0
			originSaver.bufferMutex.Unlock()
		case <-originSaver.doneChannel:
			return
		}
	}
}

//Save Add prize to buffer for save by timeout
func (originSaver *Saver) Save(prizeToSave prize.Prize) error {
	originSaver.bufferMutex.Lock()
	if len(originSaver.buffer) == originSaver.capacity {
		if originSaver.shiftOverriting >= originSaver.capacity {
			originSaver.shiftOverriting = 0
		}
		originSaver.buffer[originSaver.shiftOverriting] = prizeToSave
		originSaver.shiftOverriting++
		originSaver.state.LostedData++
	} else {
		originSaver.buffer = append(originSaver.buffer, prizeToSave)
	}
	originSaver.bufferMutex.Unlock()
	return nil
}

// Close Saver and stop saving by timeout
func (originSaver *Saver) Close() error {
	if originSaver.state.IsWorking {
		originSaver.ticker.Stop()
		originSaver.doneChannel <- struct{}{}
		close(originSaver.doneChannel)
		originSaver.state.IsWorking = false
	} else {
		return errors.New("saver already closed")
	}
	return nil
}

// GetState by order from result channel
func (originSaver *Saver) GetState() State {
	return originSaver.state
}
