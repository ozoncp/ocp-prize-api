package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/ozoncp/ocp-prize-api/internal/flusher"
	"github.com/ozoncp/ocp-prize-api/internal/mocks"
	"github.com/ozoncp/ocp-prize-api/internal/prize"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher

		prizesToAdd []prize.Prize
		leftPrizes  []prize.Prize
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)
	})

	Context("correct flushing prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockRepo.EXPECT().AddPrizes(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes = testFlusher.Flush(prizesToAdd)
			Expect(leftPrizes).Should(BeNil())
		})
	})

	Context("incorrect flushing prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockRepo.EXPECT().AddPrizes(gomock.Any()).Return(errors.New("add prize error")).MinTimes(1)
		})

		It("", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes = testFlusher.Flush(prizesToAdd)
			Expect(leftPrizes).Should(BeEquivalentTo(prizesToAdd))
		})
	})

	Context("incorrect flushing last 2 prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}
			gomock.InOrder(
				mockRepo.EXPECT().AddPrizes(gomock.Any()).Return(nil).Times(1),
				mockRepo.EXPECT().AddPrizes(gomock.Any()).Return(errors.New("add prize error")).Times(1),
			)
		})

		It("", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes = testFlusher.Flush(prizesToAdd)
			Expect(len(leftPrizes)).Should(BeEquivalentTo(2))
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
