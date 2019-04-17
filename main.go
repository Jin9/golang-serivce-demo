package main

import (
	"gitlab.com/chinnawat.w/golang-service-demo/controller"
)

const (
	port = ":12001"
)

func main() {
	c := controller.Init()
	c.Start(port)
}
