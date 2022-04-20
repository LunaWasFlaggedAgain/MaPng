package main

import (
	"fmt"
	"image/png"
	"io/fs"
	"os"
	"runtime"
	"strings"
	"sync"
)

var FileCh = make(chan fs.FileInfo, 1000)
var wg sync.WaitGroup

func init() {
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			ExportLoop()
			wg.Done()
		}()
	}
}

func ExportLoop() {
	for file := range FileCh {
		cleanName := file.Name()[:strings.Index(file.Name(), ".")]

		file, err := os.Open("./maps/" + file.Name())
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		i, err := MapToImage(file)
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		file.Close()

		destFile, err := os.Create("./images/" + cleanName + ".png")
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		err = png.Encode(destFile, i)
		if err != nil {
			fmt.Printf("got error while processing file %s: %v\n", file.Name(), err)
			continue
		}

		destFile.Close()
	}
}
