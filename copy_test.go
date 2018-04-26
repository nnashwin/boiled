package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestCopyFile(t *testing.T) {
	cases := []struct {
		FilePath string
		FailMsg  string
	}{
		{
			"./fixtures/copy-dir/copy-file-test.txt",
			"The CopyFileTest did not copy correctly",
		},
	}

	for _, c := range cases {
		expected, err := ioutil.ReadFile(c.FilePath)
		if err != nil {
			t.Error(err)
		}

		CopyFile(c.FilePath, "./fixtures/copy-dir/copy-file-test2.txt")
		actual, err := ioutil.ReadFile("./fixtures/copy-dir/copy-file-test2.txt")
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(expected, actual) == false {
			t.Error(c.FailMsg)
		}

		err = os.RemoveAll("./fixtures/copy-dir/copy-file-test2.txt")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCopyDir(t *testing.T) {
	filePath := "./fixtures/copy-dir"
	nestPath := "/nested-dir"

	ignoreList := make(map[string]struct{})
	// add .gitignore to the map
	ignoreList[".gitignore"], ignoreList["skip-dir"] = struct{}{}, struct{}{}

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		t.Error(err)
	}

	nestFiles, err := ioutil.ReadDir(filePath + nestPath)
	if err != nil {
		t.Error(err)
	}

	CopyDir(filePath, "./fixtures/copy-dir2", ignoreList)

	copiedFiles, err := ioutil.ReadDir("./fixtures/copy-dir2")
	if err != nil {
		t.Error(err)
	}

	// remove the .gitignore from the slice of files in the copy-dir so the names between the two will match
	for i := len(files) - 1; i >= 0; i-- {
		if files[i].Name() == ".gitignore" || files[i].Name() == "skip-dir" {
			files = append(files[:i], files[i+1:]...)
		}
	}

	for i, file := range files {
		if file.Name() != copiedFiles[i].Name() {
			t.Errorf("The files copied in the copy-dir directories do not match:\nfileName: %s\ncopyFileName: %s", file.Name(), copiedFiles[i].Name())
		}
	}

	nestCopyFiles, err := ioutil.ReadDir("./fixtures/copy-dir2" + nestPath)
	if err != nil {
		t.Error(err)
	}

	for i, file := range nestFiles {
		if file.Name() != nestCopyFiles[i].Name() {
			t.Errorf("The files copied in the nested-dir directories do not match:\nfileName: %s\nnestedFileName: %s", file.Name(), nestCopyFiles[i].Name())
		}
	}

	err = os.RemoveAll("./fixtures/copy-dir2")
	if err != nil {
		t.Error(err)
	}
}
