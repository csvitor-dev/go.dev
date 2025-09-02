package slices

func Some[T any](slice []T, fn func(T, int) bool) bool {
	result := Filter(slice, fn)

	return len(result) != 0
}
