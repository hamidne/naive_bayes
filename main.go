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

	for _, train := range train {
		class := train[0:1]
		data := train[1:]

		var re = regexp.MustCompile(`(?m) \w+`)

		for _, match := range re.FindAllString(data, -1) {
			fmt.Println(class, " : ", match)
		}
	}

	//var test = "iran isfahan shiraz tehran"
	//fmt.Println(test)
}
