package controller

import (
	"bufio"
	"errors"
	"os"
)

// Parse input config file.
// Return two arrays: 3 config lines and all event lines
func Parse(confPath string) ([]string, []string, error) {
	var confStrings []string
	var taskStrings []string
	f, err := os.Open(confPath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; i < 3; i++ {
		stop := scanner.Scan()
		if !stop {
			break
		}
		confStrings = append(confStrings, scanner.Text())
	}

	if len(confStrings) != 3 {
		return nil, nil, errors.New(confStrings[len(confStrings)-1])
	}

	for scanner.Scan() {
		taskStrings = append(taskStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return confStrings, taskStrings, nil
}
