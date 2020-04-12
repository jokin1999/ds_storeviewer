package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/gehaxelt/ds_store"
	"io/ioutil"
)

func main() {

	var fileName string

	flag.StringVar(&fileName, "i", "", "DS_Store input file")
	flag.Parse()

	if fileName == "" {
		panic("FileName (-i): Must be specified")
	} else if _, err := os.Stat(fileName); os.IsNotExist(err) {
		panic("FileName (-i): File does not exist:" + fileName)
	}

	dat, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	a, err := ds_store.NewAllocator(dat)

	filenames, err := a.TraverseFromRootNode()
	if err != nil {
		panic(err)
	}

	for _,f := range filenames {
		fmt.Printf("File: %s \n", f)
	}
}