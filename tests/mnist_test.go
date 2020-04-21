package tests

import (
	"go-ml.dev/pkg/dataset/mnist"
	"gotest.tools/assert"
	"testing"
)

func Test_Mnist(t *testing.T) {
	n := mnist.Data.LuckyCount()
	assert.Assert(t, n == 60000)
	n = mnist.T10k.LuckyCount()
	assert.Assert(t, n == 10000)
}
