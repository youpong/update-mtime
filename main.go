package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

var dryRun bool

func main() {
	var logFilename string
	flag.StringVar(&logFilename, "l", "update-mtime.log", "Write log to filename")
	flag.BoolVar(&dryRun, "d", false, "Write files to log that would be updated but do not update them")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [OPTIONS] [DIR]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	f, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	if dryRun {
		log.Println("dry run")
	}

	if err := filepath.Walk(".", traverse); err != nil {
		fmt.Println(err)
	}
}

func traverse(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	suffixes := []string{"jpeg", "jpg"}

	for _, s := range suffixes {
		if strings.HasSuffix(info.Name(), s) {
			taken := readTakenDate(path)
			if !dryRun {
				updateModTime(path, taken)
			}
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
