package main

import (
	"fmt"
	"github.com/dionomusuko/GrassChecker/encoder"
)

func main() {
	url := "https://api.github.com/users/dionomusuko/repos"
	fmt.Println(encoder.JsonEncoder(url))
}
