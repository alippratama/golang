package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Loop utama untuk menampilkan menu dan menerima input
	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Hello, World!")
		fmt.Println("2. Kalkulator")
		fmt.Println("3. Kelola Data (Array)")
		fmt.Println("4. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Input tidak valid.")
			continue
		}
		// Pilihan menu utama menggunakan switch-case

		switch pilihan {
		case 1:
			helloWorld()
		case 2:
			kalkulator()
		case 3:
			kelolaData()
		case 4:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menampilkan "Hello, World!"
func helloWorld() {
	fmt.Println("Hello, World!")
}

// Fungsi untuk menampilkan menu kalkulator
func kalkulator() {
	for {
		fmt.Println("\nMenu Kalkulator:")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Faktorial")
		fmt.Println("6. Jumlah Banyak Angka")
		fmt.Println("7. Kembali")

		var pilihan int
		fmt.Print("Pilih operasi: ")
		fmt.Scanln(&pilihan)

		// Pilihan menu kalkulator menggunakan switch-case
		switch pilihan {
		case 1, 2, 3, 4:
			var angka1, angka2 float64
			var operator string
			fmt.Print("Masukkan angka pertama: ")
			fmt.Scanln(&angka1)
			fmt.Print("Masukkan operator (+, -, *, /): ")
			fmt.Scanln(&operator)
			fmt.Print("Masukkan angka kedua: ")
			fmt.Scanln(&angka2)

			switch operator {
			case "+":
				fmt.Printf("%.2f + %.2f = %.2f\n", angka1, angka2, angka1+angka2)
			case "-":
				fmt.Printf("%.2f - %.2f = %.2f\n", angka1, angka2, angka1-angka2)
			case "*":
				fmt.Printf("%.2f * %.2f = %.2f\n", angka1, angka2, angka1*angka2)
			case "/":
				if angka2 == 0 {
					fmt.Println("Tidak dapat membagi dengan nol")
				} else {
					fmt.Printf("%.2f / %.2f = %.2f\n", angka1, angka2, angka1/angka2)
				}

			default:
				fmt.Println("Operator tidak valid")
			}

			// Untuk menghitung faktorial
		case 5:
			var n int
			fmt.Print("Masukkan angka untuk faktorial: ")
			fmt.Scanln(&n)
			fmt.Printf("Faktorial dari %d adalah %d\n", n, factorial(n))

			// Untuk menjumlahkan beberapa angka yang dipisahkan dengan spasi
		case 6:
			var input string
			fmt.Print("Masukkan angka-angka (pisahkan dengan spasi): ")
			fmt.Scanln(&input)

			angkaStr := strings.Split(input, " ")
			var numbers []int
			for _, s := range angkaStr {
				num, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println("Input tidak valid:", s)
					continue
				}
				numbers = append(numbers, num)
			}
			fmt.Printf("Jumlah dari angka-angka tersebut adalah %d\n", sum(numbers...))
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk mengelola data array
func kelolaData() {
	var data []int
	for {
		fmt.Println("\nMenu Kelola Data:")
		fmt.Println("1. Tambah data")
		fmt.Println("2. Tampilkan data")
		fmt.Println("3. Kembali")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var angka int
			fmt.Print("Masukkan angka: ")
			fmt.Scanln(&angka)
			data = append(data, angka)
			fmt.Println("Data berhasil ditambahkan")
		case 2:
			if len(data) == 0 {
				fmt.Println("Data masih kosong")
			} else {
				fmt.Println("Data:")
				for _, nilai := range data {
					fmt.Println(nilai)
				}
			}
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi untuk menghitung faktorial dari sebuah angka
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Fungsi untuk menjumlahkan banyak angka
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
