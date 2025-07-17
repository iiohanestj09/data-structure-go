package main

import "fmt"

type Mahasiswa struct {
    Nama string
    NIM  string
}

func tampilkanNilai(data map[Mahasiswa]map[string]float64) {
    for mhs, nilaiMap := range data {
        fmt.Printf("Nama: %s | NIM: %s\n", mhs.Nama, mhs.NIM)
        for matkul, nilai := range nilaiMap {
            fmt.Printf(" - %s: %.2f\n", matkul, nilai)
        }
        fmt.Println()
    }
}

func main() {
    nilaiMahasiswa := make(map[Mahasiswa]map[string]float64)

    mhs1 := Mahasiswa{"Putra", "240535"}
    mhs2 := Mahasiswa{"Nadiyah", "240636"}

    nilaiMahasiswa[mhs1] = map[string]float64{
        "Matematika": 87.5,
        "Fisika":     78.0,
    }

    nilaiMahasiswa[mhs2] = map[string]float64{
        "Matematika": 92.0,
        "Bahasa":     88.5,
    }

    tampilkanNilai(nilaiMahasiswa)
}