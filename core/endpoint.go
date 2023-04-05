package core

import "fmt"

type Endpoint struct {
	Ip   string
	Port int
}

func (e Endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Ip, e.Port)
}
