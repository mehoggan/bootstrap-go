package main

import (
	"fmt"

	stringutil "github.com/mehoggan/stringutil-go"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
	fmt.Println(stringutil.ToLower("Hello"))
}
