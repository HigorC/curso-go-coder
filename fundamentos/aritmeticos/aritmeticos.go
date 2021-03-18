package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma - ", a+b)
	fmt.Println("Subtração - ", a-b)
	fmt.Println("divisao - ", a/b)
	fmt.Println("multiplicadao - ", a*b)
	fmt.Println("modulo - ", a%b)

	//bitwise
	fmt.Println("e - ", a&b)
	fmt.Println("ou - ", a|b)
	fmt.Println("xor - ", a^b)

	c := 3.0
	d := 2.0

	//operaçoes com math
	fmt.Println("maior - ", math.Max(c, d))
	fmt.Println("menor - ", math.Min(c, d))
	fmt.Println("exponencial - ", math.Pow(c, d))
}
