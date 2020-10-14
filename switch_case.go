package main

import (
    "fmt"
    "time"
	
	
)

func main() {
	
    fmt.Print("Write between 1 to 5 ")
	var i int
	fmt.Scanf("%d",&i)
   
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
	  case 4:
        fmt.Println("four")
	  case 5:
        fmt.Println("five")
	
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    
}