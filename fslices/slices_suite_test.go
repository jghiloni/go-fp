package fslices_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFslices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fslices Suite")
}
