package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	//fmt.Print(" ")
	for i := 1; i <= n+1; i++ {
		//fmt.Print(i-1, " ")
		for j := 1; j <= n+1; j++ {
			if i == 1 {
				if (j - 1) == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(j-1, " ")
				}
			} else if j == 1 {
				if (i - 1) == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(i-1, " ")
				}
			} else {
				fmt.Print((i-1)*(j-1), " ")
			}
		}
		fmt.Println()
	}
}
