// Correção: 0,1. 
package main

import (
	"fmt"
	"os"
)

func expand(s string, f func(string) string) string {
	if len(s)-1 == 0 {
		return ""
	}

	if s[0] == '$' {
		var tamanho_s int
		tamanho_s = len(s) - 1
	}
}

func joker(s string) string {
	return "HaHAhAhAhaHA " + s
}

func main() {
	var input = os.Args[1]
	input = expand(input, joker)
	fmt.Printf(input)
}
