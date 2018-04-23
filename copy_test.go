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
