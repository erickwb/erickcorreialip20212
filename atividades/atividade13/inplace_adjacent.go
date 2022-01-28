// Correção: 0,1.
package main

import "fmt"

//criando array de String
var palavras [10]string = [...]string{0: "aaa",
	1: "aaa",
	2: "aaa",
	3: "aaa",
	4: "ddd",
	5: "bbb",
	6: "ccc",
	7: "ccc",
	8: "eee",
	9: "fff",
}

//função in-place
func in_place(vet []string) {

	for i:=0; i<len(vet); i++{
		if(vet[i] == vet[i+1])
		

	}
	fmt.Println(vet)
}

func main() {

	fmt.Println(palavras)
	in_place(palavras[:])

}
