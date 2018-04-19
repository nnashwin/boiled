package main

import (
	"testing"
)

func TestdoesFileExist(t *testing.T) {
	cases := []struct {
		path     string
		expected bool
		failMsg  string
	}{
		{
			"./fixtures/willFail.go",
			false,
			"The path \"./fixtures/willFail.go\" should return false and fail.  It did not, so something is wrong",
		},
		{
			"./fixtures/test-file.txt",
			true,
			"The test-file.txt was not found in the fixtures.  Please return it and run the tests again",
		},
	}
	for _, c := range cases {
		actual := doesFileExist(c.path)
		if actual != c.expected {
			t.Errorf("%q", c.failMsg)
		}
	}
}
