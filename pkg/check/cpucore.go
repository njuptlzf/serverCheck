package check

import (
	"fmt"
	"runtime"

	"github.com/juju/errors"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
)

var _ v1.Checker = &CPUCoreChecker{}

type CPUCoreChecker struct {
	// Name
	name string
	// Detailed description
	description string
	// Suggestion on failure
	suggestionOnFail string
	// Return code: fail, warn, ok
	rc v1.ReturnCode
	// Actual check result
	result string
	// Dedicated retrieval interface
	retriever CPUCoreRetriever
}

func init() {
	register.RegisterCheck(newCPUCoreChecker(&RealCPUCoreRetriever{exp: &expCPUCoreOption{Option: option.Opt}}))
}

func newCPUCoreChecker(retriever CPUCoreRetriever) *CPUCoreChecker {
	return &CPUCoreChecker{
		name:        "CPUCore",
		description: "check CPU core",
		retriever:   retriever,
	}
}

func (c *CPUCoreChecker) Check() error {
	exp, act, err := c.retriever.Collect()
	if err != nil {
		return errors.Trace(err)
	}

	c.rc = v1.WARN

	ok, err := c.diff(exp, act)
	if err != nil {
		return errors.Trace(err)
	}
	if ok {
		c.rc = v1.PASS
	}
	return nil
}

func (c *CPUCoreChecker) diff(exp *expCPUCoreOption, act *actCPUCoreOption) (bool, error) {
	pass := true
	coreNumInfo := fmt.Sprintf("[number of cores] acutal: %d, expect: %d", act.number, exp.CPUCoreNum)
	c.result += coreNumInfo

	if act.number < exp.CPUCoreNum {
		pass = false
		c.suggestionOnFail += "[number of cores] increase server's CPU"
	}
	return pass, nil
}

func (c *CPUCoreChecker) Name() string {
	return c.name
}

func (c *CPUCoreChecker) Description() string {
	return c.description
}

func (c *CPUCoreChecker) ReturnCode() v1.ReturnCode {
	return c.rc
}

func (c *CPUCoreChecker) Result() string {
	return c.result
}

func (c *CPUCoreChecker) SuggestionOnFail() string {
	return c.suggestionOnFail
}

// CPUCoreOption is a dedicated check item
type RealCPUCoreRetriever struct {
	// expect option value
	exp *expCPUCoreOption

	// actual option value
	act *actCPUCoreOption
}

type expCPUCoreOption struct {
	*optionv1.Option
}

type actCPUCoreOption struct {
	// Number of cores
	number int
}

type CPUCoreRetriever interface {
	Collect() (*expCPUCoreOption, *actCPUCoreOption, error)
}

var _ CPUCoreRetriever = &RealCPUCoreRetriever{}

func (r *RealCPUCoreRetriever) Collect() (*expCPUCoreOption, *actCPUCoreOption, error) {
	r.act = &actCPUCoreOption{}
	r.act.number = runtime.NumCPU()
	return r.exp, r.act, nil
}
