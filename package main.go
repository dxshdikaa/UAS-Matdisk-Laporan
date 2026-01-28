package main

import "fmt"

// Struktur data mahasiswa
type Mahasiswa struct {
	NIM  string
	Nama string
}

// Fungsi Sequential Search
func sequentialSearch(data []Mahasiswa, nim string) int {
	for i, mhs := range data {
		if mhs.NIM == nim {
			return i
		}
	}
	return -1
}

func main() {
	// Data mahasiswa
	mahasiswa := []Mahasiswa{
		{NIM: "251552010030", Nama: "RIFQI"},
		{NIM: "251552010044", Nama: "DIKA"},
	}

	var nimDicari string
	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scanln(&nimDicari)

	// Cari mahasiswa
	index := sequentialSearch(mahasiswa, nimDicari)

	// Tampilkan hasil
	if index != -1 {
		fmt.Printf("Mahasiswa ditemukan: %s\n", mahasiswa[index].Nama)
	} else {
		fmt.Println("Mahasiswa tidak ditemukan")
	}
}