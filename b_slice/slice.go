package main

import "fmt"

type Buah struct {
    Nama  string
    Warna string
    Berat float64
}

func tampilkanBuah(buahSlice []Buah) {
    fmt.Println("Daftar Buah:")
    for i, buah := range buahSlice {
        fmt.Printf("%d. %s (Warna: %s, Berat: %.2f g)\n", i+1, buah.Nama, buah.Warna, buah.Berat)
    }
    fmt.Printf("Panjang Slice : %d\n", len(buahSlice))
    fmt.Printf("Kapasitas Slice: %d\n\n", cap(buahSlice))
}

func hapusBuah(slice []Buah, nama string) []Buah {
    for i, b := range slice {
        if b.Nama == nama {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    fmt.Println("Buah tidak ditemukan:", nama)
    return slice
}

func main() {
    buah := []Buah{
        {"Apel", "Merah", 150.0},
        {"Jeruk", "Oranye", 200.5},
        {"Mangga", "Kuning", 300.0},
    }

    tampilkanBuah(buah)
    buahBaru := Buah{"Pisang", "Kuning", 120.0}
    buah = append(buah, buahBaru)
    fmt.Println("Setelah Menambahkan Pisang:")
    tampilkanBuah(buah)
    buah = hapusBuah(buah, "Jeruk")
    fmt.Println("Setelah Menghapus Jeruk:")
    tampilkanBuah(buah)
}