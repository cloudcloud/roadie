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

// DiskDetails will accept a string that represents a path in the local
// file system, to then provide usage information for it.
func DiskDetails(d string) Disk {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	disk := Disk{}

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(d))),
		uintptr(unsafe.Pointer(&disk.Free)),
		uintptr(unsafe.Pointer(&disk.Size)),
		uintptr(unsafe.Pointer(&disk.Available)))

	disk.Used = disk.Size - disk.Free

	return disk
}
