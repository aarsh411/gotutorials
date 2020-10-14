package main

import "fmt"

func main() {
	fmt.Println("number is 7")

    if 7%2 == 0 {
        fmt.Println("number is even")
    } else {
        fmt.Println("number is odd")
    }

    if 7%4 == 0 {
        fmt.Println("number is divisible by 4")
    }else {
	fmt.Println("number is not divisible by 4")
	}

    if num := 122; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}