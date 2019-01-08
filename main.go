package main

import (
	"fmt"
	"regexp"
	//"regexp"
)

func main() {
	var train [4]string

	train[0] = "a tehran iran tehran tehran iran"
	train[1] = "a isfahan iran isfahan"
	train[2] = "a iran shiraz fars"
	train[3] = "b fars isfahan shiraz"

	processedData := make(map[byte]map[string]byte)

	classes := []byte{'a', 'b'}
	for _, class := range classes {
		processedData[class] = map[string]byte{}
	}

	for _, train := range train {
		class := []byte(train[0:1])[0]
		data := train[1:]

		var re = regexp.MustCompile(`(?m) \w+`)

		for _, match := range re.FindAllString(data, -1) {
			processedData[class][match]++
		}
	}

	fmt.Println(processedData)

	//var test = "iran isfahan shiraz tehran"
	//fmt.Println(test)
}
