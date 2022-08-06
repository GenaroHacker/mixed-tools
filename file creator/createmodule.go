//This file creates a folder named test and creates a file named main.py in that folder
package main

import (
	"os"
)

func main() {
	//Check if the folder test exists, if it does, exit the program
	if _, err := os.Stat("test"); err == nil {
		os.Exit(1)
	}

	//Create a folder named test with the file main.py in it
	os.Mkdir("test", 0777)
	os.Create("test/main.py")
}
