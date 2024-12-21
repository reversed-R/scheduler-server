package internal

func SliceMap[T interface{}, U interface{}](slice []U, operation func(U) T) []T {
	var result []T
	for _, value := range slice {
		result = append(result, operation(value))
	}

	return result
}
