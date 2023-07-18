package ascart

func RuneToStringDbl(db [][]rune) [][]string {
	var mainRes [][]string
	var res []string
	for i := 0; i < len(db); i++ {
		for j := 0; j < len(db[i]); j++ {
			res = append(res, string(db[i][j]))
		}
		mainRes = append(mainRes, res)
		res = nil
	}
	return mainRes
}

func RuneToStringDouble(db [][]rune) string {
	var res string
	for i := 0; i < len(db); i++ {
		for j := 0; j < len(db[i]); j++ {
			res += string(db[i][j])
		}
	}
	return res
}

func RuneToString(r []rune) string {
	var res string
	for _, i := range r {
		res += string(i)
	}
	return res
}

func Index(s string) []int {
	var res []int
	for i := 0; i < len(s); i++ {
		res = append(res, int(s[i]-32))
	}
	return res
}
