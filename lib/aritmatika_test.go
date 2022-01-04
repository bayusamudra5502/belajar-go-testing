package lib_test

import (
	"fmt"
	"testing"

	"github.com/bayusamudra5502/belajar-go-testing/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPenjumlahan1(t *testing.T) {
	result := lib.Penjumlahan(1,2)

	if result != 3 {
		// Assert Failed - Pake panic, not reccomended
		panic("Hasil penjumlahan tidak 3")
	}
}

func TestPengurangan(t *testing.T) {
	result := lib.Pengurangan(1,1)

	if result != 0 {
		// Assert Failed
		t.Fail() // Ini gagal tapi masih melanjutkan ke kode selanjutnya
		fmt.Println("HAII") // BUktinya ini keliatan di console
		t.FailNow() // Gagal total, next ke unit test selanjutnya.
		fmt.Println("Ga keliatan") // Ini ga keliatan
	}
}

func TestPerkalian(t *testing.T){
	result := lib.Perkalian(2,3)

	if result != 6 {
		t.Error("Error bos!") // Setara dengan fail cuma masih lanjut kek Fail
		fmt.Println("INi keliatan")
		t.Fatal("Ckckckc..") // Fatal menyebabkan kode selanjutnya gaakan diproses. Mirip Failnow
		fmt.Println("Gak keliatan")
	}
}

func TestPembagian(t *testing.T){
	assert.Equal(t, 2, lib.PembagianInt(4,2))
	assert.Equal(t, "Kucing", "Kucing") // Klo gagal manggil fail
	assert.Condition(t, func() (success bool) {
		success = true
		return success
	}, "Kondisi")

	require.Equal(t,1,1) // INi klo gagal manggil failNow
}

func TestSkipdeh(t *testing.T){
	if true {
		t.Skip("Lagi mau nyekip dulu")

		// t.SkipNow() -> Tanpa message yang dikembalikan
	}

	assert.Equal(t,1,1)
}