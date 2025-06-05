package main

import (
    "fmt"
    "strings"
	"os"
	"os/exec"
)

const maks int = 100
const makx int = 100

type Pemain struct {
    nama   string
    mvp    int
}
type DataPemain struct {
    nama string
    tim  string
    mvp  int
}

type Tim struct {
    nama       string
    player     [5]Pemain
    menang     int
    kalah      int
    match      pertandingan
    skorMenang int
    skorKalah  int
}

type dataTim [makx]Tim

type pertandingan struct {
    tim1      string
    tim2      string
    skor1     int
    skor2     int
    tanggal   string
    mvp       string
    sudahMain bool
}

type dataPertandingan [maks]pertandingan

func clearscreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func tampilkanBorder(judul string) {
    panjangJudul := len(judul) + 12
    borderAtas := "┌" + strings.Repeat("─", panjangJudul) + "┐"
    borderBawah := "└" + strings.Repeat("─", panjangJudul) + "┘"

    fmt.Println(borderAtas)
    fmt.Printf("│      %s      │\n", judul)
    fmt.Println(borderBawah) 
}

func tambahtim(tim *dataTim, n *int) {
    var tambah int
	clearscreen()
    tampilkanBorder("TAMBAH TIM")
    fmt.Println("Tambah berapa tim ?")
    fmt.Scan(&tambah)

    for i := 0; i < tambah; i++ { 
        fmt.Println("Masukkan nama tim : ")
        fmt.Scan(&(*tim)[*n].nama)
        for j := 0; j < 5; j++ {
            fmt.Println("Masukkan nama pemain : ")
            fmt.Scan(&(*tim)[*n].player[j].nama)
        }
        *n++
    }
    fmt.Println("Tim berhasil ditambahkan.")
}

func ubahtim(tim *dataTim, jumlahTim int) {
    var namaCari string
    var ditemukan bool = false
    var pilihan int
    var idx int
	clearscreen()

    tampilkanBorder("UBAH TIM")
    fmt.Print("Masukkan nama tim yang ingin diubah: ")
    fmt.Scan(&namaCari)

    for i := 0; i < jumlahTim; i++ {
        if (*tim)[i].nama == namaCari {
            ditemukan = true
            fmt.Println("Tim ditemukan. Data saat ini:")
            fmt.Println("Nama:", (*tim)[i].nama)
            fmt.Println("Menang:", (*tim)[i].menang)
            fmt.Println("Kalah:", (*tim)[i].kalah)
            fmt.Println("1. Ubah nama tim")
            fmt.Println("2. Ubah jumlah menang")
            fmt.Println("3. Ubah jumlah kalah")
            fmt.Println("4. Ubah data pemain")
            fmt.Print("Pilih yang ingin diubah: ")
            fmt.Scan(&pilihan)

            switch pilihan {
            case 1:
                fmt.Print("Masukkan nama baru: ")
                fmt.Scan(&(*tim)[i].nama)
            case 2:
                fmt.Print("Masukkan jumlah menang baru: ")
                fmt.Scan(&(*tim)[i].menang)
            case 3:
                fmt.Print("Masukkan jumlah kalah baru: ")
                fmt.Scan(&(*tim)[i].kalah)
            case 4:
				fmt.Println("Daftar pemain:")
				for j := 0; j < 5; j++ {
					fmt.Printf("%d. %s\n", j+1, (*tim)[i].player[j].nama)
				}

				fmt.Print("Pilih pemain ke-berapa (1-5): ")
				fmt.Scan(&idx)
				idx = idx - 1

				if idx >= 0 && idx < 5 {
					fmt.Print("Nama pemain baru: ")
					fmt.Scan(&(*tim)[i].player[idx].nama)
				} else {
					fmt.Println("Nomor pemain tidak valid.")
				}

            default:
                fmt.Println("Pilihan tidak valid.")
            }
        }
    }

    if !ditemukan {
        fmt.Println("Tim tidak ditemukan.")
    } else {
        fmt.Println("Tim berhasil diubah.")
    }
}

func hapusTim(tim *dataTim, jumlah *int) {
    var nama string
    var ditemukan bool = false
	clearscreen()

    tampilkanBorder("HAPUS TIM")
    fmt.Print("Masukkan nama tim yang ingin dihapus: ")
    fmt.Scan(&nama)

    i := 0
    for i < *jumlah {
        if (*tim)[i].nama == nama {
            ditemukan = true
            for j := i; j < *jumlah-1; j++ {
                (*tim)[j] = (*tim)[j+1]
            }
            *jumlah--
            fmt.Println("Tim berhasil dihapus.")
            i = *jumlah
        } else {
            i++
        }
    }

    if !ditemukan {
        fmt.Println("Tim tidak ditemukan.")
    }
}

func tampilkanSemuaTim(tim dataTim, jumlah int) {
    var pilih int
	clearscreen()
    tampilkanBorder("KLASEMEN TIM DAN POIN")
    fmt.Printf("%-5s %-20s %-6s\n", "No.", "Nama Tim", "Poin")
    fmt.Println("======================================")

    for i := 0; i < jumlah; i++ {
        poin := tim[i].menang * 3
        fmt.Printf("%-5d %-20s %-6d\n", i+1, tim[i].nama, poin)
    }

    for pilih != 3 {
        fmt.Println()
        tampilkanBorder("SUBMENU KLASEMEN")
        fmt.Println(" [1] Cari Tim")
        fmt.Println(" [2] Urutkan Tim")
        fmt.Println(" [3] Keluar")
        fmt.Print("Pilih Menu: ")
        fmt.Scan(&pilih)

        switch pilih {
        case 1:
            cariTim(tim, jumlah)
        case 2:
            urutkanTim(&tim, jumlah)
        case 3:
			clearscreen()
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

func tampilkanKlasemen(tim dataTim, jumlah int) {
	clearscreen()
    tampilkanBorder("KLASEMEN TIM")
    if jumlah == 0 {
        fmt.Println("Tidak ada tim yang tersedia.")
        return
    }

    fmt.Printf("%-20s %-6s %-6s %-6s %-6s\n", "Tim", "Main", "Menang", "Kalah", "Poin")
    fmt.Println("-----------------------------------------------------")

    for i := 0; i < jumlah; i++ {
        main := tim[i].menang + tim[i].kalah
        poin := tim[i].menang * 3
        fmt.Printf("%-20s %-6d %-6d %-6d %-6d\n", tim[i].nama, main, tim[i].menang, tim[i].kalah, poin)
    }
  
}

func cariTim(tim dataTim, jumlah int) {
    var nama string
	clearscreen()
    tampilkanBorder("CARI TIM")
    fmt.Print("Masukkan nama tim yang dicari: ")
    fmt.Scan(&nama)

    low := 0
    high := jumlah - 1
    ditemukan := false
    posisi := -1

    for low <= high && !ditemukan {
        mid := (low + high) / 2
        if tim[mid].nama == nama {
            ditemukan = true
            posisi = mid
        } else if tim[mid].nama < nama {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

   if ditemukan {
    fmt.Printf("Tim ditemukan: %s | Menang: %d | Kalah: %d | Total Main: %d\n", tim[posisi].nama, tim[posisi].menang, tim[posisi].kalah, tim[posisi].menang + tim[posisi].kalah)

    fmt.Printf("%-10s %-23s %-11s %-12s\n", "No.", "Nama Pemain", "Total MVP", "Total Main")
	fmt.Println("---------- ----------------------- ----------- ------------")
	for j := 0; j < 5; j++ {
		fmt.Printf("%-10d %-23s %-11d %-12d\n", j+1, tim[posisi].player[j].nama, tim[posisi].player[j].mvp, tim[posisi].menang+tim[posisi].kalah)
	}

    fmt.Println()
	} else {
    fmt.Println("Tim tidak ditemukan.")
}


}

func urutkanTim(tim *dataTim, jumlah int) {
    var pilihan int
	clearscreen()
    tampilkanBorder("URUTKAN TIM")
    fmt.Println(" [1] Urutkan berdasarkan poin tertinggi")
    fmt.Println(" [2] Urutkan berdasarkan poin terendah")
    fmt.Print("Pilih opsi: ")
    fmt.Scan(&pilihan)

    for i := 0; i < jumlah-1; i++ {
        idx := i
        for j := i + 1; j < jumlah; j++ {
            if (pilihan == 1 && (*tim)[j].menang > (*tim)[idx].menang) ||
                (pilihan == 2 && (*tim)[j].menang < (*tim)[idx].menang) {
                idx = j
            }
        }
        (*tim)[i], (*tim)[idx] = (*tim)[idx], (*tim)[i]
    }

    tampilkanSemuaTim(*tim, jumlah)
}

func tampilkanJadwalPertandingan(match dataPertandingan, jumlah int) {
	clearscreen()
    tampilkanBorder("JADWAL PERTANDINGAN")
    for i := 0; i < jumlah; i++ {
        fmt.Printf("%d. %s vs %s | Tanggal: %s\n", i+1, match[i].tim1, match[i].tim2, match[i].tanggal)
    }
}

func tampilkanHasilPertandingan(match dataPertandingan, jumlah int) {
	clearscreen()
    tampilkanBorder("HASIL PERTANDINGAN")
    fmt.Printf("%-4s %-15s %-5s %-5s %-15s %-10s %-12s\n", "No", "Tim 1", "Skor1", "Skor2", "Tim 2", "MVP", "Tanggal")
	fmt.Println("---- --------------- ----- ----- --------------- ---------- ------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-4d %-15s %-5d %-5d %-15s %-10s %-12s\n", i+1, match[i].tim1, match[i].skor1, match[i].skor2, match[i].tim2, match[i].mvp, match[i].tanggal)
	}

}

func tambahPertandingan(match *dataPertandingan, m *int, tim *dataTim) {
	clearscreen()
    var pilihan [100]int
    var pemenang string
    var pilih int
    var pilihMVP int
    var sudahDiproses bool = false
    jumlahPilihan := 0

    tampilkanBorder("TAMBAH HASIL PERTANDINGAN")
    fmt.Println("Daftar jadwal pertandingan yang belum dimainkan:")
    for i := 0; i < *m; i++ {
        if !(*match)[i].sudahMain {
            fmt.Printf("[%d] %s vs %s pada %s\n", jumlahPilihan+1, (*match)[i].tim1, (*match)[i].tim2, (*match)[i].tanggal)
            pilihan[jumlahPilihan] = i
            jumlahPilihan++
        }
    }

    if jumlahPilihan == 0 {
        fmt.Println("Tidak ada jadwal yang bisa dimasukkan hasilnya.")
        return
    }

    fmt.Print("Pilih nomor pertandingan yang ingin diisi hasilnya: ")
    fmt.Scan(&pilih)

    if pilih < 1 || pilih > jumlahPilihan {
        fmt.Println("Pilihan tidak valid.")
        return
    }

    idx := pilihan[pilih-1]
    fmt.Printf("Memasukkan hasil untuk: %s vs %s\n", (*match)[idx].tim1, (*match)[idx].tim2)

    fmt.Print("Skor Tim 1: ")
    fmt.Scan(&(*match)[idx].skor1)
    fmt.Print("Skor Tim 2: ")
    fmt.Scan(&(*match)[idx].skor2)

    (*match)[idx].sudahMain = true

    tim1 := (*match)[idx].tim1
    tim2 := (*match)[idx].tim2
    skor1 := (*match)[idx].skor1
    skor2 := (*match)[idx].skor2

    if skor1 > skor2 {
        pemenang = tim1
    } else if skor2 > skor1 {
        pemenang = tim2
    }

    for j := 0; j < 10; j++ {
        if tim[j].nama == tim1 {
            if skor1 > skor2 {
                tim[j].menang++
            } else if skor1 < skor2 {
                tim[j].kalah++
            }
            tim[j].skorMenang += skor1
            tim[j].skorKalah += skor2
        }
        if tim[j].nama == tim2 {
            if skor2 > skor1 {
                tim[j].menang++
            } else if skor2 < skor1 {
                tim[j].kalah++
            }
            tim[j].skorMenang += skor2
            tim[j].skorKalah += skor1
        }
    }
    if pemenang != "" {
        for j := 0; j < 10; j++ {
            if !sudahDiproses && tim[j].nama == pemenang {
                fmt.Println("Pilih MVP dari tim", pemenang)
                for k := 0; k < 5; k++ {
                    fmt.Printf("%d. %s\n", k+1, tim[j].player[k].nama)
                }

                fmt.Print("Masukkan nomor pemain MVP: ")
                fmt.Scan(&pilihMVP)

                if pilihMVP >= 1 && pilihMVP <= 5 {
                    mvpNama := tim[j].player[pilihMVP-1].nama
                    tim[j].player[pilihMVP-1].mvp++
                    (*match)[idx].mvp = mvpNama
                } else {
                    fmt.Println("Pilihan tidak valid. MVP tidak tercatat.")
                }
                sudahDiproses = true
            }
        }
    }

    fmt.Println("Hasil pertandingan berhasil ditambahkan.")
 
}

func tambahJadwalPertandingan(match *dataPertandingan, m *int) {
    var jumlah int
	clearscreen()
    tampilkanBorder("TAMBAH JADWAL PERTANDINGAN")
    fmt.Print("Masukkan jumlah jadwal pertandingan yang ingin ditambahkan: ")
    fmt.Scan(&jumlah)

    for i := 0; i < jumlah; i++ {
        fmt.Printf("Jadwal ke-%d\n", *m+1)
        fmt.Print("Tim 1: ")
        fmt.Scan(&(*match)[*m].tim1)
        fmt.Print("Tim 2: ")
        fmt.Scan(&(*match)[*m].tim2)
        fmt.Print("Tanggal (DD-MM-YYYY): ")
        fmt.Scan(&(*match)[*m].tanggal)

        (*match)[*m].skor1 = 0
        (*match)[*m].skor2 = 0
        (*match)[*m].mvp = "-"
        (*match)[*m].sudahMain = false

        *m++
    }
    fmt.Println("Jadwal pertandingan berhasil ditambahkan.")
}

func tampilkanPemain(tim dataTim, jumlahTim int) {
	clearscreen()
    var semuaPemain [makx * 5]DataPemain
    var jumlahPemain int = 0
    var pilih int

    for i := 0; i < jumlahTim; i++ {
        for j := 0; j < 5; j++ {
            semuaPemain[jumlahPemain] = DataPemain{
                nama: tim[i].player[j].nama,
                tim:  tim[i].nama,
                mvp:  tim[i].player[j].mvp,
            }
            jumlahPemain++
        }
    }

    for pilih != 3 {
        tampilkanBorder("DAFTAR PEMAIN DAN TOTAL MVP")
        fmt.Printf("%-4s %-20s %-15s %-5s\n", "No", "Nama", "Tim", "MVP")
		fmt.Println("---- -------------------- --------------- -----")
		for i := 0; i < jumlahPemain; i++ {
			fmt.Printf("%-4d %-20s %-15s %-5d\n", i+1, semuaPemain[i].nama, semuaPemain[i].tim, semuaPemain[i].mvp)
		}


        tampilkanBorder("SUBMENU PEMAIN")
        fmt.Println(" [1] Urutkan MVP Terbanyak")
        fmt.Println(" [2] Urutkan MVP Terkecil")
        fmt.Println(" [3] Kembali")
        fmt.Print("Pilih: ")
        fmt.Scan(&pilih)

        if pilih == 1 {
            for i := 0; i < jumlahPemain-1; i++ {
                maxIdx := i
                for j := i + 1; j < jumlahPemain; j++ {
                    if semuaPemain[j].mvp > semuaPemain[maxIdx].mvp {
                        maxIdx = j	
                    }
                }
                semuaPemain[i], semuaPemain[maxIdx] = semuaPemain[maxIdx], semuaPemain[i]
            }
        } else if pilih == 2 {
            for i := 0; i < jumlahPemain-1; i++ {
                minIdx := i
                for j := i + 1; j < jumlahPemain; j++ {
                    if semuaPemain[j].mvp < semuaPemain[minIdx].mvp {
                        minIdx = j
                    }
                }
                semuaPemain[i], semuaPemain[minIdx] = semuaPemain[minIdx], semuaPemain[i]
            }
        }
    }
	if pilih == 3{
		clearscreen()
	}	
}

func tampilkanStatistik(tim dataTim, jumlah int) {
    var pilih int
	clearscreen()
    for pilih != 2 {
        tampilkanBorder("STATISTIK PERFORMA TIM")
        fmt.Printf("| %-20s | %-4s | %-6s | %-6s | %-10s | %-10s | %-12s | %-6s |\n",
            "Nama Tim", "Main", "Menang", "Kalah", "Skor Menang", "Skor Kalah", "Selisih Skor", "Poin")
        fmt.Println("|----------------------|------|--------|--------|-------------|------------|--------------|--------|")

        for i := 0; i < jumlah; i++ {
            main := tim[i].menang + tim[i].kalah
            selisih := tim[i].skorMenang - tim[i].skorKalah
            poin := tim[i].menang * 3

            fmt.Printf("| %-20s | %-4d | %-6d | %-6d | %-10d | %-10d | %-12d | %-6d |\n",
                tim[i].nama, main, tim[i].menang, tim[i].kalah, tim[i].skorMenang, tim[i].skorKalah, selisih, poin)
        }

        tampilkanBorder("SUBMENU STATISTIK")
        fmt.Println(" [1] Urutkan Statistik")
        fmt.Println(" [2] Keluar")
        fmt.Print("Pilih Menu: ")
        fmt.Scan(&pilih)

        switch pilih {
        case 1:
            urutkanTim(&tim, jumlah)
        case 2:
			clearscreen()
            fmt.Println("Keluar dari menu statistik.")
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

func main() {
    var tim dataTim
    var match dataPertandingan
    var n, m, pilihan, peran, maksperan int
    var pw string
    var pwBenar bool = false
	var pilih int
    n = 5
    m = 6
    maksperan = 5

    tim[0] = Tim{
        nama:   "EVOS",
        menang: 3,
        kalah:  0,
        player: [5]Pemain{
            {"Wann", 2},
            {"Rexxy", 0},
            {"Antimage", 1},
            {"Clover", 0},
            {"Luminaire", 0},
        },
    }
    tim[1] = Tim{
        nama:   "NAVI",
        menang: 0,
        kalah:  2,
        player: [5]Pemain{
            {"Vyn", 0},
            {"Albert", 0},
            {"Lemon", 0},
            {"Skylar", 0},
            {"R7", 0},
        },
    }
    tim[2] = Tim{
        nama:   "ONIC",
        menang: 2,
        kalah:  1,
        player: [5]Pemain{
            {"Butsss", 0},
            {"Cw", 1},
            {"Sanz", 0},
            {"Kiboy", 1},
            {"Drian", 0},
        },
    }
    tim[3] = Tim{
        nama:   "AURA",
        menang: 1,
        kalah:  1,
        player: [5]Pemain{
            {"High", 0},
            {"Fluffy", 0},
            {"Godiva", 1},
            {"Facehugger", 0},
            {"Kabuki", 0},
        },
    }
    tim[4] = Tim{
        nama:   "GEEK",
        menang: 0,
        kalah:  2,
        player: [5]Pemain{
            {"Baloyskie", 0},
            {"Janaaqt", 0},
            {"Luke", 0},
            {"Caderaa", 0},
            {"Aboy", 0},
        },
    }

    match[0] = pertandingan{"EVOS", "NAVI", 2, 1, "01-05-2025", "Wann", true}
    match[1] = pertandingan{"ONIC", "EVOS", 0, 2, "02-05-2025", "Wann", true}
    match[2] = pertandingan{"AURA", "EVOS", 0, 2, "03-05-2025", "Antimage", true}
    match[3] = pertandingan{"ONIC", "GEEK", 2, 0, "04-05-2025", "Kiboy", true}
    match[4] = pertandingan{"NAVI", "ONIC", 0, 2, "05-05-2025", "Cw", true}
    match[5] = pertandingan{"GEEK", "AURA", 0, 2, "06-05-2025", "Godiva", true}

    tampilkanBorder("REGISTER")
    fmt.Println(" [1] Admin")
    fmt.Println(" [2] Pengguna")
    fmt.Print("Pilih peran: ")
    fmt.Scan(&peran)

    for peran != 1 && peran != 2 {
        fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.")
        fmt.Print("Pilih peran: ")
        fmt.Scan(&peran)
    }

    if peran == 1 {
        fmt.Println("Anda masuk sebagai Admin")
    } else {
        fmt.Println("Anda masuk sebagai Pengguna")
        pilihan = 1
    }

    if peran == 1 {
        for i := 0; i < maksperan; i++ {
            tampilkanBorder("LOGIN ADMIN")
            fmt.Print("Masukkan password: ")
            fmt.Scan(&pw)

            if pw == "admin1234" {
                fmt.Println("Akses diterima.")
                pwBenar = true
                i = maksperan
                pilihan = 1
            } else {
                fmt.Println("Password salah.")
            }
        }
        if !pwBenar {
            fmt.Println("Percobaan habis. Akses ditolak.")
            pilihan = 0
        }
    }

    for peran == 2 && pilihan != 0 {
    tampilkanBorder("MENU APLIKASI TURNAMEN")

    fmt.Println("-----------------------------------------------------")
    fmt.Println("| No  | Menu                                        |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  0  | Keluar                                      |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  1  | Tampilkan Semua Tim                         |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  2  | Tampilkan Klasemen                          |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  3  | Cari Tim                                    |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  4  | Tampilkan Jadwal Pertandingan               |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  5  | Tampilkan Hasil Pertandingan                |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  6  | Tampilkan Pemain                            |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  7  | Tampilkan Statistik                         |")
    fmt.Println("-----------------------------------------------------")

    fmt.Print("Pilih Menu: ")
    fmt.Scan(&pilih)

		switch pilih {
		case 0:
			pilihan = 0
		case 1:
			tampilkanSemuaTim(tim, n)
		case 2:
			tampilkanKlasemen(tim, n)
		case 3:
			cariTim(tim, n)
		case 4:
		}
	}	

    for peran == 1 && pilihan != 0 {
    tampilkanBorder("MENU APLIKASI TURNAMEN")

    fmt.Println("-----------------------------------------------------")
    fmt.Println("| No  | Menu                                        |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  0  | Keluar                                      |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  1  | Tambah Tim                                  |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  2  | Ubah Tim                                    |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  3  | Hapus Tim                                   |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  4  | Tampilkan Klasemen                          |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  5  | Cari Tim                                    |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  6  | Tampilkan Hasil Pertandingan                |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  7  | Tampilkan Jadwal Pertandingan               |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  8  | Tambahkan Jadwal Pertandingan               |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("|  9  | Tambahkan Hasil Pertandingan                |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("| 10  | Tampilkan Pemain                            |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("| 11  | Tampilkan Semua Tim                         |")
    fmt.Println("-----------------------------------------------------")
    fmt.Println("| 12  | Tampilkan Statistik                         |")
    fmt.Println("-----------------------------------------------------")

    fmt.Print("Pilih Menu: ")
    fmt.Scan(&pilih)

		switch pilih {
		case 0:
				pilihan = 0
		case 1:
				tambahtim(&tim, &n)
		case 2:
				ubahtim(&tim, n)
		case 3:
				hapusTim(&tim, &n)
		case 4:
				tampilkanKlasemen(tim, n)
		case 5:
				cariTim(tim, n)
		case 6:
				tampilkanHasilPertandingan(match, m)
		case 7:
				tampilkanJadwalPertandingan(match, m)
		case 8:
				tambahJadwalPertandingan(&match, &m)
		case 9:
				tambahPertandingan(&match, &m, &tim)
		case 10:
				tampilkanPemain(tim, n)
		case 11:
				tampilkanSemuaTim(tim, n)
		case 12:
				tampilkanStatistik(tim, m)
		default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		
		}
	}	
}
