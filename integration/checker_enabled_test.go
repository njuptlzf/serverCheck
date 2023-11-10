package integration_test

import (
	"testing"

	"github.com/njuptlzf/servercheck/pkg/inspector"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnabled(t *testing.T) {
	for _, c := range register.Checks {
		_, err := inspector.CheckerEnabled(c)
		assert.NoError(t, err)
	}
}
