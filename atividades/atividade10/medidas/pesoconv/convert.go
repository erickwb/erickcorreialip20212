package pesoconv

//convertendo Libras para Kilo
func LToK(l Libra) Kilo { return Kilo(l / 2.2046) }

//convertendo Kilo para Libras
func KToL(k Kilo) Libra { return Libra(k * 2.2046) }
