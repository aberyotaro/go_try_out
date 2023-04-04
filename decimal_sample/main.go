package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

func main() {
	f1 := 0.000000000000000002
	f2 := 0.000000000000000001
	f3 := 0.000000000000000001

	a := f1 * f2 * f3
	fmt.Println("float64: ", strconv.FormatFloat(a, 'f', -1, 64))
	fmt.Println("桁数: ", len(strconv.FormatFloat(a, 'f', -1, 64)))

	x, _ := decimal.NewFromString(strconv.FormatFloat(f1, 'f', -1, 64))
	y, _ := decimal.NewFromString(strconv.FormatFloat(f2, 'f', -1, 64))
	z, _ := decimal.NewFromString(strconv.FormatFloat(f3, 'f', -1, 64))

	x = x.Mul(y).Mul(z)
	fmt.Println("decimal: ", x)
	fmt.Println("桁数: ", len(x.String()))
}
