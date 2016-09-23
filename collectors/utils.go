package collectors

func MapToArray(m map[string]string) [] string  {

	res := make([]string, 0)

	for ind,val := range m {
		res = append(res,ind+"="+val)
	}

	return res
}
