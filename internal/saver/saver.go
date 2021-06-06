package saver

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ozoncp/ocp-prize-api/internal/flusher"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
)

type SaverResultCode int

const (
	OKSaverResultCode SaverResultCode = iota
	ErrorSaverResultCode
)

// SaveResult to write in result channel
type SaverState struct {
	ResultCode    SaverResultCode
	ErrorOfSaving error
	LostedData    uint64
}

// ISaver interface for prize Saver
type ISaver interface {
	Save(prizeToSave prize.Prize) error
	Init() error
	Close() error
	GetState() SaverState
}

// Saver for prize buffer
type Saver struct {
	capacity        int
	flusher         flusher.IFlusher
	timeout         int
	ticker          *time.Ticker
	state           SaverState
	doneChannel     chan struct{}
	bufferMutex     sync.Mutex
	buffer          []prize.Prize
	shiftOverriting int
}

// NewSaver return Saver struct with setted capacity, flusher and timeout
func NewSaver(
	capacity int,
	flusher flusher.IFlusher,
	timeout int,
) ISaver {
	return &Saver{
		capacity: capacity,
		flusher:  flusher,
		timeout:  timeout,
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
	originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
	originSaver.doneChannel = make(chan struct{})
	originSaver.state = SaverState{
		ResultCode:    OKSaverResultCode,
		ErrorOfSaving: nil,
	}
	originSaver.bufferMutex = sync.Mutex{}
	originSaver.shiftOverriting = 0
	go originSaver.savingLoop()
	return nil
}

// savingLoop for flushing data by timeout
func (originSaver *Saver) savingLoop() {
	originSaver.ticker = time.NewTicker(1000)

	fmt.Println("Start loop")
	for {
		select {
		case <-originSaver.ticker.C:

			fmt.Println("Timer ticked Len prizes", len(originSaver.buffer))
			if len(originSaver.buffer) == 0 {
				continue
			}
			originSaver.bufferMutex.Lock()
			fmt.Println("Loop mutex lock")
			leftPrizes, err := originSaver.flusher.Flush(originSaver.buffer)
			fmt.Println("Flush ok")
			if leftPrizes != nil {
				originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
				originSaver.buffer = append(originSaver.buffer, leftPrizes...)
				originSaver.state.ResultCode = ErrorSaverResultCode
				originSaver.state.ErrorOfSaving = errors.New("prizes to save left")
				fmt.Println("Prizes left")
			} else {
				originSaver.buffer = make([]prize.Prize, 0, originSaver.capacity)
				originSaver.state.ResultCode = OKSaverResultCode
				originSaver.state.ErrorOfSaving = nil
				fmt.Println("Prizes not left")
			}
			originSaver.shiftOverriting = 0
			fmt.Println("Loop mutex try unlock")
			originSaver.bufferMutex.Unlock()
			fmt.Println("Loop mutex unlock")
			if err != nil {

				fmt.Println("Error saving")
				continue
			}
		case <-originSaver.doneChannel:
			fmt.Println("CloseChannel")
			return
		}
	}
}

//Save Add prize to buffer for save by timeout
func (originSaver *Saver) Save(prizeToSave prize.Prize) error {

	fmt.Println("Save prize%i", prizeToSave.ID)
	originSaver.bufferMutex.Lock()
	fmt.Println("Save mutex lock")
	if len(originSaver.buffer) == originSaver.capacity {
		originSaver.buffer[originSaver.shiftOverriting] = prizeToSave
		originSaver.shiftOverriting++
		originSaver.state.LostedData++
	} else {
		originSaver.buffer = append(originSaver.buffer, prizeToSave)
	}
	originSaver.bufferMutex.Unlock()
	fmt.Println("Save mutex unlock")
	return nil
}

// Close Saver and stop saving by timeout
func (originSaver *Saver) Close() error {
	fmt.Println("Close func")
	originSaver.ticker.Stop()
	originSaver.doneChannel <- struct{}{}
	close(originSaver.doneChannel)
	return nil
}

// GetResult by order from result channel
func (originSaver *Saver) GetState() SaverState {
	return originSaver.state
}
