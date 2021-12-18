package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var sum float64 = 0.0

func Integral(valor_1 float64) float64 {
	//Função a integrar
	return math.Sqrt(1 + math.Sin(valor_1)*math.Sin(valor_1))
}

func main() {

	//"pegando" os argumentos
	var1 := os.Args[1]
	var2 := os.Args[2]
	var3 := os.Args[3]

	//convertendo os argumentos para float 64
	Limite_inferior, erro1 := strconv.ParseFloat(var1, 64)
	Limite_superior, erro2 := strconv.ParseFloat(var2, 64)
	Intervalos, erro3 := strconv.ParseInt(var3, 10, 64)

	//Se a conversao ter erro,  imprima o primeiro erro e saia do programa
	if (erro1 != nil) || (erro2 != nil) || (erro3 != nil) {
		fmt.Fprintf(os.Stderr, "cf: %v\n", erro1)
		os.Exit(1)
	}
	//printando as variaveis
	fmt.Println("Limite inferior: ", Limite_inferior)
	fmt.Println("Limite superior: ", Limite_superior)
	fmt.Println("Número de Partiçõe:", Intervalos)

	//realizando o calculo
	Limite_superior = (5*Limite_inferior)/2 - Limite_superior //manipulação na variavel limite superior
	// Calculando a dimensao de cada intervalo
	h := (Limite_superior - Limite_inferior) / float64(Intervalos)

	for i := 0; i < int(Intervalos); i++ {
		sum = sum + Integral(Limite_inferior) + Integral((Limite_inferior+h))*h/2
	}
	fmt.Printf("O resultado da soma e: %g", sum)

}
