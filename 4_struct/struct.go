package main

import "fmt"


type Mahasiswa struct {
    Nama      string
    NIM       string
    Jurusan   string
    NilaiUAS  float64
    NilaiUTS  float64
}

func (m Mahasiswa) HitungRataRata() float64 {
    return (m.NilaiUAS + m.NilaiUTS) / 2
}

func (m *Mahasiswa) UbahNama(namaBaru string) {
    m.Nama = namaBaru
}

func CetakProfil(m Mahasiswa) {
    fmt.Println("== Data Mahasiswa ==")
    fmt.Println("Nama :", m.Nama)
    fmt.Println("NIM :", m.NIM)
    fmt.Println("Jurusan:", m.Jurusan)
}

func BuatMahasiswa(nama, nim, jurusan string, uts, uas float64) Mahasiswa {
    return Mahasiswa{
        Nama:     nama,
        NIM:      nim,
        Jurusan:  jurusan,
        NilaiUTS: uts,
        NilaiUAS: uas,
    }
}

func main() {
    mhs1 := BuatMahasiswa("Aril Salwa Hartono", "TI2025005", "Teknik Informatika", 85, 90)

    CetakProfil(mhs1)
    fmt.Printf("Rata-rata nilai: %.2f\n", mhs1.HitungRataRata())

    mhs1.UbahNama("Aril Salwa H.")

    fmt.Println("\nSetelah perubahan nama:")
    CetakProfil(mhs1)
}