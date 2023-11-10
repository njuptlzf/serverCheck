package system

import "syscall"

// GetAvailableSpace returns the available disk space of the given path in GB
func GetAvailableSpace(path string) (float64, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, err
	}
	// Available blocks * size per block = available space in bytes
	availableBytes := stat.Bavail * uint64(stat.Bsize)
	return float64(availableBytes), nil
}
