package main

import (
	"fmt"

	vinyl_types "github.com/mehoggan/vinyl-collection-service-go/types"
)

func main() {
	var albums = []vinyl_types.Album{
		{ID: "1",
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99},
		{ID: "2",
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99},
		{ID: "3",
			Title:  "Sarah Vaughan and Clifford Brown",
			Artist: "Sarah Vaughan",
			Price:  39.99},
	}
	fmt.Printf("%v", albums)
}
