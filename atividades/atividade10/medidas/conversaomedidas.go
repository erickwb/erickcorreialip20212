package main

import (
	"fmt"
	"medidas/medidasconv"
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
		p := medidasconv.Pes(t)
		m := medidasconv.Metro(t)
		fmt.Printf("\n %s = %s \n %s = %s \n ", p, medidasconv.PToM(p), m, medidasconv.MToP(m))

	}
}
