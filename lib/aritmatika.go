package lib

func Penjumlahan(a,b int) int {
	return a + b
}

func Pengurangan(a,b int) int {
	return a - b
}

func Perkalian(a,b int) int {
	return a * b
}

func PembagianInt(a,b int) int {
	result := float64(a) / float64(b)
	return int(result)
}