package usbdrivedetector

import (
	"fmt"
	"testing"
)

func TestDetect(t *testing.T) {
	drives, err := Detect()
	if drives != nil {
		fmt.Println(drives)
	}
}
