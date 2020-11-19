package string_split

func Split(s string) []string {
	var r []string
	var start int
	var flag bool
	var kk int

	for k, v := range s {
		kk = k
		if v == '"' {
			flag = !flag
		}
		if flag == false && v == ' ' {
			r = append(r, s[start:kk])
			start = kk + 1
		}
	}
	r = append(r, s[start:kk+1])
	return r
}
