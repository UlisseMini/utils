// +build linux darwin

package proc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Proc holds information about a process
type Proc struct {
	Name    string
	Pid     int
	State   string
	Threads int
}

// GetInfo returns the process
// information for p
func GetInfo(p int) (Proc, error) {
	filename := fmt.Sprintf(`/proc/%d/status`, p)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return Proc{}, err
	}

	return ParseStatus(string(b))
}

// GetUnparsedInfo returns a map
// of every field in the proc file.
//func GetUnparsedInfo(p int) (map[]interface{}, error) {}

// ParseStatus parses the status file
// in /proc/[pid]/status
func ParseStatus(s string) (p Proc, err error) {
	lines := strings.Split(s, "\n")
	field := make(map[string]interface{})

	// safety
	defer func() {
		if e := recover(); e != nil {
			switch t := e.(type) {
			case error:
				err = t
			case string:
				err = errors.New(t)
			default:
				err = fmt.Errorf("%v\n", t)
			}
		}
	}()

	// parse the lines into the map
	for _, line := range lines {
		colon := strings.Index(line, ":")
		if colon == -1 {
			continue
		}

		// get key and value
		key := strings.TrimSpace(line[:colon])
		value := strings.TrimSpace(line[colon+1:])

		field[key] = value
	}

	// convert some of the fields to ints
	// if it fails i want the value to be zero
	// so i don't care about the error
	field["Pid"], _ = strconv.Atoi(field["Pid"].(string))
	field["Threads"], _ = strconv.Atoi(field["Threads"].(string))

	// finally return
	p = Proc{
		Name:    field["Name"].(string),
		Pid:     field["Pid"].(int),
		State:   field["State"].(string),
		Threads: field["Threads"].(int),
	}
	// i have named return values, this is less confusing though
	return p, err
}
