package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	n := time.Now()
	m := n.Month().String()
	folder := os.Getenv("GOOGLE_DRIVE_FOLDER")
	currFolder := folder + m[:3] + "-" + strconv.Itoa(n.Year()) + "/"
	if _, err := os.Stat(currFolder); os.IsNotExist(err) {
		os.Mkdir(currFolder, os.ModePerm)
		fmt.Println("Create folder for this month " + currFolder)
	}

	// create TODO for today
	currFile := currFolder + strconv.Itoa(n.Day()) + "_" + m[:3] + ".TODO"
	if _, err := os.Stat(currFile); os.IsNotExist(err) {
		_, _ = os.Create(currFile)
	}

	// check yesterday TODO file and clean up if I don't add any todo
	yesterdayFile := currFolder + strconv.Itoa(n.Day()-1) + "_" + m[:3] + ".TODO"
	if f, err := os.Stat(yesterdayFile); err == nil {
		if f.Size() == int64(0) {
			fmt.Println("Clean up yesterday file")
			_ = os.Remove(yesterdayFile)
		}
	}
}
