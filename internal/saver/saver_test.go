package saver_test

import (
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-prize-api/internal/mocks"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
	"github.com/ozoncp/ocp-prize-api/internal/saver"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockIFlusher
		testSaver   saver.ISaver

		prizesToAdd []prize.Prize
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockIFlusher(ctrl)
	})

	Context("Test saving prizes with correct flush", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockFlusher.EXPECT().Flush(gomock.Any()).Return(nil, nil).MinTimes(1)
		})

		It("Test with correct saving", func() {
			testSaver = saver.NewSaver(10, mockFlusher, 100000)
			err := testSaver.Init()
			Expect(err).Should(BeNil())
			for _, currentPrize := range prizesToAdd {
				errSave := testSaver.Save(currentPrize)
				Expect(errSave).Should(BeNil())
			}
			time.Sleep(400000)
			testSaver.Close()
			state := testSaver.GetState()
			Expect(state.ResultCode).Should(BeEquivalentTo(saver.OKSaverResultCode))
			Expect(state.ErrorOfSaving).Should(BeNil())
			Expect(state.LostedData).Should(BeEquivalentTo(0))
		})

		It("Test with incorrect init", func() {
			testSaver = saver.NewSaver(0, mockFlusher, 100000)
			err := testSaver.Init()
			Expect(err).ShouldNot(BeNil())
		})

		It("Test with oversaving data", func() {
			testSaver = saver.NewSaver(3, mockFlusher, 100000)
			err := testSaver.Init()
			Expect(err).Should(BeNil())
			for _, currentPrize := range prizesToAdd {
				errSave := testSaver.Save(currentPrize)
				Expect(errSave).Should(BeNil())
			}
			time.Sleep(400000)
			testSaver.Close()
			state := testSaver.GetState()
			Expect(state.ResultCode).Should(BeEquivalentTo(saver.OKSaverResultCode))
			Expect(state.ErrorOfSaving).Should(BeNil())
			Expect(state.LostedData).ShouldNot(BeEquivalentTo(0))
		})
	})
	Context("Test saving prizes with error flush", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockFlusher.EXPECT().Flush(gomock.Any()).Return(prizesToAdd, errors.New("error flushing")).MinTimes(1)
		})

		It("Test error saving", func() {
			testSaver = saver.NewSaver(10, mockFlusher, 100000)
			err := testSaver.Init()
			Expect(err).Should(BeNil())
			for _, currentPrize := range prizesToAdd {
				errSave := testSaver.Save(currentPrize)
				Expect(errSave).Should(BeNil())
			}
			time.Sleep(400000)
			testSaver.Close()
			state := testSaver.GetState()
			Expect(state.ResultCode).Should(BeEquivalentTo(saver.ErrorSaverResultCode))
			Expect(state.ErrorOfSaving).ShouldNot(BeNil())
		})
	})
})
