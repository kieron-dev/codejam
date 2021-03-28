package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Alien", func() {

	DescribeTable("scoring", func(prog string, score int) {

		Expect(Score(prog)).To(Equal(score))
	},
		Entry("S", "S", 1),
		Entry("CS", "CS", 2),
		Entry("SCCSSC", "SCCSSC", 9),
	)

	DescribeTable("swap at index", func(prog string, idx int, res string) {
		Expect(SwapAtIndex(prog, idx)).To(Equal(res))
	},

		Entry("1", "CCCCSCC", 3, "CCCSCCC"),
		Entry("2", "CSCCCCC", 0, "SCCCCCC"),
	)
})
