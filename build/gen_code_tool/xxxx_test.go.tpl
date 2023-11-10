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
			checker := newXXXXChecker((&mockXXXXRetriever{exp: &expXXXXOption{}, act: &actXXXXOption{}, err: nil}))
			err := checker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, checker.ReturnCode())
		})
	}
}

type mockXXXXRetriever struct {
	exp *expXXXXOption
	act *actXXXXOption
	err error
}

func (r *mockXXXXRetriever) Collect() (*expXXXXOption, *actXXXXOption, error) {
	return r.exp, r.act, r.err
}
