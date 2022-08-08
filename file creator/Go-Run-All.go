// This file reads all the names of the files with the .go extension in the current directory.
// It then creates a .bat file for each one that runs the go run command on the file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("*.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file)
		f, err := os.Create(file + ".bat")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		f.WriteString("go run " + file + "\ntimeout /t 5")
	}
}
