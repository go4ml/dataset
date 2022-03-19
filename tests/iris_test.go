package tests

import (
	"go4ml.xyz/dataset/iris"
	"gotest.tools/assert"
	"testing"
)

func Test_Iris(t *testing.T) {
	n := iris.Data.LuckyCount()
	assert.Assert(t, n == 150)
}
