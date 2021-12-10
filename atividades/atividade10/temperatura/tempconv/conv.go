package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit((c*9)/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit(1.8*(k-273) + 32) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273) }
