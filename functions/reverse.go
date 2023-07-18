package ascart

import "fmt"

var globalNums []int

func Reverse(fileName string) {
	FileIn := ReadWholeFile(fileName)
	s1, s2, s3 := "./arts/standard.txt", "./arts/shadow.txt", "./arts/thinkertoy.txt"
	arr_first := ReadWholeFile(s1)
	arr_all := Connect_Slice(arr_first, ReadWholeFile(s2), ReadWholeFile(s3))
	Output(getAll(arr_all), letter(string(FileIn), Divide(string(FileIn))))
}

func Output(main, oldWords [][]rune) {
	if checkDistance(globalNums) {
		fmt.Println("Incorrect symbols on your file! Please enter correct symbols!")
		return
	}
	var res string
	words := Change(oldWords)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(main); j++ {
			if RuneToString(words[i]) == RuneToString(main[j]) {
				k := j
				if k >= 127 && k < 222 {
					k -= 95
				} else if k >= 222 && k < 319 {
					k -= 190
				}
				res += string(rune(k + 32))
			}
		}
	}
	if len(globalNums) != len(res) {
		fmt.Println("Incorrect symbols on your file! Please enter correct symbols!")
		return
	}
	if res != "" {
		fmt.Println(res)
	} else {
		fmt.Println("Incorrect symbols on your file! Please enter correct symbols!")
	}
}

func Change(words [][]rune) [][]rune {
	var res [][]rune
	for i := 0; i < len(words); i++ {
		var tmRes []rune
		tmRes = append(tmRes, 10)
		for j := 0; j < len(words[i]); j++ {
			if j+1 != len(words[i]) {
				tmRes = append(tmRes, rune(words[i][j]))
			} else {
				tmRes = append(tmRes, 10)
			}
		}
		res = append(res, tmRes)
	}
	return res
}

func getAll(arr []rune) [][]rune {
	var doubleArr [][]rune
	var res []rune
	for i := 0; i < 285; i++ {
		num := i * 9
		x := 0
		j := 0
		for j < len(arr) {
			if x < num {
				if arr[j] == 10 || arr[j] == 13 {
					x++
				}
			} else {
				break
			}
			j++
		}
		y := 0
		for k := j; k < len(arr); k++ {
			if arr[k] == 10 {
				y++
			}
			if y < 10 {
				res = append(res, arr[k])
			}
		}
		doubleArr = append(doubleArr, res)

		res = []rune{}
	}
	return doubleArr
}

func Divide(s string) []int {
	var nums []int
	var check int
	var slsh int
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			slsh = i
			break
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			check = i
			j := i
			cnt := 0
			for j < len(s) {
				if s[j] == '\n' {
					if j+check < len(s) && s[j+check] == ' ' {
						cnt++
					} else {
						break
					}
				}
				j++
			}
			if cnt == 7 {
				nums = append(nums, check)
			}
		}
	}
	nums = append(nums, slsh)
	globalNums = nums
	return nums
}

func checkDistance(n []int) bool {
	for i := 0; i < len(n); i++ {
		if i+1 < len(n) {
			if n[i+1]-n[i] > 15 {
				return true
			}
		}
	}
	return false
}

func letter(s string, nums []int) [][]rune {
	var res [][]rune
	var part []rune
	for i := 0; i < len(nums); i++ {
		j := 0
		timer := nums[i]
		timerS := nums[i]
		if i != 0 {
			j = nums[i-1] + 1
		}
		jStart := j
		x := j
		eight := 0
		for eight < 8 {
			if i != 0 && eight == 0 {
				x--
			}
			for x < timer {
				if x < len(s) {
					part = append(part, rune(s[x]))
				} else {
					break
				}
				x++
			}
			if eight != 0 {
				part = append(part, ' ')
			}
			if eight == 7 {
				part = append(part, ' ')
			}
			for x < len(s) {
				if s[x] == '\n' && x+1 != len(s) {
					timer = timerS + x
					x += jStart
					if i != 0 {
						part = append(part, '\n')
					}
					break
				} else {
					x++
				}
			}
			eight++
		}
		res = append(res, part)
		part = nil
	}
	return res
}

func Connect_Slice(arr, arr2, arr3 []rune) []rune {
	for i := 0; i < len(arr2); i++ {
		arr = append(arr, arr2[i])
	}
	for i := 0; i < len(arr3); i++ {
		arr = append(arr, arr3[i])
	}
	return arr
}

