package main

import (
	"awesomeProject/server"
)

func main() {
	s := server.NewSever()

	s.Run()
}
