package main

import (
	"fmt"
	"errors"
)

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("%+v\n", errNotFound)

}