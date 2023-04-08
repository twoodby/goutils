package goutils

import (
	"encoding/binary"
	"strconv"
)

func Int64FromString(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0
	}
	return out
}

func Int64ToBytes(in int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(in))
	return b
}

func Int64FromBytes(in []byte) int64 {
	return int64(binary.BigEndian.Uint64(in))
}
