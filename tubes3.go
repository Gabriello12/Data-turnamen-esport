package main

import "fmt"

const maks int = 16
const makx int = 100

type Pemain struct {
	nama   string
	kill   int
	death  int
	assist int
}

type Tim struct {
	nama   string
	player [5]Pemain
	menang int
	kalah  int
	match  pertandingan
}

type dataTim [makx]Tim

type pertandingan struct {
	tim1    string
	tim2    string
	skor1   int
	skor2   int
	tanggal string
	mvp     string
}

type dataPertandingan [maks]pertandingan

func main() {
	var tim dataTim
	var match dataPertandingan
	var n, m, pilihan, peran, maksperan int
	var pw string
	var pwBenar bool = false
	n = 2
	m = 2
	maksperan = 5

	tim[0] = Tim{
		nama:   "EVOS",
		menang: 3,
		kalah:  1,
		player: [5]Pemain{
			{"Wann", 10, 2, 8},
			{"Rexxy", 7, 3, 9},
			{"Antimage", 5, 5, 6},
			{"Clover", 8, 4, 7},
			{"Luminaire", 2, 1, 14},
		},
	}
	tim[1] = Tim{
		nama:   "RRQ",
		menang: 2,
		kalah:  2,
		player: [5]Pemain{
			{"Vyn", 6, 4, 10},
			{"Albert", 12, 5, 4},
			{"Lemon", 4, 2, 9},
			{"Skylar", 7, 3, 8},
			{"R7", 5, 5, 5},
		},
	}
	tim[2] = Tim{
		nama:   "ONIC",
		menang: 4,
		kalah:  0,
		player: [5]Pemain{
			{"Butsss", 11, 3, 10},
			{"CW", 9, 2, 12},
			{"Sanz", 8, 4, 11},
			{"Kiboy", 6, 5, 14},
			{"Drian", 7, 3, 9},
		},
	}
	tim[3] = Tim{
		nama:   "AURA",
		menang: 1,
		kalah:  3,
		player: [5]Pemain{
			{"High", 6, 7, 5},
			{"Fluffy", 5, 6, 8},
			{"Godiva", 4, 4, 9},
			{"Facehugger", 3, 5, 7},
			{"Kabuki", 5, 6, 6},
		},
	}
	tim[4] = Tim{
		nama:   "GEEK",
		menang: 0,
		kalah:  4,
		player: [5]Pemain{
			{"Baloyskie", 3, 8, 5},
			{"Janaaqt", 4, 7, 6},
			{"Luke", 2, 6, 4},
			{"Caderaa", 5, 6, 5},
			{"Aboy", 3, 5, 7},
		},
	}

	match[0] = pertandingan{"EVOS", "RRQ", 2, 1, "2025-05-01", "Albert"}
	match[1] = pertandingan{"RRQ", "EVOS", 0, 2, "2025-05-10", "Wann"}

	fmt.Println("Peran : ")
	fmt.Println("1. Panitia")
	fmt.Println("2. Pengguna")
	fmt.Print("Pilih peran : ")
	fmt.Scan(&peran)

	if peran == 1 {
		for i := 0; i < maksperan; i++ {
			fmt.Print("Masukkan password: ")
			fmt.Scan(&pw)

			if pw == "admin1234" {
				fmt.Println("Akses diterima.")
				pwBenar = true
				i = maksperan
			} else {
				fmt.Println("Password salah.")
			}
		}
		if !pwBenar {
			fmt.Println("Percobaan habis. Akses ditolak.")
			pilihan = 9
		}
	}

	for peran == 2 && pilihan != 9 {
		fmt.Println("\nMenu Aplikasi Turnamen")
		fmt.Println("1. Tampilkan Semua Tim")
		fmt.Println("2. Cari Tim")
		fmt.Println("3. Urutkan Tim")
		fmt.Println("4. Tampilkan Jadwal Pertandingan dan Hasil")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanSemuaTim(tim, n)
		case 2:
			cariTim(tim, n)
		case 3:
			urutkanTim(&tim, n)
		case 4:
			tampilkanJadwalPertandingan(match, m)
			tampilkanHasilPertandingan(match, m)
		}
	}

	for peran == 1 && pilihan != 9 {
		fmt.Println("\nMenu Aplikasi Turnamen")
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Ubah Tim")
		fmt.Println("3. Hapus Tim")
		fmt.Println("4. Tampilkan Semua Tim")
		fmt.Println("5. Cari Tim")
		fmt.Println("6. Urutkan Tim")
		fmt.Println("7. Tampilkan Jadwal Pertandingan dan Hasil")
		fmt.Println("8. Tambahkan Jadwal Pertandingan dan Hasil")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahtim(&tim, &n)
		case 2:
			ubahtim(&tim, n)
		case 3:
			hapusTim(&tim, &n)
		case 4:
			tampilkanSemuaTim(tim, n)
		case 5:
			cariTim(tim, n)
		case 6:
			urutkanTim(&tim, n)
		case 7:
			tampilkanJadwalPertandingan(match, m)
			tampilkanHasilPertandingan(match, m)
		case 8:
			tambahPertandingan(&match, &m)
		}
	}
}

func tambahtim(tim *dataTim, n *int) {
	var tambah int
	fmt.Println("Tambah berapa tim ?")
	fmt.Scan(&tambah)

	for i := 0; i < tambah; i++ {
		fmt.Println("Masukkan nama tim : ")
		fmt.Scan(&(*tim)[*n].nama)
		fmt.Println("Total Menang : ")
		fmt.Scan(&(*tim)[*n].menang)
		fmt.Println("Total Kalah : ")
		fmt.Scan(&(*tim)[*n].kalah)
		for j := 0; j < 5; j++ {
			fmt.Println("Masukkan nama pemain : ")
			fmt.Scan(&(*tim)[*n].player[j].nama)
			fmt.Println("Masukkan Kill, Death, Assist : ")
			fmt.Scan(&(*tim)[*n].player[j].kill, &(*tim)[*n].player[j].death, &(*tim)[*n].player[j].assist)
		}
		*n++
	}
}

func ubahtim(tim *dataTim, jumlahTim int) {
	var namaCari string
	var ditemukan bool = false
	var pilihan int
	var idx int

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
				fmt.Print("Pilih pemain ke-berapa (1-5): ")
				fmt.Scan(&idx)
				idx = idx - 1
				if idx >= 0 && idx < 5 {
					fmt.Print("Nama pemain baru: ")
					fmt.Scan(&(*tim)[i].player[idx].nama)
					fmt.Print("Kill, Death, Assist baru: ")
					fmt.Scan(&(*tim)[i].player[idx].kill, &(*tim)[i].player[idx].death, &(*tim)[i].player[idx].assist)
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
	}
}

func hapusTim(tim *dataTim, jumlah *int) {
	var nama string
	var ditemukan bool = false

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
	fmt.Println("\nDaftar Semua Tim:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Tim %d: %s | Menang: %d | Kalah: %d\n", i+1, tim[i].nama, tim[i].menang, tim[i].kalah)
		for j := 0; j < 5; j++ {
			fmt.Printf("  Pemain %d: %s | K/D/A Pertandingan terbaru: %d/%d/%d\n", j+1, tim[i].player[j].nama, tim[i].player[j].kill, tim[i].player[j].death, tim[i].player[j].assist)
		}
	}
}

func cariTim(tim dataTim, jumlah int) {
	var nama string
	var ditemukan bool = false
	fmt.Print("Masukkan nama tim yang dicari: ")
	fmt.Scan(&nama)

	for i := 0; i < jumlah; i++ {
		if tim[i].nama == nama {
			ditemukan = true
			fmt.Printf("Tim ditemukan: %s | Menang: %d | Kalah: %d\n", tim[i].nama, tim[i].menang, tim[i].kalah)
			for j := 0; j < 5; j++ {
				fmt.Printf("  Pemain %d: %s | K/D/A: %d/%d/%d\n", j+1, tim[i].player[j].nama, tim[i].player[j].kill, tim[i].player[j].death, tim[i].player[j].assist)
			}
		}
	}
	if !ditemukan {
		fmt.Println("Tim tidak ditemukan.")
	}
}

func urutkanTim(tim *dataTim, jumlah int) {
	var pilihan int
	fmt.Println("1. Urutkan berdasarkan poin tertinggi")
	fmt.Println("2. Urutkan berdasarkan poin terendah")
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
	fmt.Println("\nJadwal Pertandingan:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s vs %s | Tanggal: %s\n", i+1, match[i].tim1, match[i].tim2, match[i].tanggal)
	}
}

func tampilkanHasilPertandingan(match dataPertandingan, jumlah int) {
	fmt.Println("\nHasil Pertandingan:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s %d - %d %s | MVP: %s\n", i+1, match[i].tim1, match[i].skor1, match[i].skor2, match[i].tim2, match[i].mvp)
	}
}
func tambahPertandingan(match *dataPertandingan, m *int) {
	var jumlah int
	fmt.Print("Masukkan jumlah pertandingan yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Pertandingan ke-%d\n", *m+1)
		fmt.Print("Tim 1: ")
		fmt.Scan(&(*match)[*m].tim1)
		fmt.Print("Tim 2: ")
		fmt.Scan(&(*match)[*m].tim2)
		fmt.Print("Skor Tim 1: ")
		fmt.Scan(&(*match)[*m].skor1)
		fmt.Print("Skor Tim 2: ")
		fmt.Scan(&(*match)[*m].skor2)
		fmt.Print("Tanggal (YYYY-MM-DD): ")
		fmt.Scan(&(*match)[*m].tanggal)
		fmt.Print("MVP: ")
		fmt.Scan(&(*match)[*m].mvp)
		*m++
	}
	fmt.Println("Pertandingan berhasil ditambahkan.")
}
