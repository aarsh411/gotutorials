package main


import (
	"fmt"
)

func main() {

	var languages = map[int]string{

		3: "English",

		4: "Hindi",

		5: "Gujarati",
	}

	var products = make(map[int]string)

	products[1] = "item1"
	products[2] = "item2"

	var i int
	var value string

	for i, value = range languages {

		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 1", products[1])

	delete(products, 2)

	fmt.Println("products", products)
}
