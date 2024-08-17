package utils

func FilterItems[T any](filter func(item T) bool, items []T) []T {
	res := []T{}

	for _, item := range items {
		if filter(item) {
			res = append(res, item)
		}
	}

	return res
}
