package prize_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-prize-api/internal/prize"
)

var _ = Describe("Prize", func() {
	var (
		testSlice []prize.Prize
	)
	Context("test prize", func() {
		It("Test prize to string", func() {
			testPrize := prize.NewPrize(1, 2, "www")
			Expect(testPrize.String()).Should(BeEquivalentTo("ID: 1 IssueID: 2 Link: www"))
		})
	})
	Context("test multiple prize slice processing", func() {
		BeforeEach(func() {
			testSlice = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www"), prize.NewPrize(6, 2, "www")}
		})
		It("Test split multiple prize slice", func() {
			var splittedSize int = 2
			resultSlice, err := prize.SplitPrizeSliceToBunches(testSlice, splittedSize)

			Expect(err).Should(BeNil())
			Expect(len(resultSlice)).Should(BeEquivalentTo(len(testSlice) / splittedSize))
			Expect(len(resultSlice[0])).Should(BeEquivalentTo(splittedSize))
			Expect(len(resultSlice[len(resultSlice)-1])).Should(BeEquivalentTo(splittedSize))
		})
	})

	Context("test non multiple prize slice processing", func() {
		BeforeEach(func() {
			testSlice = []prize.Prize{prize.NewPrize(1, 2, "www"),
				prize.NewPrize(2, 2, "www"), prize.NewPrize(3, 2, "www"),
				prize.NewPrize(4, 2, "www"), prize.NewPrize(5, 2, "www")}
		})
		It("Test split non multiple prize slice", func() {
			var splittedSize int = 2
			resultSlice, err := prize.SplitPrizeSliceToBunches(testSlice, splittedSize)
			Expect(err).Should(BeNil())
			Expect(len(resultSlice)).Should(BeEquivalentTo(len(testSlice)/splittedSize + 1))
			Expect(len(resultSlice[0])).Should(BeEquivalentTo(splittedSize))
			Expect(len(resultSlice[len(resultSlice)-1])).Should(BeEquivalentTo(len(testSlice) % splittedSize))
		})

		It("Test prize slice with zero size", func() {
			resultSlice, err := prize.SplitPrizeSliceToBunches(testSlice, 0)
			Expect(err).ShouldNot(BeNil())
			Expect(resultSlice).Should(BeNil())
		})

		It("Test prize slice with size equal origin", func() {
			resultSlice, err := prize.SplitPrizeSliceToBunches(testSlice, len(testSlice))
			result := reflect.DeepEqual(resultSlice[0], testSlice)
			Expect(err).Should(BeNil())
			Expect(result).Should(BeEquivalentTo(true))
		})

		It("Test prize slice with size more than origin", func() {
			resultSlice, err := prize.SplitPrizeSliceToBunches(testSlice, len(testSlice)+1)
			Expect(err).ShouldNot(BeNil())
			Expect(resultSlice).Should(BeNil())
		})
	})

})
