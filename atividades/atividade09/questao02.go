package main

import "fmt"

var valor_pi float64 = 22.0 / 7.0
var raio float64 = 1.984

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

	fmt.Printf("valor do comprimento e igual a: %f \n", calcularComprimento(valor_pi, raio))
	fmt.Printf("valor da área e igual a: %f \n", calcularArea(valor_pi, raio))
}
