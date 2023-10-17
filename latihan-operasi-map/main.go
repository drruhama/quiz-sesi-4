package main

import "fmt"

func main() {
	fruits := map[string]int{
		"Apple":  10,
		"Banana": 15,
		"Orange": 8,
		"Grape":  20,
	}
	fmt.Println("Sebelum ditambah/hapus:")
	fmt.Println(fruits)
	fmt.Println("")

	fruits["Mango"] = 12
	fruits["Strawberry"] = 7
	fmt.Println("Setelah menambahkan Mango dan Strawberry:")
	fmt.Println(fruits)
	fmt.Println("")

	delete(fruits, "Orange")
	fmt.Println("Setelah menghapus buah Orange:")
	fmt.Println(fruits)
	fmt.Println("")

	key := "Apple"
	val, isExist := fruits[key]

	if isExist {
		fmt.Println("Jumlah apel yang tersedia adalah", val)
	} else {
		fmt.Println("Buah apel tidak ditemukan")
	}
	fmt.Println("")

	daftarBuah := []map[string]int{fruits}
	fmt.Println("Daftar buah-buahan beserta jumlahnya:")
	for _, daftarbuah := range daftarBuah {
		fmt.Println("Banana :", daftarbuah["Banana"])
		fmt.Println("Grape :", daftarbuah["Grape"])
		fmt.Println("Mango :", daftarbuah["Mango"])
		fmt.Println("Strawberry :", daftarbuah["Strawberry"])
		fmt.Println("Apple :", daftarbuah["Apple"])
		//fmt.Println(key, "	\t:", val)
	}
}
