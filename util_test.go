package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestInSlice(t *testing.T) {

	intList := []int{1, 2, 3}

	isInSlice := InSlice(1, intList)
	notInSlice := InSlice(4, intList)
	assert.Assert(t, isInSlice)
	assert.Assert(t, !notInSlice)
}
