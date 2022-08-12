package main

//gofiber.oi

import (
	"fmt"

	gosayhello "github.com/aftafiandi26/golang-go-say-hello/v2"
)

func main() {
	fmt.Println(gosayhello.SayHello("Dede"))
    fmt.Println(gosayhello.Name())
}
