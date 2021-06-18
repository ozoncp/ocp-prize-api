package flusher_test

import (
	"context"
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
		ctx         context.Context
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockIRepo
		testFlusher flusher.IFlusher

		prizesToAdd []prize.Prize
		leftPrizes  []prize.Prize
		errorFlush  error
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockIRepo(ctrl)
	})

	Context("correct flushing prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockRepo.EXPECT().AddPrizes(gomock.Any(), gomock.Any()).Return([]uint64{1}, nil).MinTimes(1)
		})

		It("Test flushing with correct input", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes, _, errorFlush = testFlusher.Flush(ctx, prizesToAdd, nil)
			Expect(leftPrizes).Should(BeNil())
			Expect(errorFlush).Should(BeNil())
		})
	})

	Context("incorrect flushing prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}

			mockRepo.EXPECT().AddPrizes(gomock.Any(), gomock.Any()).Return([]uint64{0}, errors.New("add prize error")).MinTimes(1)
		})

		It("Test flushing with error at first try to add bunch", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes, _, errorFlush = testFlusher.Flush(ctx, prizesToAdd, nil)
			Expect(leftPrizes).Should(BeEquivalentTo(prizesToAdd))
			Expect(errorFlush).ShouldNot(BeNil())
		})
	})

	Context("incorrect flushing last 2 prizes", func() {

		BeforeEach(func() {
			prizesToAdd = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}
			gomock.InOrder(
				mockRepo.EXPECT().AddPrizes(gomock.Any(), gomock.Any()).Return([]uint64{1}, nil).Times(1),
				mockRepo.EXPECT().AddPrizes(gomock.Any(), gomock.Any()).Return([]uint64{0}, errors.New("add prize error")).Times(1),
			)
		})

		It("Test flushing with error at second try to add bunch", func() {
			testFlusher = flusher.NewFlusher(mockRepo, 3)
			leftPrizes, _, errorFlush = testFlusher.Flush(ctx, prizesToAdd, nil)
			Expect(len(leftPrizes)).Should(BeEquivalentTo(2))
			Expect(errorFlush).ShouldNot(BeNil())
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})
})
