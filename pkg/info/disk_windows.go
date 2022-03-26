//go:build windows

// Package info is a house for gathering further information about specific
// components relevant to roadie.
package info

import (
	"syscall"
	"unsafe"
)

// Disk is a structure to capture some space utilisation details about a
// specific disk.
type Disk struct {
	Available int64  `json:"available"`
	Free      int64  `json:"free"`
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	Used      int64  `json:"used"`
}

// DiskDetails will accept a list of strings that comprise a list of paths
// in the local file system to then provide usage information for each.
func DiskDetails(d []string) []Disk {
	i := []Disk{}

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	for _, x := range d {
		disk := Disk{}

		c.Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(x))),
			uintptr(unsafe.Pointer(&disk.Free)),
			uintptr(unsafe.Pointer(&disk.Size)),
			uintptr(unsafe.Pointer(&disk.Available)))

		disk.Used = disk.Size - disk.Free
		i = append(i, disk)
	}

	return i
}
