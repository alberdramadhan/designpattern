package stringutil_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"unit_test/stringutil"
)



var _ = Describe("Stringutil", func() {
    Describe("Reverse", func() {
        Context("when given a string", func() {
            It("should reverse the string", func() {
                Expect(stringutil.Reverse("hello")).To(Equal("olleh"))
            })
        })
    })

    Describe("ToUpper", func() {
        Context("when given a string", func() {
            It("should convert the string to upper case", func() {
                Expect(stringutil.ToUpper("hello")).To(Equal("HELLO"))
            })
        })
    })

    Describe("Concat", func() {
        Context("when given multiple strings", func() {
            It("should concatenate the strings", func() {
                Expect(stringutil.Concat("hello", " ", "world")).To(Equal("hello world"))
            })
        })
    })

    Describe("IsEmpty", func() {
        Context("when given an empty string", func() {
            It("should return true", func() {
                Expect(stringutil.IsEmpty("")).To(BeTrue())
            })
        })

        Context("when given a non-empty string", func() {
            It("should return false", func() {
                Expect(stringutil.IsEmpty("hello")).To(BeFalse())
            })
        })
    })

    Describe("FirstRune", func() {
        Context("when given a non-empty string", func() {
            It("should return the first rune of the string", func() {
                Expect(stringutil.FirstRune("hello")).To(Equal('h'))
            })
        })

        Context("when given an empty string", func() {
            It("should return 0", func() {
                Expect(stringutil.FirstRune("")).To(Equal(rune(0)))
            })
        })
    })

    Describe("GetStringFunction", func() {
        Context("when given a valid function type", func() {
            It("should return the corresponding function", func() {
                funcType := "reverse"
                fn := stringutil.GetStringFunction(funcType)
                Expect(fn("hello")).To(Equal("olleh"))
            })
        })

        Context("when given an invalid function type", func() {
            It("should return an error message", func() {
                funcType := "invalid"
                fn := stringutil.GetStringFunction(funcType)
                Expect(fn("hello")).To(Equal("Invalid function type"))
            })
        })
    })
})

