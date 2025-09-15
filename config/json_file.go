package config

import "os"

var FilePath string = "../books.json"

func CreateFile() {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
