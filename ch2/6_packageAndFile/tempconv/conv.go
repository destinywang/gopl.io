package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}