package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func naiveBaise(data []string, test string) map[byte]float32 {

	classCount := make(map[byte]byte)
	itemCount := make(map[string]byte)
	regex := regexp.MustCompile(`(?m) \w+`)
	processedData := make(map[byte]map[string]byte)

	for _, train := range data {

		class := []byte(train[0:1])[0]
		classCount[class]++

		if len(processedData[class]) == 0 {
			processedData[class] = map[string]byte{}
		}

		for _, match := range regex.FindAllString(train[1:], -1) {
			itemCount[match]++
			processedData[class][match]++
		}
	}

	result := make(map[byte]float32)

	for _, match := range regex.FindAllString(test, -1) {
		for class := range processedData {
			classCount[class] *= float32(processedData[class][match]+1) / float32(len(uniqueCity)+cityCount[class])
		}
	}

	for class := range processedData {
		fmt.Println(string(class), classCount[class]/float32(len(data)))
	}

	return result
}

func getData() []string {

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

	uniqueCity := make(map[string]bool)
	classCount := make(map[byte]float32)

	trainData := getData()

	var test = " Chinese Chinese Chinese Tokyo Japan"

}
