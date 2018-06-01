package main

import (
	"os"
	"fmt"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	file.Close()
}


func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "data")
}