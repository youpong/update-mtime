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
	// updateTime()
	fname := "sample.jpeg"
	tm := readTakenDate(fname)
	updateTime(fname, tm)
}

func updateTime(path string, mtime time.Time) {
	if err := os.Chtimes(path, time.Time{}, mtime); err != nil {
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

func readTakenDate(path string) time.Time {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	tm, _ := x.DateTime()

	return tm
}
