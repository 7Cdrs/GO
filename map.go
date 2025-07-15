package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alice",
		"address": "123 Main St",
	}
	person["DevOps"] = "Hazel Ray"

	fmt.Println(person)
	fmt.Println(person["name"])    // Output: Alice
	fmt.Println(person["address"]) // Output: 123 Main St

	// Perbaikan di sini
	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "John Doe"
	book["publisher"] = "Tech Books Publishing"

	fmt.Println(book)
	fmt.Println(len(book)) // Output: 3

	delete(book, "publisher")

	fmt.Println(book)       // Output: map[author:John Doe title:Go Programming]
	fmt.Println(len(book))  // Output: 2
}
