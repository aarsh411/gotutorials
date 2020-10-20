package main
import (
	"fmt"
	"strconv"
)

func selectionSort(arr []int) {
	var length = len(arr)
	for i := 0; i < length - 1; i++ {
		var minIndex = i;
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
            
		var temp = arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int { 15, 3, 12, 6, -9, 9, 0 }

	fmt.Print("Before Sorting: ")
	printArray(arr)

	selectionSort(arr)
	fmt.Print("After Sorting: ")
	printArray(arr)
}