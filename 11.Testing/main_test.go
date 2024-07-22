package main

import "testing"

func TestSumar(t *testing.T) {
	total := Sumar(3, 4)
	if total != 7 {
		t.Errorf("Sumar(3, 4) = %d; want 7", total)
	}
}

func BenchmarkSumar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sumar(3, 4)
	}
}
