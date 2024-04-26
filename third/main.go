package main

import "fmt"

func isPrime(n int) bool {

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func solution(a, b int) []int {
	var ans []int
	for i := a; i < b; i++ {
		if isPrime(i) == true {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := solution(a, b)

	fmt.Println(ans)
}
