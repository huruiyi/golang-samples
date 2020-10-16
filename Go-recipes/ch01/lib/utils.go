package lib

import (
	"fmt"
)

// Print all favorites
func PrintFavorites() {
	for i, v := range favorites {
		fmt.Println(i, ":", v)
	}
}
