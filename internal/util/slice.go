package util

func Map[S, T any](source []S, converter func(source S) T) []T {
	converted := make([]T, 0, len(source))
	for _, s := range source {
		converted = append(converted, converter(s))
	}

	return converted
}

func Filter[T any](source []T, filter func(source T) bool) []T {
	filtered := make([]T, 0, len(source))
	for _, s := range source {
		if filter(s) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

func Reduce[T any](source []T, reducer func(elem1, elem2 T) T) (t T) {
	switch len(source) {
	case 0:
		return
	case 1:
		return source[0]
	default:
		t = source[0]
		for _, s := range source[1:] {
			t = reducer(t, s)
		}

		return
	}
}
