// Correção: 0,5
package main

import "fmt"

var valor_pi float64 = 22.0 / 7.0
var raio float64 = 1.984

func main() {
	var comprimento float64 = (2 * valor_pi * raio)
	var area float64 = valor_pi * (raio * raio)

	fmt.Printf("valor do comprimento e igual a: %f \n", comprimento)
	fmt.Printf("valor da área e igual a: %f \n", area)
}
