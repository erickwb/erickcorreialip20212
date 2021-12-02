package main

import "fmt"

var Valor_pi float64 = 22.0 / 7.0
var Raio float64 = 1.984

func main() {
	var comprimento float64 = (2 * Valor_pi * Raio)
	var area float64 = Valor_pi * (Raio * Raio)

	fmt.Printf("valor do comprimento e igual a: %f \n", comprimento)
	fmt.Printf("valor da área e igual a: %f \n", area)
}
