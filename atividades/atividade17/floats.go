package main

import (
	"fmt"
	"os"
	"strconv"
)

func floats(floats []float64) ([]float64, bool, error) {

	var auxiliar []float64
	var erro error
	erro = nil
	var verifica_value bool = true

	for index, _ := range floats {
		if floats[index] < 0 {
			return auxiliar, verifica_value, fmt.Errorf("Números negativos não são aceitos!")
		}

		if floats[index] > 10 {
			verifica_value = false
		}

		auxiliar = append(auxiliar, floats[index]*2.0)
	}

	return auxiliar, verifica_value, erro
}

func main() {
	var input []float64

	for _, a := range os.Args[1:] {
		a, err := strconv.ParseFloat(a, 64)
		if err == nil {
			input = append(input, a)
		}
	}

	output, verifica, err := floats(input)

	if err == nil {
		fmt.Println("Nova fatia:", output)
		if verifica {
			fmt.Println("Todos os valores passados são menores que 10.")
		}
	} else {
		fmt.Println(err)
	}
}
