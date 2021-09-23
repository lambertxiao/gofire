package main

import (
	"gofire"
)

func main() {
	server := gofire.NewServer("127.0.0.1", 7777)
	err := server.Serve("")

	if err != nil {
		panic(err)
	}
}
