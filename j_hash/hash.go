package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func generateSHA256Hash(data string) string {
    hash := sha256.Sum256([]byte(data))
    return fmt.Sprintf("%x", hash)
}

func compareHashes(hash1, hash2 string) bool {
    return hash1 == hash2
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("=== Program Hashing SHA-256 ===")
    fmt.Print("Masukkan data pertama: ")
    input1, _ := reader.ReadString('\n')
    input1 = strings.TrimSpace(input1)
    if input1 == "" {
        fmt.Println("Input tidak boleh kosong.")
        return
    }
    fmt.Print("Masukkan data kedua: ")
    input2, _ := reader.ReadString('\n')
    input2 = strings.TrimSpace(input2)

    if input2 == "" {
        fmt.Println("Input tidak boleh kosong.")
        return
    }
    hash1 := generateSHA256Hash(input1)
    hash2 := generateSHA256Hash(input2)
    fmt.Println("\n=== Hasil Hashing ===")
    fmt.Printf("Data 1 : %s\n", input1)
    fmt.Printf("Hash SHA-256: %s\n", hash1)
    fmt.Printf("Data 2 : %s\n", input2)
    fmt.Printf("Hash SHA-256: %s\n", hash2)
    fmt.Println("\nApakah kedua hash sama?")
    if compareHashes(hash1, hash2) {
        fmt.Println("Ya, hash sama.")
    } else {
        fmt.Println("Tidak, hash berbeda.")
    }
}
