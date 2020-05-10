package main

import (
	"fmt"
	"github.com/tkech17/jack_compiler/compiler/tokenizer"
	"github.com/tkech17/jack_compiler/utils/files"
	"log"
	"os"
	"strings"
)

//noinspection GoVarAndConstTypeMayBeOmitted
func main() {
	var args = os.Args
	checkArgs(args)
	var jackFiles []string = getJackFilesToRead(args[1])

	for _, jackFile := range jackFiles {
		compileJackFile(jackFile)
	}

	fmt.Println(jackFiles)
}

func compileJackFile(file string) {
	var content string = files.ReadFile(file)
	var result string = tokenizer.GetTokensContent(content)
	var targetFilePath string = getTargetFilePath(file)
	files.Write(targetFilePath, result)
}

func getTargetFilePath(file string) string {
	var targetFolder string = "target/"
	var fileNameStartIndex int = strings.LastIndex(file, "/") + 1
	return file[0:fileNameStartIndex] + targetFolder + file[fileNameStartIndex:]
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
