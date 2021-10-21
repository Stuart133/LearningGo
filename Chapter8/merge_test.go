package sample

import "testing"

var n []int

func init() {
	for i := 0; i < 1_000_000; i++ {
		n = append(n, i)
	}
}

func BenchmarkSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		single(n)
	}
}

func BenchmarkUnlimited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unlimited(n)
	}
}
func BenchmarkNumCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numCPU(n, 0)
	}
}
