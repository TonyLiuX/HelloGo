package myMath

func Sub(x, y int) int {
	return x - y
}

func Swap(x, y int) (int, int) {
	Add(x, y)
	return y, x
}
