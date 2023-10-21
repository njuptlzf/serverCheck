package check

import (
	"testing"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	"github.com/stretchr/testify/assert"
)

func TestCPUChecker(t *testing.T) {
	testCases := []struct {
		desc     string
		actual   int
		expectRC v1.ReturnCode
	}{
		{
			desc:     "greater than expected value",
			actual:   5,
			expectRC: v1.PASS,
		},
		{
			desc:     "equal to expected value",
			actual:   4,
			expectRC: v1.PASS,
		},
		{
			desc:     "less than expected value",
			actual:   3,
			expectRC: v1.WARN,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cpuChecker := newCPUChecker(&CPUCoreOption{
				number: 4,
			}, &mockCPURetriever{actual: &CPUCoreOption{
				number: tc.actual,
			}, err: nil})
			err := cpuChecker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, cpuChecker.ReturnCode())
		})
	}
}

type mockCPURetriever struct {
	actual *CPUCoreOption
	err    error
}

func (r *mockCPURetriever) Get() (*CPUCoreOption, error) {
	return r.actual, r.err
}
