package system

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/juju/errors"
)

// GetAvailableSpace returns the available space in bytes of the file system containing the given path.
// If the path does not exist, it will recursively check the parent directory until it finds an existing directory.
// It returns the available space in bytes as a float64 and an error if any.
func GetAvailableSpace(path string) (float64, error) {
	return getAvailableSpace(path, getDiskAvail)
}

func getAvailableSpace(path string, calculate func(path string) (float64, error)) (float64, error) {
	if calculate == nil {
		calculate = getDiskAvail
	}
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			parent := filepath.Dir(path)
			fmt.Printf("%s is not exist, use parent %s\n", path, parent)
			if parent == path {
				return 0, errors.Trace(err)
			}
			return getAvailableSpace(parent, calculate)
		}
		return 0, errors.Trace(err)
	}

	return calculate(path)
}

func getDiskAvail(path string) (float64, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0, errors.Trace(err)
	}
	availableBytes := stat.Bavail * uint64(stat.Bsize)
	return float64(availableBytes), nil
}
