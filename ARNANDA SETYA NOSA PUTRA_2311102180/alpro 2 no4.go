package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int
	for {
		var x int
		fmt.Scan(&x)

		if x == -5313541 {
			break
		}

		if x == 0 {
			sort.Ints(data)
			n := len(data)
			if n == 0 {
				continue
			}
			var median int
			if n%2 == 1 {
				median = data[n/2]
			} else {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data = append(data, x)
		}
	}
}
