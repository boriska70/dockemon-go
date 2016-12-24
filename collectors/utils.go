package collectors

// MapToArray is an auxiliary method to create array from maps (for example for image or container labels).
func MapToArray(m map[string]string) []string {

	res := make([]string, 0)

	for ind, val := range m {
		res = append(res, ind+"="+val)
	}

	return res
}
