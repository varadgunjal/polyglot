package main

import "fmt"

func main() {
      a := [5]int{1, 2, 3, 4, 5}
      fmt.Println("start: ", a)

      var twoD [2][2]int
      for i := 0; i < 2; i++ {
            for j := 0; j < 2; j++ {
                  twoD[i][j] = i+ j
            }
      }

      fmt.Println(twoD)
}