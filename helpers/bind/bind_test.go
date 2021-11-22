package bind

import (
	"github.com/lubovskiy/crud/helpers"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialGenerator_Next_Success(t *testing.T) {
	startNum := helpers.Int(1, 10000)
	gen := NewSequentialGenerator(startNum)

	n1 := gen.Next()
	assert.Equal(t, startNum, n1)
	n2 := gen.Next()
	assert.Equal(t, startNum+1, n2)
	n3 := gen.Next()
	assert.Equal(t, startNum+2, n3)
	for i := 0; i < 1000; i++ {
		_ = gen.Next()
	}
	nX := gen.Next()
	assert.Equal(t, startNum+1003, nX)
}

func TestSequentialGenerator_NextBind_Success(t *testing.T) {
	startNum := helpers.Int(1, 10000)
	gen := NewSequentialGenerator(startNum)

	n1 := gen.NextBind()
	assert.Equal(t, "$"+strconv.FormatInt(int64(startNum), 10), n1)
	n2 := gen.NextBind()
	assert.Equal(t, "$"+strconv.FormatInt(int64(startNum+1), 10), n2)
	n3 := gen.NextBind()
	assert.Equal(t, "$"+strconv.FormatInt(int64(startNum+2), 10), n3)
	for i := 0; i < 1000; i++ {
		_ = gen.NextBind()
	}
	nX := gen.NextBind()
	assert.Equal(t, "$"+strconv.FormatInt(int64(startNum+1003), 10), nX)
}
