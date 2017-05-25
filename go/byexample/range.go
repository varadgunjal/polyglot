package main

import "fmt"

func main() {
    a := map[string]int{"a": 1, "b":2}

    for k, v := range a {
        fmt.Println(k, v)
    }
}