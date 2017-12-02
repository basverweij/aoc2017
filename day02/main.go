package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	d, err := readData(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(d)

	log.Printf("checksum: %d", checksum(d))
}
