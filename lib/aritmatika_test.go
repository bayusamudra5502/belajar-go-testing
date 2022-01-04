package lib_test

import (
	"testing"

	"github.com/bayusamudra5502/belajar-go-testing/lib"
)

func TestPenjumlahan1(t *testing.T) {
	result := lib.Penjumlahan(1,2)

	if result != 3 {
		panic("Hasil penjumlahan tidak 3")
	}
}

func TestPengurangan(t *testing.T) {
	result := lib.Pengurangan(1,1)

	if result != 0 {
		// Assert Failed
		panic("Hasil pengurangan tidak 0")
	}
}



