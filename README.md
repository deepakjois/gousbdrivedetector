# Detecting USB Storage Drives in Go

[![GoDoc](https://godoc.org/github.com/deepakjois/gousbdrivedetector?status.svg)](https://godoc.org/github.com/deepakjois/gousbdrivedetector)

Simple API to detect USB Storage device for Go. Works OS X, Linux and (with some issues, see note below) on Windows.

### Usage
Check the [example] folder.

[example]:https://github.com/deepakjois/gousbdrivedetector/tree/master/example

### Issues with some USB flash drives on Windows

On Windows, we use the following command to detect removable media on the system:

```
wmic logicaldisk where drivetype=2 get deviceid
```

This checks for a special bit descriptor on the flash drive which indicates that the USB flash drive is a removable device. This should work as expected on most new flash drives. On some older drives, this bit may not be set leading to drives getting recognized as a local disk instead of a removable drive, which can lead to a false negative.

### Credits
The code was inspired by [USB Drive Detector](https://github.com/samuelcampos/usbdrivedetector), which is written for Java.
