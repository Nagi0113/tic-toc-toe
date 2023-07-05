package main

import (
//"fmt"
)

func judge(banmen [3][3]int) string {
	var tmp int
	for i := 0; i < 3; i++ {
		tmp = 0
		for j := 0; j < 3; j++ {
			tmp = tmp + banmen[i][j]
			if tmp == 3 {
				return "O win"
			} else if tmp == 12 {
				return "X win"
			}
		}
	}
	return "draw"
}

//○：1
//x：4
//.：0
