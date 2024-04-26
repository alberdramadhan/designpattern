package stringutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStringutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringutil Suite")
}
