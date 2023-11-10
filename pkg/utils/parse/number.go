package parse

import (
	"strings"

	"github.com/docker/go-units"
	"github.com/juju/errors"
)

// ParseDiskForDir parse disk element for dir, expected size and failed suggestion
func ParseDiskForDir(element string) (dir string, expSize int64, failedSug string, err error) {
	list := strings.SplitN(element, ";", 3)
	if len(list) != 3 {
		return "", 0, "", errors.Errorf("invalid disk element %s", element)
	}
	expSize, err = ParseToNumber(list[1])
	if err != nil {
		return "", 0, "", errors.Errorf("invalid size format %s, %v", list[1], err)
	}
	return list[0], expSize, list[2], nil
}

// Convertir números a unidades de bytes
func ParseSize(size float64) string {
	return units.BytesSize(size)
}

func ParseToNumber(size string) (int64, error) {
	return units.RAMInBytes(size)
}
