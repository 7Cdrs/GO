package main

import "fmt"

func main() {
	var names = [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names[0]) // Output: Alice
	fmt.Println(names[1]) // Output: Bob
	fmt.Println(names[2]) // Output: Charlie

	var values = [3] int {
		10,
		20,
		30,

	}
	fmt.Println(values) 

	fmt.Println("Panjang array names:", len(names))
	fmt.Println("Panjang array values:", len(values)) // Output: 3
}