package main

import(
"fmt"
"math"
) 

const s string="value"

func main(){
fmt.Println(s)
const n = 1000
const m = 304654/n
fmt.Println(int64(m))
fmt.Println(float32(m))
fmt.Println(math.Cos(m))
}