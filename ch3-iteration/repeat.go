package ch3iteration

func RepeatN(str string, n uint) string {
	var repeated string
	for i := uint(0); i < n; i++ {
		repeated += str
	}
	return repeated
}
