package inttostring_test

import (
	"testing"

	"github.com/xescugc/benchmarks/inttostring"
)

const (
	x = 10
	y = 20
	z = 30
)

func BenchmarkStrconvItoaConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inttostring.StrconvItoaConcat(x, y, z)
	}
}

func BenchmarkStrconvItoaSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inttostring.StrconvItoaSprintf(x, y, z)
	}
}

func BenchmarkStrconvParseIntConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inttostring.StrconvParseIntConcat(x, y, z)
	}
}

func BenchmarkStrconvParseIntSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inttostring.StrconvParseIntSprintf(x, y, z)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inttostring.Sprintf(x, y, z)
	}
}
