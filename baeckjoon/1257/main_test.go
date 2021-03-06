package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	assert := assert.New(t)
	// assert.Equal(uint64(100001), Solve(1000000001, []uint64{10000, 1}), "basic")
	assert.Equal(uint64(1), Solve(3, []uint64{3, 2, 1}), "test1")
	assert.Equal(uint64(2), Solve(4, []uint64{3, 2, 1}), "test2")
	assert.Equal(uint64(2), Solve(10, []uint64{1, 4, 5, 7}), "test3")
	assert.Equal(uint64(8), Solve(107, []uint64{1, 69, 100}), "test4")
	// assert.Equal(uint64(1000000000000000000), Solve(1000000000000000000, []uint64{1}), "test5")
	assert.Equal(uint64(10), Solve(100000, []uint64{1, 10000, 10000}), "test6")
	assert.Equal(uint64(2), Solve(8, []uint64{1, 4, 5, 6}), "test7")
	assert.Equal(uint64(3), Solve(65, []uint64{1, 4, 30, 31}), "test8")
}
