package files

import (
	"github.com/tkech17/jack_compiler/utils/functional"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Could Not Read File")
	}
	return string(data)
}

func IsDir(path string) bool {
	return strings.LastIndex(path, ".") == -1
}

//noinspection GoVarAndConstTypeMayBeOmitted
func GetFilesWithPrefix(dir string, prefix string) []string {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for _, file := range fileInfos {
		fileNames = append(fileNames, file.Name())
	}

	var jackFiles []string = functional.Filter(fileNames, isJackFile)
	var directories []string = functional.Filter(fileNames, IsDir)

	var result []string = functional.MapString(jackFiles, func(filename string) string {
		return dir + "/" + filename
	})
	for _, directory := range directories {
		var files []string = GetFilesWithPrefix(dir+"/"+directory, prefix)
		result = append(result, files...)
	}
	return result
}

func isJackFile(filename string) bool {
	return strings.HasSuffix(filename, ".jack")
}
