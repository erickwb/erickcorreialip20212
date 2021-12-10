package distconv

//convertendo pés para metro
func PToM(p Pes) Metro { return Metro(p / 3.2808) }

//convertendo metro para pés
func MToP(m Metro) Pes { return Pes(m * 3.2808) }
