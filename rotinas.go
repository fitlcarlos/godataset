package godataset

import "slices"

func ClearSlice[S ~[]E, E any](s S) S {
	size := len(s)
	if size == 0 {
		return nil
	}
	return slices.Delete(s, 0, size-1)
}
