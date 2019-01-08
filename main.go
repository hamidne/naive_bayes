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

	cityCount := make(map[byte]int)
	classCount := make(map[byte]float32)
	processedData := make(map[byte]map[string]byte)

	uniqueCity := make(map[string]bool)

	for _, train := range getTrainData() {

		class := []byte(train[0:1])[0]
		classCount[class]++

		if len(processedData[class]) == 0 {
			processedData[class] = map[string]byte{}
		}

		for _, match := range regex.FindAllString(train[1:], -1) {
			cityCount[class]++
			uniqueCity[match] = true
			processedData[class][match]++
		}
	}

	fmt.Println(uniqueCity)
	fmt.Println(cityCount)
	fmt.Println(processedData)

	var test = " " + "iran isfahan shiraz tehran"

	for _, match := range regex.FindAllString(test, -1) {
		for class := range processedData {
			fmt.Println(processedData[class][match], processedData[class][match]+1)
			fmt.Println(len(uniqueCity), cityCount[class])
			fmt.Println(float32(processedData[class][match]+1) / float32(len(uniqueCity)+cityCount[class]))
			classCount[class] *= float32(processedData[class][match]+1) / float32(len(uniqueCity)+cityCount[class])
		}
	}

}
