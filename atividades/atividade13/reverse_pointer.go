package main

import "fmt"

//criando array 5 posição
var vet [5]int

//função reverse
func reverse(vet *[5]int) {
	for i, j := 0, (5 - 1); i < j; i, j = i+1, j-1 {
		vet[i], vet[j] = vet[j], vet[i]
	}

}

func main() {

	vet = [5]int{1, 2, 3, 4, 5} //inicializando valores

	//exibe os valores
	fmt.Println(vet)

	//chamado a função
	reverse(&vet)
	fmt.Println(vet) //exibindo valores invertidos

}
