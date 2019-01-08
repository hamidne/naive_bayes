package main

import "fmt"

func main()  {
	var train [5]string

	train[1] = "a tehran iran tehran tehran iran"
	train[2] = "a isfahan iran isfahan"
	train[3] = "a iran shiraz fars"
	train[4] = "b fars isfahan shiraz"

	var test = "iran isfahan shiraz tehran"

	fmt.Println(test)
}