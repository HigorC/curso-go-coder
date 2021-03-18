package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("o valor maximo de int é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamaneto da tabela unicode (int32)
	fmt.Println("o rune é ", reflect.TypeOf(i2))

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é ", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "oii"
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho é ", len(s1))
	fmt.Println("o tipo é ", reflect.TypeOf(s1))

	s2 := `oi
	tudo 
	bem`
	fmt.Println(s2)

	// char?
	char := 'a'
	fmt.Println("tipo de char é", reflect.TypeOf(char))

}
