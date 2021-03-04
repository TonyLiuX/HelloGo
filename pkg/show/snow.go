package show

import (
	"fmt"
	"gitee.com/tonyliu88/go-test/pkg/calc"
)

func init() {
	fmt.Println("package show init.")
}

func Calc(x, y int) {
	add := calc.Add(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, add)
	minus := calc.Minus(x, y)
	fmt.Printf("%d - %d = %d\n", x, y, minus)
	multi := calc.Multi(x, y)
	fmt.Printf("%d * %d = %d\n", x, y, multi)
	divide := calc.Divide(x, y)
	fmt.Printf("%d / %d = %d\n", x, y, divide)
}
