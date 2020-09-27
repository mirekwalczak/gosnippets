package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func setup() (string, func()) {
	const testdb = "db.test.bin"

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatalf("could not create temp direstory : %v", err)
	}

	teardown := func() { os.RemoveAll(dir) }

	return filepath.Join(dir, testdb), teardown
}

func main() {
	testdb, teardown := setup()
	log.Println(testdb)
	defer teardown()
}
