package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// dry run/ run

func main() {
	f, err := os.OpenFile("update-mtime.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	if err := filepath.Walk(".", traverse); err != nil {
		fmt.Println(err)
	}
}

// FileInfo.ModTime time.Time
func traverse(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	suffixes := []string{"jpeg", "jpg"}

	for _, s := range suffixes {
		if strings.HasSuffix(info.Name(), s) {
			taken := readTakenDate(path)
			updateModTime(path, taken)
			log.Printf("%s: %v, %v", path, taken.Format(time.RFC3339), info.ModTime().Format(time.RFC3339))
			break
		}
	}

	return nil
}

func updateModTime(path string, t time.Time) {
	if err := os.Chtimes(path, time.Time{}, t); err != nil {
		log.Fatal(err)
	}
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

	tm, err := x.DateTime()
	if err != nil {
		log.Fatal(err)
	}

	return tm
}
