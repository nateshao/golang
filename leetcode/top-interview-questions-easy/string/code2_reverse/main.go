package main

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if res > (math.MaxInt32/10) || res < (math.MaxInt32/10) {
			return 0
		}
		res = res*10 + digit
	}
	return res
}
func main() {

}
