// Package usbdrivedetector detects all USB storage devices connected to a computer.
// It currently works on OS X, with Linux and Windows support coming soon.
//
// Source code and other details for the project are available at Github:
//
// https://github.com/deepakjois/gousbdrivedetector
//
package usbdrivedetector

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Detect returns a list of file paths pointing to the root folder of
// USB storage devices connected to the system.
func Detect() ([]string, error) {
	var drives []string
	driveMap := make(map[string]bool)
	macOSPattern := regexp.MustCompile("^.*Mount Point: (.+)$")

	cmd := "system_profiler"
	args := []string{"SPUSBDataType"}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		return drives, err
	}

	b := bufio.NewReader(bytes.NewReader(out))
	for {
		line, err := b.ReadString('\n')
		line = strings.TrimSpace(line)
		if macOSPattern.MatchString(line) {
			d := macOSPattern.FindStringSubmatch(line)[1]
			driveMap[d] = true
		}
		if err == io.EOF {
			break
		}
	}

	for k := range driveMap {
		_, err := os.Open(k)
		if err == nil {
			drives = append(drives, k)
		}
	}

	return drives, nil
}
