package main

import (
	"embed"
)

//go:embed single_file.txt
var fileString string

//go:embed single_file.txt
var fileByte []byte

//go:embed single_file.txt
//go:embed *.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("file2.hash")
	print(string(content2))
}
