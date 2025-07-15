package main 

import "fmt"

func main() {
	var name1 = "Alice"
	var name2 = "Hazel"

	var result bool = name1 > name2
	fmt.Println("Apakah name1 lebih dari name2?", result)


	var value1 = 15
	var value2 = 20

	fmt.Println("Apakah value1 lebih kecil dari value2?", value1 < value2)
	fmt.Println("Apakah value1 lebih besar dari value2?", value1 > value2)	
	fmt.Println("Apakah value1 sama dengan value2?", value1 == value2)
	fmt.Println("Apakah value1 tidak sama dengan value2?", value1 != value2)
}