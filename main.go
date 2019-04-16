package main

import (
	"./controller"
)

const (
	port = ":12001"
)

func main() {
	c := controller.Init()
	c.Start(port)
}
