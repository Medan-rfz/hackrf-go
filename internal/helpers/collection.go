package helpers

func Where[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

func Select[T any, R any](collection []T, selector func(T) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = selector(item)
	}
	return result
}
