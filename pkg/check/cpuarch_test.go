package check

import (
	"testing"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	"github.com/stretchr/testify/assert"
)

func TestCPUArchChecker(t *testing.T) {
	testCases := []struct {
		desc     string
		arch     string
		expectRC v1.ReturnCode
	}{
		{
			desc:     "arch compatible",
			arch:     "amd64",
			expectRC: v1.PASS,
		},
		{
			desc:     "arch is not compatible",
			arch:     "arm32",
			expectRC: v1.FAIL,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			checker := newCPUArchChecker(&CPUArchOption{
				arch: []string{"amd64", "arm64"},
			}, &mockCPUArchRetriever{actual: &CPUArchOption{
				arch: []string{tc.arch},
			}, err: nil})
			err := checker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, checker.ReturnCode())
		})
	}
}

type mockCPUArchRetriever struct {
	actual *CPUArchOption
	err    error
}

func (r *mockCPUArchRetriever) Get() (*CPUArchOption, error) {
	return r.actual, r.err
}
