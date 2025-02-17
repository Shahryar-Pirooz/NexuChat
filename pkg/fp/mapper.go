package fp

func Map[T, U any](src []T, mapper func(T) U) []U {
	result := make([]U, len(src))

	for i := range src {
		result[i] = mapper(src[i])
	}
	return result
}
