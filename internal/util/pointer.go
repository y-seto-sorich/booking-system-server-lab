package util

func Pointer[T any](v T) *T {
	return &v
}

func Dereference[T any](p *T) (v T) {
	if p != nil {
		v = *p
	}
	return
}
