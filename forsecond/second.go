package main

import "fmt"

var ans map[int]int

func solution(n int) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			ans[i] += 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Initialize the ans map
	ans = make(map[int]int)

	for i := 0; i < n; i++ {
		solution(arr[i])
	}

	var sol []int

	for k, v := range ans {
		if v >= 2 {
			sol = append(sol, k)
		}
	}

	// Print the values from sol
	for _, val := range sol {
		fmt.Print(val, " ")
	}
}
