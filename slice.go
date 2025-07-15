package main

import "fmt"

func main() {
	var months = [...]string{ // gunakan [...] agar panjang array otomatis
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[0:3] // Mengambil elemen Januari, Februari, Maret
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Panjang slice 1:", len(slice1))
	fmt.Println("Kapasitas slice 1:", cap(slice1)) // Output: 12

	slice1[0] = "Bulan Pertama"
	fmt.Println("Array months setelah diubah melalui slice1:", months)

	var slice2 = months[10:] // Mengambil elemen November, Desember
	fmt.Println("Slice 2:", slice2)

	var slice3 = append(slice2, "Hazel")
	fmt.Println("Slice 3:", slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hazel"
	newSlice[1] = "Rayes"

	fmt.Println("New Slice:", newSlice)
	fmt.Println("Panjang New Slice:", len(newSlice))
	fmt.Println("Kapasitas New Slice:", cap(newSlice))

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Array iniArray:", iniArray)
	fmt.Println("Slice iniSlice:", iniSlice)
}
