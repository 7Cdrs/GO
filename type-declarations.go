package main
import "fmt"

func main() {
	type NoKTP string
	type married bool

	var NoKtpHazel NoKTP = "1234567890"
	fmt.Println(NoKtpHazel)
	var isMarried married = false
	fmt.Println(isMarried)

}