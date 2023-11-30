package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func ExampleRead() {
	fname := "sample.jpeg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)
}

func main() {
	ExampleRead()
}
