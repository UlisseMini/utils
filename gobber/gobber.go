// Package gobber implements methods for easy saving of data to files.
package gobber

import (
	"encoding/gob"
	"os"
)

// Write writes a struct to a file
func Write(filename string, thing interface{}) error {
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

// Read reads a struct from a file
func Read(filename string, thing interface{}) error {
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
