package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAlien(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alien Suite")
}
