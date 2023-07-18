package ascart

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadWholeFile(style string) []rune {
	contents, err := ioutil.ReadFile(style)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(string(contents)) == 0 {
		fmt.Println("Error! Empty parametre!")
		return nil
	}
	res := string(contents)
	var arr []rune
	for i := 0; i < len(res); i++ {
		arr = append(arr, rune(res[i]))
	}
	return arr
}

func WriteToFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	content = content[:len(content)-1]
	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		fmt.Println(err.Error())
		return
	}
}
