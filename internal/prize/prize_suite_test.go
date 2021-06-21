package prize_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrize(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prize Suite")
}
