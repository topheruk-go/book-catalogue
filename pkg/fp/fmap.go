package fp

func FMap[T any, U any](ns []T, fn func(T) U) []U {
	ms := make([]U, len(ns))
	for i, v := range ns {
		ms[i] = fn(v)
	}
	return ms
}
