// Example from Everyday Go book
package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	fmt.Printf("Well done %s for having your first Go\n", name)
}
