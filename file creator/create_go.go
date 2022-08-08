//This file creates a folder named test and creates a file named main.go in that folder and writes a go hello world program to the file main.go.
//Then it creates a file named build.bat in the folder test that contains the script to compile the program.
package main

import (
	"fmt"
	"os"
)

func main() {
	//Check if the folder test exists, if it does, exit the program
	if _, err := os.Stat("test go"); err == nil {
		os.Exit(1)
	}

	//Create a folder named test with the file main.go in it
	os.Mkdir("test go", 0777)
	os.Create("test go/main.go")
	//Write a go hello world program to the file main.go
	f, err := os.OpenFile("test go/main.go", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(`// The file prints "Hello world!" and then waits for 3 seconds.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")
	time.Sleep(3 * time.Second)
}`)
	//Create a file named build.bat in the folder test that contains the script to compile the program
	os.Create("test go/build.bat")
	//Write a script to compile the program to the file build.bat
	f, err = os.OpenFile("test go/build.bat", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(`go build main.go`)
}
