package main 

import "fmt"

func main() {
	var result int = 10 + 5
	fmt.Println("Hasil penjumlahan:", result)

	var a = 10
	var b = 5
	var c =  a * b

	fmt.Println("Hasil perkalian:", c)

	var d = 20
	 d += 20

	fmt.Println("Hasil penambahan d =", d)

	d++
	fmt.Println("Hasil increment d =", d)

	var negative = -10
	var positive = +10
	fmt.Println("Nilai negatif:", negative)
	fmt.Println("Nilai positif:", positive)


}