package main

import (
	"fmt"
	"ripflux/core/server"
)

func main() {

	server.ServerStart()

	fmt.Println("This part shouldn't be reached")

}
