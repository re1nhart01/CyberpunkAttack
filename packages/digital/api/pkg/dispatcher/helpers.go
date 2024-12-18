package dispatcher

func filter[T any](ss []T, test func(T, int) bool) (ret []T) {
	for i, s := range ss {
		if test(s, i) {
			ret = append(ret, s)
		}
	}
	return
}
