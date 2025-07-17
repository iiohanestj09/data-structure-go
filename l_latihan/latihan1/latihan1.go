// SISTEM ANTRIAN PADA RUMAH SAKIT

package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Pasien struct {
    Nama        string
    Kondisi     string
    Prioritas   int
    WaktuDatang time.Time
}

type AntrianDarurat []*Pasien

func (pq AntrianDarurat) Len() int { return len(pq) }

func (pq AntrianDarurat) Less(i, j int) bool {
    return pq[i].Prioritas < pq[j].Prioritas
}

func (pq AntrianDarurat) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *AntrianDarurat) Push(x any) {
    item := x.(*Pasien)
    *pq = append(*pq, item)
}

func (pq *AntrianDarurat) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

type AntrianReguler []*Pasien

func (q *AntrianReguler) Enqueue(p *Pasien) {
    *q = append(*q, p)
}

func (q *AntrianReguler) Dequeue() *Pasien {
    if len(*q) == 0 {
        return nil
    }
    item := (*q)[0]
    *q = (*q)[1:]
    return item
}

func (q *AntrianReguler) IsEmpty() bool {
    return len(*q) == 0
}

func LayananPasien(darurat *AntrianDarurat, reguler *AntrianReguler) {
    fmt.Println("Urutan Pelayanan Pasien:")
    for darurat.Len() > 0 || !reguler.IsEmpty() {
        if darurat.Len() > 0 {
            p := heap.Pop(darurat).(*Pasien)
            fmt.Printf("[DARURAT] %s (Prioritas: %d)\n", p.Nama, p.Prioritas)
        } else if !reguler.IsEmpty() {
            p := reguler.Dequeue()
            fmt.Printf("[REGULER] %s (Datang: %s)\n", p.Nama, p.WaktuDatang.Format("15:04:05"))
        }
    }
}

func main() {
    darurat := &AntrianDarurat{}
    reguler := &AntrianReguler{}

    heap.Init(darurat)

    heap.Push(darurat, &Pasien{Nama: "Pak Budi", Kondisi: "darurat", Prioritas: 1})
    heap.Push(darurat, &Pasien{Nama: "Bu Siti", Kondisi: "darurat", Prioritas: 2})
    heap.Push(darurat, &Pasien{Nama: "Pak Darto", Kondisi: "darurat", Prioritas: 0})

    now := time.Now()
    reguler.Enqueue(&Pasien{Nama: "Andi", Kondisi: "reguler", WaktuDatang: now.Add(1 * time.Minute)})
    reguler.Enqueue(&Pasien{Nama: "Rina", Kondisi: "reguler", WaktuDatang: now.Add(2 * time.Minute)})
    reguler.Enqueue(&Pasien{Nama: "Sari", Kondisi: "reguler", WaktuDatang: now.Add(3 * time.Minute)})

    LayananPasien(darurat, reguler)
}
