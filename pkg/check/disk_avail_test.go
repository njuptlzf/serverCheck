package check

import (
	"testing"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	"github.com/njuptlzf/servercheck/pkg/utils/parse"

	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		desc          string
		actDiskForDir []string
		expDiskForDir []string
		expSug        string
		expectRC      v1.ReturnCode
	}{
		{
			desc:          "act > exp",
			actDiskForDir: []string{"/var;101G;"},
			expDiskForDir: []string{"/var;100G;The minimum available space for is 100G"},
			expSug:        "",
			expectRC:      v1.PASS,
		},
		{
			desc:          "act == exp",
			actDiskForDir: []string{"/var;100G;"},
			expDiskForDir: []string{"/var;100G;The minimum available space for is 100G"},
			expSug:        "",
			expectRC:      v1.PASS,
		},
		{
			desc:          "act < exp",
			actDiskForDir: []string{"/var;100G;"},
			expDiskForDir: []string{"/var;200G;The minimum available space for is 100G"},
			expSug:        "/var: The minimum available space for is 100G",
			expectRC:      v1.FAIL,
		},
		{
			desc:          "act list < exp list",
			actDiskForDir: []string{"/var;100G;", "/home;200G;"},
			expDiskForDir: []string{"/var;200G;The minimum available space for is 200G", "/home;300G;>=300G"},
			expSug:        "/var: The minimum available space for is 200G\n/home: >=300G",
			expectRC:      v1.FAIL,
		},
		{
			desc:          "first element error",
			actDiskForDir: []string{"/var;99G;", "/home;200G;"},
			expDiskForDir: []string{"/var;100G;The minimum available space for is 200G", "/home;200G;>=200G"},
			expSug:        "/var: The minimum available space for is 200G",
			expectRC:      v1.FAIL,
		},
		{
			desc:          "last element error",
			actDiskForDir: []string{"/var;100G;", "/home;100G;"},
			expDiskForDir: []string{"/var;100G;The minimum available space for is 200G", "/home;200G;>=200G"},
			expSug:        "/home: >=200G",
			expectRC:      v1.FAIL,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			exp := &expDiskAvailOption{
				Option: &optionv1.Option{
					DiskOfDir: tc.expDiskForDir,
				},
			}
			act := &actDiskAvailOption{
				diskOfDir: tc.actDiskForDir,
			}
			checker := newDiskAvailChecker(&mockDiskAvailRetriever{exp: exp, act: act, err: nil})
			err := checker.Check()
			assert.NoError(t, err)
			assert.Equal(t, tc.expectRC, checker.ReturnCode())
			assert.Equal(t, tc.expSug, checker.SuggestionOnFail())
		})
	}
}

type mockDiskAvailRetriever struct {
	exp *expDiskAvailOption
	act *actDiskAvailOption
	err error
}

func (r *mockDiskAvailRetriever) Collect() (*expDiskAvailOption, *actDiskAvailOption, error) {
	return r.exp, r.act, r.err
}

func TestCollect(t *testing.T) {
	testCases := []struct {
		desc      string
		actSize   string
		srcOpt    *expDiskAvailOption
		expActOpt *actDiskAvailOption
		err       error
	}{
		{
			desc: "exp == act",
			srcOpt: &expDiskAvailOption{
				Option: &optionv1.Option{
					DiskOfDir: []string{"/var;100G;The minimum available space for is 100G"},
				},
			},
			actSize: "100G",
			expActOpt: &actDiskAvailOption{
				diskOfDir: []string{"/var;100GiB;"},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			retriever := &RealDiskAvailRetriever{
				exp: tc.srcOpt,
				Get: func(path string) (float64, error) {
					size, err := parse.ParseToNumber(tc.actSize)
					return float64(size), err
				},
			}
			_, act, err := retriever.Collect()
			assert.Equal(t, tc.expActOpt, act)
			assert.Equal(t, tc.err, err)
		})
	}
}
