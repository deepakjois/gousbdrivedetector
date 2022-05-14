package usbdrivedetector

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
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

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		if macOSPattern.MatchString(line) {
			d := macOSPattern.FindStringSubmatch(line)[1]
			driveMap[d] = true
		}
	}

	for k := range driveMap {
		file, err := os.Open(k)
		if err == nil {
			drives = append(drives, k)
		}
		file.Close()
	}

	return drives, nil
}
