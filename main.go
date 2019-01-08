package main

import (
	"fmt"
	"regexp"
)

func main() {
	var train [4]string

	train[0] = "a tehran iran tehran tehran iran"
	train[1] = "a isfahan iran isfahan"
	train[2] = "a iran shiraz fars"
	train[3] = "b fars isfahan shiraz"

	processedData := make(map[byte]map[string]byte)

	var allItems []string
	var regex = regexp.MustCompile(`(?m) \w+`)

	for _, train := range train {

		class := []byte(train[0:1])[0]
		data := train[1:]

		if len(processedData[class]) == 0 {
			processedData[class] = map[string]byte{}
		}

		for _, match := range regex.FindAllString(data, -1) {
			processedData[class][match]++
			allItems = append(allItems, match)
		}
	}

	fmt.Println(processedData)

	var test = " " + "iran isfahan shiraz tehran"

	for _, match := range regex.FindAllString(test, -1) {
		var point float32 = 1
		for class, _ := range processedData {
			fmt.Println(match, float32(processedData[class][match]+1)/12)
			point *= float32(processedData[class][match]+1) / 12
		}
		fmt.Println(point)
	}

	//fmt.Println(test)
}
