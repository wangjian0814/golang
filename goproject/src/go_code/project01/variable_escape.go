package main

import (
	"fmt"
)

type Date struct {
}

func dummy() *Date {
	var c Date
	return &c
}

func main() {
	fmt.Println(dummy())
}
