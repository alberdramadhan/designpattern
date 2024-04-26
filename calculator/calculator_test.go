package calculator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"unit_test/calculator"
)



var _ = Describe("Calculator", func() {
	Describe("GetOperation", func() {
		Context("when given 'addition'", func() {
			It("should return an addition operation", func() {
				op := calculator.GetOperation("addition")
				Expect(op.Operate(9, 4)).To(Equal(13))
			})
		})

		Context("when given 'multiplication'", func() {
			It("should return a multiplication operation", func() {
				op := calculator.GetOperation("multiplication")
				Expect(op.Operate(9, 4)).To(Equal(36))
			})
		})

		Context("when given an unknown operation type", func() {
			It("should panic", func() {
				Expect(func() { calculator.GetOperation("division") }).To(Panic())
			})
		})
	})
})
