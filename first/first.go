package first

import "fmt"

func padezhi(n int) (int, string) {
	ans := n % 100
	if ans >= 5 && ans <= 20 {
		return n, "компьюьтеров"
	}
	ans %= 10
	if ans == 1 {
		return n, "компьютер"
	}
	if ans == 2 || ans == 4 {
		return n, "компьютера"
	}
	return n, "компьюьтеров"
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(padezhi(n))

}
