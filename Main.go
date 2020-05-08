package main

import (
	"fmt"
	"github.com/tkech17/jack_compiler/utils/files"
	"log"
	"os"
)

//noinspection GoVarAndConstTypeMayBeOmitted
func main() {
	var args = os.Args
	checkArgs(args)

	var jackFiles []string = getJackFilesToRead(args[1])
	fmt.Println(jackFiles)
}

func getJackFilesToRead(filename string) []string {
	if files.IsDir(filename) {
		return files.GetFilesWithPrefix(filename, ".jack")
	} else {
		return []string{filename}
	}
}

func checkArgs(args []string) {
	if len(args) != 2 {
		log.Fatal("arguments len must be 2")
	}
}
