package main

import (
	"gofire"
)

func main() {
	client := gofire.NewClient("127.0.0.1:7777")
	client.Send([]byte("hello"))
}
