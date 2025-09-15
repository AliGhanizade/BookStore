package main

import (
	"BookStore/config"
	"BookStore/model"
	"BookStore/routers"
)

func main() {
	config.CreateFile()
	model.GetAll()
	routers.Run()
	
}