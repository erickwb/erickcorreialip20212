package main

import "fmt"

var Valor_pi float64 = 22.0 / 7.0
var Raio float64 = 1.984

func calcularComprimento(valor_1 float64, valor_2 float64) float64 {
	var comprimento float64 = 0.0
	comprimento = (2 * valor_1 * valor_2)
	return comprimento

}

func calcularArea(valor_1 float64, valor_2 float64) float64 {
	var area float64 = 0.0
	area = valor_1 * (valor_2 * valor_2)
	return area
}

func main() {

	fmt.Printf("valor do comprimento e igual a: %f \n", calcularComprimento(Valor_pi, Raio))
	fmt.Printf("valor da área e igual a: %f \n", calcularArea(Valor_pi, Raio))
}
