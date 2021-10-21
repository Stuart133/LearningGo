package sample

import (
	"fmt"
	"testing"
)

var gs string
var a []int

func BenchmarkSprint(b *testing.B) {
	b.Run("None", benchmarkSprint)
	b.Run("Format", benchmarkSprintf)
}

func benchmarkSprint(b *testing.B) {
	var s string

	a = append(a, b.N)
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	if len(a) > 4 {
		// fmt.Println(a)
	}

	gs = s
}

func benchmarkSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gs = s
}
