package main

import (
	"fmt"

	child "github.com/bstoll/import-visibility-child"
)

func main() {
	fmt.Println("Hello from main v2")
	child.Child()
}
