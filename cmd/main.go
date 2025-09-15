package main

import (
	"BookStore/config"
	"BookStore/routers"
)

func main() {
	config.CreateFile()
	routers.Run()
}