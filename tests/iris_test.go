package tests

import (
	"go-ml.dev/pkg/dataset/iris"
	"gotest.tools/assert"
	"testing"
)

func Test_Iris(t *testing.T) {
	n := iris.Data.LuckyCount()
	assert.Assert(t, n == 150)
}
