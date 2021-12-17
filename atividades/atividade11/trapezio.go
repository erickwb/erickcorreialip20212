package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//"pegando" os argumentos
	var1 := os.Args[1]
	var2 := os.Args[2]
	var3 := os.Args[3]

	//convertendo os argumentos para float 64
	Limite_inferior, erro1 := strconv.ParseFloat(var1, 64)
	Limite_superior, erro2 := strconv.ParseFloat(var2, 64)
	Intervalos, erro3 := strconv.ParseFloat(var3, 64)

	//Se a conversao ter erro,  imprima o primeiro erro e saia do programa
	if (erro1 != nil) || (erro2 != nil) || (erro3 != nil) {
		fmt.Fprintf(os.Stderr, "cf: %v\n", erro1)
		os.Exit(1)
	}
	//printando as variaveis
	fmt.Println("Limite inferior: ", Limite_inferior)
	fmt.Println("Limite superior: ", Limite_superior)
	fmt.Println("Número de Partiçõe:", Intervalos)

}
