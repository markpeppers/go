package popcount_test

import (
	"popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(123456789)
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.LoopPopCount(123456789)
	}
}

func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ShiftPopCount(123456789)
	}
}

func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.ClearPopCount(123456789)
	}
}

