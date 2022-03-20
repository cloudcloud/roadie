// Package info is a house for gathering further information about specific
// components relevant to roadie.
package info

import "syscall"

// Disk is a structure to capture some space utilisation details about a
// specific disk.
type Disk struct {
	Free uint64
	Size uint64
	Used uint64

	stat *syscall.Statfs_t
}

// DiskDetails will accept a list of strings that comprise a list of paths
// in the local file system to then provide usage information for each.
func DiskDetails(d []string) []Disk {
	i := []Disk{}

	for _, x := range d {
		var stat syscall.Statfs_t
		syscall.Statfs(x, &stat)

		i = append(i, Disk{
			stat: &stat,
			Free: stat.Bfree * uint64(stat.Bsize),
			Size: uint64(stat.Blocks) * uint64(stat.Bsize),
			Used: (uint64(stat.Blocks) * uint64(stat.Bsize)) - (stat.Bfree * uint64(stat.Bsize)),
		})
	}

	return i
}
