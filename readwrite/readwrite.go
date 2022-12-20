package main

import (
	"log"
	"os"
)

func main() {
	file_name := "file.txt"
	data, error := os.ReadFile(file_name)

	if error != nil {
		log.Fatal(error)
	}
	os.Stdout.Write(data)

	err := os.WriteFile(file_name, []byte("im writing this using functions!"), 0666)

	if err != nil {
		log.Fatal(err)
	}

}
