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

	return ParseIntoProc(string(b))
}

// GetUnparsedInfo returns a map
// of every field in the proc file.
func GetUnparsedInfo(p int) (map[string]string, error) {
	filename := fmt.Sprintf(`/proc/%d/status`, p)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return make(map[string]string, 0), err
	}
	return ParseStatus(string(b))
}

// ParseStatus parses the status file
// in /proc/[pid]/status
// and returns a key value pair map
func ParseStatus(s string) (fields map[string]string, err error) {
	lines := strings.Split(s, "\n")
	fields = make(map[string]string)

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

		fields[key] = value
	}

	// i have named return values, this is less confusing though
	return fields, err
}

func ParseIntoProc(s string) (p Proc, err error) {
	fields, err := ParseStatus(s)
	if err != nil {
		return Proc{}, err
	}

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

	// convert some of the fieldss to ints
	// if it fails i want the value to be zero
	// so i don't care about the error
	pid, _ := strconv.Atoi(fields["Pid"])
	threads, _ := strconv.Atoi(fields["Threads"])

	// finally return
	p = Proc{
		Name:    fields["Name"],
		Pid:     pid,
		State:   fields["State"],
		Threads: threads,
	}

	return p, err
}
