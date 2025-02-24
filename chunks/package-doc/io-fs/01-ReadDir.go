package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// DirFS returns a file system (an fs.FS) for the tree of
	// files rooted at the directory dir.
	dir := os.DirFS("/home/raf/Go")

	// ReadDir reads the named directory and returns a list of
	// directory entries sorted by filename.
	files, err := fs.ReadDir(dir, ".")
	if err != nil {
		fmt.Println("fs.ReadDir():", err.Error())
	}

	for i, f := range files {
		fmt.Println(i, f.Name())
	}

}
