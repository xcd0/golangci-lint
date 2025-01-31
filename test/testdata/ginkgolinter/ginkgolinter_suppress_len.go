//golangcitest:config_path configs/ginkgolinter_suppress_len.yml
//golangcitest:args --disable-all -Eginkgolinter
package ginkgolinter

import (
	"errors"
	. "github.com/onsi/gomega"
)

// LenUsecase_len should not trigger any warning
func LenUsecase_len() {
	var fakeVarUnderTest []int
	Expect(fakeVarUnderTest).Should(BeEmpty())     // valid
	Expect(fakeVarUnderTest).ShouldNot(HaveLen(5)) // valid

	Expect(len(fakeVarUnderTest)).Should(Equal(0))
	Expect(len(fakeVarUnderTest)).ShouldNot(Equal(2))
	Expect(len(fakeVarUnderTest)).To(BeNumerically("==", 0))

	fakeVarUnderTest = append(fakeVarUnderTest, 3)
	Expect(len(fakeVarUnderTest)).ShouldNot(Equal(0))
	Expect(len(fakeVarUnderTest)).Should(Equal(1))
	Expect(len(fakeVarUnderTest)).To(BeNumerically(">", 0))
	Expect(len(fakeVarUnderTest)).To(BeNumerically(">=", 1))
	Expect(len(fakeVarUnderTest)).To(BeNumerically("!=", 0))
}

func NilUsecase_len() {
	y := 5
	x := &y
	Expect(x == nil).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion; consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\). instead"
	Expect(nil == x).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion; consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\). instead"
	Expect(x != nil).To(Equal(true)) // want "ginkgo-linter: wrong nil assertion; consider using .Expect\\(x\\)\\.ToNot\\(BeNil\\(\\)\\). instead"
	Expect(x == nil).To(BeTrue())    // want "ginkgo-linter: wrong nil assertion; consider using .Expect\\(x\\)\\.To\\(BeNil\\(\\)\\). instead"
	Expect(x == nil).To(BeFalse())   // want "ginkgo-linter: wrong nil assertion; consider using .Expect\\(x\\)\\.ToNot\\(BeNil\\(\\)\\). instead"
}
func BooleanUsecase_len() {
	x := true
	Expect(x).To(Equal(true)) // want "ginkgo-linter: wrong boolean assertion; consider using .Expect\\(x\\)\\.To\\(BeTrue\\(\\)\\). instead"
	x = false
	Expect(x).To(Equal(false)) // want "ginkgo-linter: wrong boolean assertion; consider using .Expect\\(x\\)\\.To\\(BeFalse\\(\\)\\). instead"
}

func ErrorUsecase_len() {
	err := errors.New("fake error")
	funcReturnsErr := func() error { return err }

	Expect(err).To(BeNil())              // want "ginkgo-linter: wrong error assertion; consider using .Expect\\(err\\)\\.ToNot\\(HaveOccurred\\(\\)\\). instead"
	Expect(err == nil).To(Equal(true))   // want "ginkgo-linter: wrong error assertion; consider using .Expect\\(err\\)\\.ToNot\\(HaveOccurred\\(\\)\\). instead"
	Expect(err == nil).To(BeFalse())     // want "ginkgo-linter: wrong error assertion; consider using .Expect\\(err\\)\\.To\\(HaveOccurred\\(\\)\\). instead"
	Expect(err != nil).To(BeTrue())      // want "ginkgo-linter: wrong error assertion; consider using .Expect\\(err\\)\\.To\\(HaveOccurred\\(\\)\\). instead"
	Expect(funcReturnsErr()).To(BeNil()) // want "ginkgo-linter: wrong error assertion; consider using .Expect\\(funcReturnsErr\\(\\)\\)\\.To\\(Succeed\\(\\)\\). instead"
}

func HaveLen0Usecase_len() {
	x := make([]string, 0)
	Expect(x).To(HaveLen(0)) // want "ginkgo-linter: wrong length assertion; consider using .Expect\\(x\\)\\.To\\(BeEmpty\\(\\)\\). instead"
}

func WrongComparisonUsecase_len() {
	x := 8
	Expect(x == 8).To(BeTrue())    // want "ginkgo-linter: wrong comparison assertion; consider using .Expect\\(x\\)\\.To\\(Equal\\(8\\)\\). instead"
	Expect(x < 9).To(BeTrue())     // want "ginkgo-linter: wrong comparison assertion; consider using .Expect\\(x\\)\\.To\\(BeNumerically\\(\"<\", 9\\)\\). instead"
	Expect(x < 7).To(Equal(false)) // want "ginkgo-linter: wrong comparison assertion; consider using .Expect\\(x\\)\\.ToNot\\(BeNumerically\\(\"<\", 7\\)\\). instead"

	p1, p2 := &x, &x
	Expect(p1 == p2).To(Equal(true)) // want "ginkgo-linter: wrong comparison assertion; consider using .Expect\\(p1\\).To\\(BeIdenticalTo\\(p2\\)\\). instead"
}
