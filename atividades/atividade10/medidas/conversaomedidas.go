package main

import (
	"fmt"
	"medidas/medidaconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		p := medidaconv.Pes(t)
		m := medidaconv.Metro(t)
		fmt.Printf("\n %s = %s \n %s = %s \n ", p, medidaconv.PToM(p), m, medidaconv.MToP(m))

	}
}
