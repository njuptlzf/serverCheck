package system

import (
	"testing"
)

func TestGetAvailableSpace(t *testing.T) {
	path := "/"
	availableSpace, err := GetAvailableSpace(path)
	if err != nil {
		t.Errorf("GetAvailableSpace(%s) returned an error: %v", path, err)
	}
	if availableSpace <= 0 {
		t.Errorf("GetAvailableSpace(%s) returned an invalid value: %f", path, availableSpace)
	}
}
