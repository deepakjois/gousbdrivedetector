package usbdrivedetector

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

func Detect() ([]string, error) {
	var drives []string
	driveMap := make(map[string]bool)
	cmd := "system_profiler"
	macOSPattern := regexp.MustCompile("^.*Mount Point: (.+)$")
	args := []string{"SPUSBDataType"}
	if out, err := exec.Command(cmd, args...).Output(); err != nil {
		return drives, err
	} else {
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
			drives = append(drives, k)
		}
	}
	return drives, nil
}
