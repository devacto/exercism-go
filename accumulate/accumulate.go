package accumulate

const testVersion = 1

func Accumulate(i []string, f func(string) string) []string {
	var r []string
	r = make([]string, len(i), len(i))
	for k := range i {
		r[k] = f(i[k])
	}
	return r
}
