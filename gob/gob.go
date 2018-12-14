// Package gob implements methods for easy saving to files.
package gob

import (
	"encoding/gob"
	"os"
)

// Writegob writes a "thing" to a file
func WriteGob(filename string, thing interface{}) error {
	file, err := os.OpenFile(
		filename,
		os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// if there was no error write the struct to the file
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(thing)

	// return possible error
	return err

}

// ReadGob reads a "thing" from a file
func ReadGob(filename string, thing interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// if there was no error decode it
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(thing)

	// return possible error
	return err
}
