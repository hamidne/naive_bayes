package main

import (
	"fmt"
	"regexp"
)

func getTrainData() []string {
	// TODO: Load From File

	var train []string

	train = append(train, "a tehran iran tehran tehran iran")
	train = append(train, "a isfahan iran isfahan")
	train = append(train, "a iran shiraz fars")
	train = append(train, "b fars isfahan shiraz")

	return train
}

func main() {

	var regex = regexp.MustCompile(`(?m) \w+`)

	points := make(map[byte]float32)
	processedData := make(map[byte]map[string]byte)

	var allItems []string

	for _, train := range getTrainData() {

		data := train[1:]
		class := []byte(train[0:1])[0]

		points[class]++

		if len(processedData[class]) == 0 {
			processedData[class] = map[string]byte{}
		}

		for _, match := range regex.FindAllString(data, -1) {
			processedData[class][match]++
			allItems = append(allItems, match)
		}
	}

	fmt.Println(points)
	fmt.Println(processedData)

	var test = " " + "iran isfahan shiraz tehran"

	for _, match := range regex.FindAllString(test, -1) {
		var point float32 = 1
		for class := range processedData {
			fmt.Println(match, float32(processedData[class][match]+1)/12)
			point *= float32(processedData[class][match]+1) / 12
		}
		fmt.Println(point)
	}

	//fmt.Println(test)
}
