package medidaconv

//convertendo pés para metro
func PToM(p Pes) Metro { return (p / 3.2808) }

//convertendo metro para pés
func MToP(m Metro) Pes { return (m * 3.2808) }
