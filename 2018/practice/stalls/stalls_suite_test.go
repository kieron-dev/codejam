package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStalls(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stalls Suite")
}
