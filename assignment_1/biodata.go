package main

import (
	"fmt"
	"os"
)

// Struct untuk merepresentasikan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
	Absen     int
}

// Fungsi untuk mendapatkan data teman berdasarkan absen
func getDataByAbsen(absen int) Teman {
	// Data teman-teman
	dataTeman := map[int]Teman{
		1: {"Aep", "Jl. Sudirman No. 123", "Software Engineer", "Ingin belajar pemrograman go lebih dalam", 1},
		2: {"Aceng", "Jl. Gatot Subroto No. 456", "Data Analyst", "Untuk meningkatkan kemampuan teknis", 2},
		3: {"Adul", "Jl. Pahlawan No. 789", "UI/UX Designer", "Agar dapat membuat produk digital yang lebih baik", 3},
	}

	// Mendapatkan data teman berdasarkan absen
	teman, found := dataTeman[absen]
	if !found {
		teman = Teman{}
	}
	return teman
}

func main() {
	// Memastikan argumen absen telah diberikan
	if len(os.Args) < 2 {
		fmt.Println("Cara menjalankan: go run biodata.go [nomor absen]")
		return
	}

	// Mendapatkan argumen absen dari command line
	absen := os.Args[1]

	// Konversi argumen absen ke tipe data integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mendapatkan data teman berdasarkan absen
	teman := getDataByAbsen(absenInt)

	// Menampilkan data teman
	if teman.Nama == "" {
		fmt.Println("Teman dengan absen", absenInt, "tidak ditemukan.")
	} else {
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
	}
}
