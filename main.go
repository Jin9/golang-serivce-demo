package main

import (
	"./controller"
)

func main() {
	c := controller.Init()
	c.Start(":12001")
}
