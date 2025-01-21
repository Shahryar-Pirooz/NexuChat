package fp

func Map[T any, U any](src []T, mapper func(T) *U) []U {
	dst := make([]U, len(src))
	for i := range src {
		dst[i] = *mapper(src[i])
	}
	return dst
}
