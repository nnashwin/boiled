package main

import (
	"fmt"
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
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		t.Error(err)
	}

	CopyDir(filePath, "./fixtures/copy-dir2")

	copiedFiles, err := ioutil.ReadDir("./fixtures/copy-dir2")
	if err != nil {
		t.Error(err)
	}

	for i, file := range files {
		fmt.Println(file.Name())
		fmt.Println(copiedFiles[i].Name())
	}

	err = os.RemoveAll("./fixtures/copy-dir2")
	if err != nil {
		t.Error(err)
	}
}
