package check

import (
	"testing"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	"github.com/stretchr/testify/assert"
)

func TestXXXXChecker(t *testing.T) {
	testCases := []struct {
		desc string

		// to complete

		expectRC v1.ReturnCode
	}{
		// to complete
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// to complete
			checker := newXXXXChecker(&XXXXOption{}, &mockXXXXRetriever{actual: &XXXXOption{}, err: nil})
			err := checker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, checker.ReturnCode())
		})
	}
}

type mockXXXXRetriever struct {
	actual *XXXXOption
	err    error
}

func (r *mockXXXXRetriever) Get() (*XXXXOption, error) {
	return r.actual, r.err
}
