package main

import "fmt"

type OrangTua struct {
	Nama string
	Umur int
}

type Siswa struct {
	Nama     string
	Umur     int
	Kelas    string
	Orangtua OrangTua
}

type Slicesiswa []Siswa

func main() {
	siswa1 := Siswa{
		Nama:  "Ali",
		Umur:  19,
		Kelas: "9A",
		Orangtua: OrangTua{
			Nama: "Budi",
			Umur: 40,
		}}

	siswa2 := Siswa{
		Nama:  "David",
		Umur:  14,
		Kelas: "8B",
		Orangtua: OrangTua{
			Nama: "Citra",
			Umur: 39,
		}}

	fmt.Println("Informasi Siswa 1:")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s \n", siswa1.Nama, siswa1.Umur, siswa1.Kelas)
	fmt.Printf("Orang Tua : %s , Umur : %d\n", siswa1.Orangtua.Nama, siswa1.Orangtua.Umur)
	fmt.Println("")
	fmt.Println("Informasi Siswa 2:")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s \n", siswa2.Nama, siswa2.Umur, siswa2.Kelas)
	fmt.Printf("Orang Tua : %s , Umur : %d\n", siswa2.Orangtua.Nama, siswa2.Orangtua.Umur)
	fmt.Println("")

	var siswa3 = struct {
		Nama  string
		Umur  int
		Kelas string
		OrangTua
	}{}
	siswa3.Nama = "Fina"
	siswa3.Umur = 16
	siswa3.Kelas = "10C"
	siswa3.OrangTua = OrangTua{"Eva", 35}
	fmt.Println("Informasi Siswa 3:")
	fmt.Printf("Nama : %s , Umur : %d , Kelas : %s \n", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
	fmt.Printf("Orang Tua : %s , Umur : %d\n", siswa3.OrangTua.Nama, siswa3.OrangTua.Umur)
	fmt.Println("")

	var daftarSiswa Slicesiswa
	daftarSiswa = append(daftarSiswa, Siswa{
		Nama:  "Eva",
		Umur:  12,
		Kelas: "6B",
		Orangtua: OrangTua{
			Nama: "Rudi",
			Umur: 40,
		}})
	daftarSiswa = append(daftarSiswa, Siswa{
		Nama:  "Faisal",
		Umur:  11,
		Kelas: "5A",
		Orangtua: OrangTua{
			Nama: "Dewi",
			Umur: 38,
		}})
	daftarSiswa = append(daftarSiswa, Siswa{
		Nama:  "Grace",
		Umur:  10,
		Kelas: "EC",
		Orangtua: OrangTua{
			Nama: "Hendro",
			Umur: 42,
		}})

	fmt.Println("Daftar Siswa:")
	for _, daftarsiswa := range daftarSiswa {
		fmt.Printf("Nama Siswa : %s , Umur : %d , Kelas : %s \n", daftarsiswa.Nama, daftarsiswa.Umur, daftarsiswa.Kelas)
		fmt.Printf("Orang Tua : %s , Umur : %d\n", daftarsiswa.Orangtua.Nama, daftarsiswa.Orangtua.Umur)
		fmt.Println("")
	}
}
