package check

import (
	"testing"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
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
			cpuChecker := newCPUCoreChecker(&mockCPUCoreRetriever{
				exp: &expCPUCoreOption{
					Option: &optionv1.Option{
						CPUCoreNum: 4,
					},
				},
				act: &actCPUCoreOption{
					number: tc.actual,
				},
			})
			err := cpuChecker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, cpuChecker.ReturnCode())
		})
	}
}

type mockCPUCoreRetriever struct {
	exp *expCPUCoreOption
	act *actCPUCoreOption
	err error
}

func (r *mockCPUCoreRetriever) Collect() (*expCPUCoreOption, *actCPUCoreOption, error) {
	return r.exp, r.act, r.err
}
