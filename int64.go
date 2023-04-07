package goutils

import "strconv"

func Int64FromString(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0
	}
	return out
}

func Int64ToBytes(in int64) []byte {
	if in > 0 {
		return []byte(strconv.FormatInt(in, 10))
	}
	return nil
}
