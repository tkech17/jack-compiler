package files_test

import (
	"github.com/tkech17/jack_compiler/utils/files"
	"github.com/tkech17/jack_compiler/utils/tests"
	"testing"
)

func TestIsDir_when_IsDirectory(t *testing.T) {
	var folders = [2]string{"folder", "folder/"}

	for _, path := range folders {
		tests.AssertTrue(t, files.IsDir(path), path+" Is Folder")
	}
}
func TestIsDir_when_IsNotDirectory(t *testing.T) {
	var filesArray = [2]string{"folder.VM", "folder/BLA.JPG"}

	for _, path := range filesArray {
		tests.AssertFalse(t, files.IsDir(path), path+" Is Not Folder")
	}
}
