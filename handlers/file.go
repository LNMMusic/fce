package handlers

import (
	"os"
	"io/ioutil"
)

func FileToBytes(filename string) ([]byte, error) {
	// File [open]
	file, err := os.Open(filename);		defer file.Close()
	if err != nil { return nil, err }

	// File [read]
	bytes, _ := ioutil.ReadAll(file)

	return bytes, nil
}