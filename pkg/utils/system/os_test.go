package system

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvailableSpace(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		diskSize float64
		expSize  float64
	}{
		{
			name:     "Valid path",
			path:     "/",
			diskSize: 100.0,
			expSize:  100.0,
		},
		{
			name:     "path not found, use parent",
			path:     "/invalid2/test/path",
			diskSize: 200.0,
			expSize:  200.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actAvail, err := getAvailableSpace(tc.path, func(path string) (float64, error) {
				return tc.diskSize, nil
			})
			assert.NoError(t, err)
			assert.Equal(t, tc.expSize, actAvail)
		})
	}
}

func TestGetSize(t *testing.T) {
	path := "/"
	availableSpace, err := GetAvailableSpace(path)
	if err != nil {
		t.Errorf("GetAvailableSpace(%s) returned an error: %v", path, err)
	}
	if availableSpace <= 0 {
		t.Errorf("GetAvailableSpace(%s) returned an invalid value: %f", path, availableSpace)
	}
	fmt.Println("availableSpace: ", availableSpace)
}
