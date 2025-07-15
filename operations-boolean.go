package main 

import "fmt"

func main() {

	var Ujian  = 80
	var Absensi = 70

	var LulusUjian = Ujian > 75
	var LulusAbsensi = Absensi > 70

	var Lulus = LulusUjian && LulusAbsensi
	fmt.Println("Apakah siswa lulus ujian?", LulusUjian)
	fmt.Println("Apakah siswa lulus absensi?", LulusAbsensi)
	fmt.Println("Apakah siswa lulus semua syarat?", Lulus)


}