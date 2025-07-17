package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
    const jumlahSiswa = 3

    var nama [jumlahSiswa]string
    var tugas [jumlahSiswa]float64
    var uts [jumlahSiswa]float64
    var uas [jumlahSiswa]float64
    var nilaiAkhir [jumlahSiswa]float64
    var status [jumlahSiswa]string

    fmt.Println("=== INPUT DATA 3 SISWA ===")
    for i := 0; i < jumlahSiswa; i++ {
        fmt.Printf("\nSiswa ke-%d:\n", i+1)
        fmt.Print("Nama : ")
        fmt.Scanln(&nama[i])
        fmt.Print("Nilai Tugas : ")
        fmt.Scanln(&tugas[i])
        fmt.Print("Nilai UTS : ")
        fmt.Scanln(&uts[i])
        fmt.Print("Nilai UAS : ")
        fmt.Scanln(&uas[i])

        nilaiAkhir[i] = tugas[i]*0.3 + uts[i]*0.3 + uas[i]*0.4

        if nilaiAkhir[i] >= 70 {
            status[i] = "LULUS"
        } else {
            status[i] = "TIDAK LULUS"
        }
    }

    file, err := os.Create("hasil_nilai.txt")
    if err != nil {
        fmt.Println("Gagal menyimpan file:", err)
        return
    }
    defer file.Close()

    header := fmt.Sprintf("%-5s %-10s %-6s %-6s %-6s %-10s %-15s\n", 
        "No", "Nama", "Tgs", "UTS", "UAS", "Akhir", "Status")
    fmt.Println("\n=== REKAP NILAI SISWA ===")
    fmt.Print(header)
    file.WriteString("=== REKAP NILAI SISWA ===\n")
    file.WriteString(header)

    for i := 0; i < jumlahSiswa; i++ {
        baris := fmt.Sprintf("%-5d %-10s %-6.1f %-6.1f %-6.1f %-10.1f %-15s\n", i+1, nama[i], tugas[i], uts[i], uas[i], nilaiAkhir[i], status[i])
        fmt.Print(baris)
        file.WriteString(baris)
    }

    type siswaData struct {
        nama        string
        nilaiAkhir  float64
    }

    var ranking []siswaData
    for i := 0; i < jumlahSiswa; i++ {
        ranking = append(ranking, siswaData{nama: nama[i], nilaiAkhir: nilaiAkhir[i]})
    }

    sort.Slice(ranking, func(i, j int) bool {
        return ranking[i].nilaiAkhir > ranking[j].nilaiAkhir
    })

    fmt.Println("\n=== PERINGKAT SISWA ===")
    file.WriteString("\n=== PERINGKAT SISWA ===\n")
    for i, s := range ranking {
        baris := fmt.Sprintf("Peringkat %d: %-10s (%.1f)\n", i+1, s.nama, s.nilaiAkhir)
        fmt.Print(baris)
        file.WriteString(baris)
    }

    fmt.Println("\n✅ Data berhasil disimpan ke file 'hasil_nilai.txt'")
}