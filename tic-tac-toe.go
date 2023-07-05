package main

import (
	"fmt"
	//  "reflect"
	"strconv"
)

func main() {
	// 	for i := 0; ; i++{
	//      if i%2 == 0 {
	//         input()[0]
	//      }
	//      else{
	//          input()[1]
	//      }
	// 	}
	var zahyoint [2]int = input()
	_ = zahyoint
	fmt.Println(zahyoint[0])
	fmt.Println(zahyoint[1])

}

func input() [2]int {
	var zahyostr string
	fmt.Println("input: ")
	fmt.Scanln(&zahyostr)
	x, _ := strconv.Atoi(zahyostr[0:1])
	y, _ := strconv.Atoi(zahyostr[2:3])
	var zahyoint [2]int = [2]int{x, y}
	return zahyoint
}
