package goutils

func SliceContains[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func SliceRemoveValue[T comparable](list []T, value T) []T {
	for idx, v := range list {
		if v == value {
			return SliceRemoveIndex(list, idx)
		}
	}
	return list
}

func SliceRemoveIndex[T comparable](list []T, idx int) []T {
	return append(list[:idx], list[idx+1:]...)
}
