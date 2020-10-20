package main
import "fmt"
func BubbleSort(numbers []int) []int {
   
   for i := len(numbers); i > 0; i-- {
      
      for j := 1; j < i; j++ {
         if numbers[j-1] > numbers[j] {
            intermediate := numbers[j]
            numbers[j] = numbers[j-1]
            numbers[j-1] = intermediate
                        }
             }
      }
   return numbers
 }
func main() {
    a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println(BubbleSort(a))
}