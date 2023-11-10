package check

import (
	"fmt"
	"runtime"

	"github.com/juju/errors"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/njuptlzf/servercheck/pkg/utils/diff"
)

var _ v1.Checker = &CPUArchChecker{}

type CPUArchChecker struct {
	// Name
	name string
	// Detailed description
	description string
	// Suggestion on failure
	suggestionOnFail string
	// Return code: fail, warn, or ok
	rc v1.ReturnCode
	// Actual check result
	result string
	// Dedicated retrieval interface
	retriever CPUArchRetriever
}

func init() {
	register.RegisterCheck(newCPUArchChecker(&RealCPUArchRetriever{
		exp: &expCPUArchOption{
			// todo: use flag
			// Option: option.Opt,
			arch: []string{"amd64", "arm64"},
		},
	}))
}

func newCPUArchChecker(retriever CPUArchRetriever) *CPUArchChecker {
	return &CPUArchChecker{
		name:        "CPUArch",
		description: "check CPU arch",
		retriever:   retriever,
	}
}

func (c *CPUArchChecker) Check() error {
	exp, act, err := c.retriever.Collect()
	if err != nil {
		return errors.Trace(err)
	}

	c.rc = v1.FAIL

	ok, err := c.diff(exp, act)
	if err != nil {
		return errors.Trace(err)
	}
	if ok {
		c.rc = v1.PASS
	}
	return nil
}

func (c *CPUArchChecker) diff(exp *expCPUArchOption, act *actCPUArchOption) (bool, error) {
	pass := true

	archInfo := fmt.Sprintf("[arch] actual: %s, expect: %s", act.arch, exp.arch)
	c.result += archInfo

	if !diff.SubContains(exp.arch, act.arch) {
		pass = false
		c.suggestionOnFail += "[arch] change to a compatible server"
	}

	return pass, nil
}

func (c *CPUArchChecker) Name() string {
	return c.name
}

func (c *CPUArchChecker) Description() string {
	return c.description
}

func (c *CPUArchChecker) ReturnCode() v1.ReturnCode {
	return c.rc
}

func (c *CPUArchChecker) Result() string {
	return c.result
}

func (c *CPUArchChecker) SuggestionOnFail() string {
	return c.suggestionOnFail
}

type RealCPUArchRetriever struct {
	// expect option value
	exp *expCPUArchOption

	// actual option value
	act *actCPUArchOption
}

type expCPUArchOption struct {
	*optionv1.Option
	// todo: use flag
	// Architecture
	arch []string
}

type actCPUArchOption struct {
	// Architecture
	arch []string
}

type CPUArchRetriever interface {
	Collect() (*expCPUArchOption, *actCPUArchOption, error)
}

var _ CPUArchRetriever = &RealCPUArchRetriever{}

func (r *RealCPUArchRetriever) Collect() (*expCPUArchOption, *actCPUArchOption, error) {
	r.act = &actCPUArchOption{}
	r.act.arch = []string{runtime.GOARCH}
	return r.exp, r.act, nil
}
