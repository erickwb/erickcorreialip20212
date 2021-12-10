package main

import (
	"fmt"
	"medidas/distconv"
	"medidas/pesoconv"
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
		p := distconv.Pes(t)
		m := distconv.Metro(t)
		l := pesoconv.Libra(t)
		k := pesoconv.Kilo(t)

		fmt.Printf("\n Conversão unidades de distancia \n %s = %s \n %s = %s \n ", p, distconv.PToM(p), m, distconv.MToP(m))

		fmt.Printf("\n Conversão unidades de peso \n %s = %s \n %s = %s \n ", l, pesoconv.LToK(l), k, pesoconv.KToL(k))

	}
}
