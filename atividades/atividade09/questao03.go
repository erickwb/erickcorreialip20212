package main

import (
	"flag"
	"fmt"
)

//parametros para flag
var r = flag.Float64("r", 0.0, "Float referente ao raio")
var c = flag.Bool("c", false, "omitir o calculo da circuferencia")
var a = flag.Bool("a", false, "omitir o calculo da area")

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

	flag.Parse()

	raio = *r

	if *c {
		//se a flag c for utilizada (true) calcula o comprimento
		fmt.Printf("valor do comprimento e igual a: %f \n", calcularComprimento(valor_pi, raio))
	}
	if *a {
		//se a flag a for utilizada (true) calcula a area
		fmt.Printf("valor da área e igual a: %f \n", calcularArea(valor_pi, raio))
	}
}
