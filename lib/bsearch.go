package lib

import "cmp"

func BinarySearch[T cmp.Ordered](data []T, target T) int {
	lo := 0
	hi := len(data) - 1
	for lo <= hi {
        mid := (lo + hi) / 2
		value := data[mid]

		if value == target {
			return mid
		}
		if value < target {
            lo = mid + 1
		} else {
            hi = mid - 1
		}
	}
	return -1
}
