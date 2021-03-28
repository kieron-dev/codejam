package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Solve", func() {

	DescribeTable("example solutions",
		func(n, k, h, l uint64) {
			high, low := solve(n, k)
			Expect(high).To(Equal(h), "high")
			Expect(low).To(Equal(l), "low")
		},
		Entry("4 2", uint64(4), uint64(2), uint64(1), uint64(0)),
		Entry("5 2", uint64(5), uint64(2), uint64(1), uint64(0)),
		Entry("6 2", uint64(6), uint64(2), uint64(1), uint64(1)),
		Entry("1000 1000", uint64(1000), uint64(1000), uint64(0), uint64(0)),
		Entry("1000 1", uint64(1000), uint64(1), uint64(500), uint64(499)),
	)
})
