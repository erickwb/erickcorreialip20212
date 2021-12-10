package medidaconv

import "fmt"

type Pes float64
type Metro float64

func (p Pes) String() string   { return fmt.Sprintf("%g pes", p) }
func (m Metro) String() string { return fmt.Sprintf("%g metro", m) }
