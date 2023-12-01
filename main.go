package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// scan folders recursively
//   read created date from exif
//   update file created date
// dry run/ run

func main() {
	/*
		if err := filepath.Walk(".", traverse); err != nil {
			fmt.Println(err)
		}
	*/
	updateTime()
}

func updateTime() {
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("sample.jpeg", atime, mtime); err != nil {
		log.Fatal(err)
	}
}

func traverse(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Printf(" dir : %s\n", path)
		return nil
	}
	fmt.Printf("file : %s\n", path)
	return nil
}

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
