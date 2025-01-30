package floattostring_test

import (
	"fmt"
	"testing"

	"github.com/xescugc/benchmarks/floattostring"
)

var (
	f float64 = 2222.222222
	s string  = "2222.222222"
)

func TestStrconv(t *testing.T) {
	fmt.Println("How do the formats look like")
	fmt.Println("b:", floattostring.Strconv(16.2, 'b'))
	fmt.Println("e:", floattostring.Strconv(17.2, 'e'))
	fmt.Println("E:", floattostring.Strconv(18.2, 'E'))
	fmt.Println("f:", floattostring.Strconv(19.2, 'f'))
	fmt.Println("g:", floattostring.Strconv(20.2, 'g'))
	fmt.Println("G:", floattostring.Strconv(21.2, 'G'))
	fmt.Println("x:", floattostring.Strconv(22.2, 'x'))
	fmt.Println("X:", floattostring.Strconv(23.2, 'X'))
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Sprintf(f)
	}
}

func BenchmarkStrconv_b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'b')
	}
}
func BenchmarkStrconv_e(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'e')
	}
}
func BenchmarkStrconv_E(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'E')
	}
}
func BenchmarkStrconv_f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'f')
	}
}
func BenchmarkStrconv_g(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'g')
	}
}
func BenchmarkStrconv_G(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'G')
	}
}
func BenchmarkStrconv_x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'x')
	}
}
func BenchmarkStrconv_X(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floattostring.Strconv(f, 'X')
	}
}
