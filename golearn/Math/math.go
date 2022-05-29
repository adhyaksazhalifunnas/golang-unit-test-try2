package golearn

func Absolute(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func Multiply(a, b int) int {
	return a * b
}
