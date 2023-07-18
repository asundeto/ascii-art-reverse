package ascart

func CheckSymbol(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 31 || s[i] > 127 {
			if s[i] != 10 && s[i] != 13 {
				return true
			}
		}
	}
	return false
}

func CheckFormat(s string) bool {
	if s != "standard" && s != "shadow" && s != "thinkertoy" {
		return true
	}
	return false
}

func Param(arr []string) (string, string, string) {
	if len(arr) == 1 {
		return "", arr[0], "standard"
	} else if len(arr) == 2 {
		boly, fileName := CheckOutput(arr[0])
		if boly {
			return fileName, arr[1], "standard"
		} else {
			return "", arr[0], arr[1]
		}
	} else {
		boly, fileName := CheckOutput(arr[0])
		if boly {
			return fileName, arr[1], arr[2]
		} else {
			return fileName, arr[1], arr[2]
		}
	}
}

func CheckOutput(s string) (bool, string) {
	if len(s) < 10 {
		return false, ""
	}
	var check string
	for i := 0; i < 9; i++ {
		check += string(s[i])
	}
	if check == "--output=" {
		var fileName string
		for i := 9; i < len(s); i++ {
			fileName += string(s[i])
		}
		return true, fileName
	}
	return false, ""
}

func CheckFileName(s string) string {
	if len(s) < 15 {
		return ""
	}
	option := s[:10]
	name := s[10:]
	if option != "--reverse=" {
		return ""
	}
	if CheckTxt(name) {
		return ""
	}
	return name
}

func CheckTxt(s string) bool {
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			x++
		}
	}
	if x != 1 {
		return true
	}
	if s == "" {
		return false
	}
	if len(s) < 4 {
		return true
	}
	rev := MyReverse(s)
	if rev[:4] != "txt." {
		return true
	}
	return false
}

func MyReverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
