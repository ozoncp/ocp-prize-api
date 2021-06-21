package producer_test

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-prize-api/internal/producer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Producer", func() {
	var (
		TestProducer producer.IProducer
	)
	Context("producer test", func() {

		BeforeEach(func() {
			TestProducer = producer.NewProducer(context.Background(), "OcpPrizeApi")
			time.Sleep(1 * time.Second)
		})

		AfterEach(func() {
			err := TestProducer.Close()
			Expect(err).Should(BeNil())
		})

		It("Test producer send message", func() {
			res := TestProducer.SendMessage("Test message")
			Expect(res).Should(BeEquivalentTo(true))
			time.Sleep(1 * time.Second)
		})

		It("Test producer send few message", func() {
			res := TestProducer.SendMessage("Test message")
			Expect(res).Should(BeEquivalentTo(true))
			res = TestProducer.SendMessage("Test message2")
			Expect(res).Should(BeEquivalentTo(true))
			res = TestProducer.SendMessage("Test message3")
			Expect(res).Should(BeEquivalentTo(true))
			time.Sleep(1 * time.Second)
		})
	})

	Context("producer test with few producers", func() {
		It("Test send to closed producer", func() {

			TestProducer = producer.NewProducer(context.Background(), "OcpPrizeApi")
			time.Sleep(1 * time.Second)
			err := TestProducer.Close()
			Expect(err).Should(BeNil())
			time.Sleep(1 * time.Second)
			res := TestProducer.SendMessage("Test message")
			Expect(res).Should(BeEquivalentTo(false))
			err = TestProducer.Close()
			Expect(err).ShouldNot(BeNil())
			time.Sleep(1 * time.Second)
		})
		It("Test multi close", func() {

			TestProducer = producer.NewProducer(context.Background(), "OcpPrizeApi")
			time.Sleep(1 * time.Second)
			err := TestProducer.Close()
			Expect(err).Should(BeNil())
			err = TestProducer.Close()
			Expect(err).ShouldNot(BeNil())
			err = TestProducer.Close()
			Expect(err).ShouldNot(BeNil())
		})
	})
})
