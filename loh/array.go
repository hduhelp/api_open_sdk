package loh

func Array[T any, V any](a []T, f func(T) V) []V {
	b := make([]V, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}
