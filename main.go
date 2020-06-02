package main

import (
	"fmt"
	"github.com/dionomusuko/GrassChecker/encoder"
)

func main() {
	fmt.Println(encoder.JsonEncoder("https://api.github.com/users/dionomusuko/repos"))
}
