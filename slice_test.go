package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceContains(t *testing.T) {
	assert.True(t, SliceContains[string]("a", []string{"c", "b", "a"}))
	assert.False(t, SliceContains[string]("z", []string{"c", "b", "a"}))

	assert.True(t, SliceContains[int](4, []int{1, 2, 3, 4}))
	assert.False(t, SliceContains[int](10, []int{1, 2, 3, 4}))

	type TestObject struct {
		S string
		I int
	}

	assert.True(t, SliceContains[TestObject](
		TestObject{S: "mystring", I: 2}, []TestObject{
			TestObject{S: "mystring", I: 0},
			TestObject{S: "mystring", I: 1},
			TestObject{S: "mystring", I: 2},
		}),
	)

	assert.False(t, SliceContains[TestObject](
		TestObject{S: "mystring", I: 3}, []TestObject{
			TestObject{S: "mystring", I: 0},
			TestObject{S: "mystring", I: 1},
			TestObject{S: "mystring", I: 2},
		}),
	)

}
