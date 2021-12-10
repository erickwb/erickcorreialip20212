package pesoconv

import "fmt"

type Libra float64
type Kilo float64

func (l Libra) String() string { return fmt.Sprintf("%g Libras", l) }
func (k Kilo) String() string  { return fmt.Sprintf("%g Kilos", k) }
