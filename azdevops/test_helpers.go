package azdevops

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func testHelperLoadString(t *testing.T, name string) string {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes[:])
}
