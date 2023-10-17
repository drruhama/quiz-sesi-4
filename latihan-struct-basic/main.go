package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {

	book1 := Book{"Pemrograman Go", "John Doe"}
	book2 := Book{"Algoritma dan Struktur Data", "Jane Smith"}
	fmt.Println("Informasi tentang Book 1:")
	fmt.Println("Judul :", book1.Title)
	fmt.Println("Penulis :", book1.Author)
	fmt.Println("")
	fmt.Println("Informasi tentang Book 2:")
	fmt.Println("Judul :", book2.Title)
	fmt.Println("Penulis :", book2.Author)
	fmt.Println("")

	book3 := struct {
		Title  string
		Author string
	}{
		Title:  "Machine Learning for Beginners",
		Author: "David Johnson",
	}
	fmt.Println("Informasi tentang Book 3:")
	fmt.Println("Judul :", book3.Title)
	fmt.Println("Penulis :", book3.Author)
	fmt.Println("")

}
