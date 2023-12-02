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
	tm := readTakenDate()
	updateTime(tm)
}

func updateTime(mtime time.Time) {
	if err := os.Chtimes("sample.jpeg", time.Time{}, mtime); err != nil {
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

func readTakenDate() time.Time {
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

	return tm
}
