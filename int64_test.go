package goutils

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestInt64FromString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int64
	}{
		{name: "test 1", in: "1", want: 1},
		{name: "test 10", in: "10", want: 10},
		{name: "test 100", in: "100", want: 100},
		{name: "test 1000", in: "1000", want: 1000},
		{name: "words", in: "nope", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int64FromString(tt.in), "Int64FromString(%v)", tt.in)
		})
	}
}

func TestInt64ToBytes(t *testing.T) {
	tests := []struct {
		name string
		in   int64
		want []byte
	}{
		{name: "test 1", in: 1, want: []byte{0x31}},
		{name: "test 10", in: 10, want: []byte{0x31, 0x30}},
		{name: "test 100", in: 100, want: []byte{0x31, 0x30, 0x30}},
		{name: "max", in: math.MaxInt64, want: []byte{0x39, 0x32, 0x32, 0x33, 0x33, 0x37, 0x32, 0x30, 0x33, 0x36, 0x38, 0x35, 0x34, 0x37, 0x37, 0x35, 0x38, 0x30, 0x37}},
		{name: "negative", in: -2, want: []byte(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int64ToBytes(tt.in), "Int64ToBytes(%v)", tt.in)
		})
	}
}
