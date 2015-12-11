package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWave(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wave Suite")
}
