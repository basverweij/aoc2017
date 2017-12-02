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

	log.Printf("checksum 1: %d", checksum(d))

	log.Printf("checksum 2: %d", checksum2(d))
}
