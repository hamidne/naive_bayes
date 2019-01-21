package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func mapSum(mapArray map[string]byte) byte {
	var sum byte

	for _, count := range mapArray {
		sum += count
	}

	return sum
}

func getMax(data map[byte]float32) byte {
	var name byte
	var max float32

	for key, value := range data {
		if value > max {
			name = key
			max = value
		}
	}

	return name
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

func naiveBasie(data []string, test string) map[byte]float32 {

	classCount := make(map[byte]byte)
	itemCount := make(map[string]byte)
	uniqeItems := make(map[string]bool)
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
			uniqeItems[match] = true
			processedData[class][match]++
		}
	}

	result := make(map[byte]float32)

	for nameC, countC := range classCount {
		result[nameC] = float32(countC) / float32(len(data))
		for _, nameI := range regex.FindAllString(test, -1) {
			result[nameC] *= float32(processedData[nameC][nameI]+1) / float32(mapSum(processedData[nameC])+byte(len(uniqeItems)))
		}
	}

	return result
}

func main() {

	//res := naiveBasie(getData(), " Chinese Chinese Chinese Tokyo Japan")
	//fmt.Println(res)

	originalData := getData()

	var tp, fp, tn, fn int

	for range originalData {
		var testData = originalData[0]
		originalData = originalData[1:]
		res := naiveBasie(originalData, testData[1:])

		if getMax(res) == 'c' {
			if testData[0] == 'c' {
				tp++
			} else {
				fp++
			}
		} else {
			if testData[0] == 'j' {
				tn++
			} else {
				fn++
			}
		}

		originalData = append(originalData, testData)
	}

	recall := float32(tp) / float32(tp+fn)
	precision := float32(tp) / float32(tp+fp)

	fmt.Println(`Recall:     `, recall)
	fmt.Println(`Precision:  `, precision)
	fmt.Println(`F1-measure: `, float32(2*recall*precision)/float32(recall+precision))

}
