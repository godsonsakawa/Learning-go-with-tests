package iteration

// Repeat returns character repeated 5 times.
func Repeat(chr string) string {
	var Repeated string
	for i :=0; i < 5; i++ {
		Repeated += chr
	}
	return Repeated
}
