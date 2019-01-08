package main

import "fmt"

func main() {
	var train [4]string

	train[0] = "a tehran iran tehran tehran iran"
	train[1] = "a isfahan iran isfahan"
	train[2] = "a iran shiraz fars"
	train[3] = "b fars isfahan shiraz"

	var test = "iran isfahan shiraz tehran"

	for _, train := range train {
		fmt.Println("original : ", train)
		fmt.Println("class : ", train[0:1])
		fmt.Println("data : ", train[2:])
		fmt.Println()
	}

	fmt.Println(test)
}
