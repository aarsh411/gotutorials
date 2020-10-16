package main

import (
	"fmt"
)

func twiceValue(slice []int) {

	var i int
	var value int

	for i, value = range slice {

		slice[i] = 2 * value

	}

}

func main() {

	var slice = []int{0, 1, 8, 20}
	twiceValue(slice)

	var i int

	for i = 0; i < len(slice); i++ {

		fmt.Println("new slice value", slice[i])
	}
}
