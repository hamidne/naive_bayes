package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func getTrainData() []string {
	var train []string
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		train = append(train, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return train
}

func main() {

	var regex = regexp.MustCompile(`(?m) \w+`)

	cityCount := make(map[byte]int)
	uniqueCity := make(map[string]bool)
	classCount := make(map[byte]float32)
	processedData := make(map[byte]map[string]byte)

	trainData := getTrainData()

	for _, train := range trainData {

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

	var test = " Chinese Chinese Chinese Tokyo Japan"

	for _, match := range regex.FindAllString(test, -1) {
		for class := range processedData {
			classCount[class] *= float32(processedData[class][match]+1) / float32(len(uniqueCity)+cityCount[class])
		}
	}

	for class := range processedData {
		fmt.Println(class, classCount[class]/float32(len(trainData)))
	}

}
