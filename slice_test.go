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

func TestSliceRemoveIndex(t *testing.T) {
	type args[T comparable] struct {
		list []T
		idx  int
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{name: "level0", args: args[string]{list: []string{"aaa", "abc", "abb"}, idx: 0}, want: []string{"abc", "abb"}},
		{name: "level1", args: args[string]{list: []string{"aaa", "abc", "abb"}, idx: 1}, want: []string{"aaa", "abb"}},
		{name: "level2", args: args[string]{list: []string{"aaa", "abc", "abb"}, idx: 2}, want: []string{"aaa", "abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceRemoveIndex(tt.args.list, tt.args.idx), "SliceRemoveIndex(%v, %v)", tt.args.list, tt.args.idx)
		})
	}
}

func TestSliceRemoveValue(t *testing.T) {
	type args[T comparable] struct {
		list  []T
		value T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test 1",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 1},
			want: []int{2, 3, 4, 5, 6},
		},
		{
			name: "test 2",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 2},
			want: []int{1, 3, 4, 5, 6},
		},
		{
			name: "test 3",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 3},
			want: []int{1, 2, 4, 5, 6},
		},
		{
			name: "test 4",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 4},
			want: []int{1, 2, 3, 5, 6},
		},
		{
			name: "test 5",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 5},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test 6",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 6},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test 77",
			args: args[int]{list: []int{1, 2, 3, 4, 5, 6}, value: 77},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SliceRemoveValue(tt.args.list, tt.args.value), "SliceRemoveValue(%v, %v)", tt.args.list, tt.args.value)
		})
	}
}
