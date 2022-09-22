package usbdrivedetector

import (
	"fmt"
	"testing"
)

func TestDetect(t *testing.T) {
	drives, err := Detect()
	if err != nil {
		fmt.Println(drives)
	}
}
