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
		{name: "test 1", in: 1, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}},
		{name: "test 10", in: 10, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}},
		{name: "test 100", in: 100, want: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64}},
		{name: "max", in: math.MaxInt64, want: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{name: "min", in: math.MinInt64, want: []byte{0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{name: "negative", in: -2, want: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int64ToBytes(tt.in), "Int64ToBytes(%v)", tt.in)
		})
	}
}

func TestInt64FromBytes(t *testing.T) {
	tests := []struct {
		name string
		want int64
		in   []byte
	}{
		{name: "test 1", want: 1, in: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}},
		{name: "test 10", want: 10, in: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}},
		{name: "test 100", want: 100, in: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x64}},
		{name: "max", want: math.MaxInt64, in: []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{name: "min", want: math.MinInt64, in: []byte{0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{name: "negative", want: -2, in: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int64FromBytes(tt.in), "Int64FromBytes(%v)", tt.in)
		})
	}
}
