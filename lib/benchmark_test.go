package lib_test

import (
	"testing"

	"github.com/bayusamudra5502/belajar-go-testing/lib"
)

func BenchmarkPenjumlahan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.Penjumlahan(1,1)
	}
}

func BenchmarkPengurangan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.Pengurangan(1,1)
	}
}

func BenchmarkPerkalian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.Perkalian(1,1)
	}
}

func BenchmarkPembagian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.PembagianInt(1,1)
	}
}