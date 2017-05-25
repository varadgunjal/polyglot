package main

import "fmt"

func plus (data ...int) int {
    sum := 0
    for _, v := range data {
        sum += v
    }
    return sum
}

func main() {
    res := plus(1, 2, 3, 4, 5, 6)
    fmt.Println(res)
}