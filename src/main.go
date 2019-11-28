package main

import (
	"demo"
	"demo2"
)

func main() {
	demo.Main()
	demo2.Test()
	demo2.Methods()
	go demo2.Say("world")
	demo2.Say("hello")
}
