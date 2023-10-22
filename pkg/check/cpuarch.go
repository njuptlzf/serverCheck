package check

import (
	"fmt"
	"runtime"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	// "github.com/njuptlzf/servercheck/pkg/option"
	"github.com/juju/errors"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/njuptlzf/servercheck/pkg/utils/diff"
)

type CPUArchChecker struct {
	// Name
	name string
	// Specific check item
	item *CPUArchOption
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

type CPUArchOption struct {
	// Architecture
	arch []string
}

var _ v1.Checker = &CPUArchChecker{}

func init() {
	register.RegisterCheck(newCPUArchChecker(&CPUArchOption{
		arch: []string{
			"amd64",
			"arm64",
		},
	}, &RealCPUArchRetriever{}))
}

type CPUArchRetriever interface {
	Get() (*CPUArchOption, error)
}

type RealCPUArchRetriever struct{}

var _ CPUArchRetriever = &RealCPUArchRetriever{}

func (r *RealCPUArchRetriever) Get() (*CPUArchOption, error) {
	return &CPUArchOption{
		arch: []string{runtime.GOARCH},
	}, nil
}

func newCPUArchChecker(item *CPUArchOption, retriever CPUArchRetriever) *CPUArchChecker {
	return &CPUArchChecker{
		name:        "CPUArch",
		item:        item,
		description: "check CPU arch",
		retriever:   retriever,
	}
}

func (c *CPUArchChecker) Check() error {
	actual, err := c.retriever.Get()
	if err != nil {
		return errors.Trace(err)
	}

	if c.diff(actual) {
		c.rc = v1.PASS
	} else {
		c.rc = v1.FAIL
	}
	return nil
}

func (c *CPUArchChecker) diff(actual *CPUArchOption) bool {
	pass := true

	archInfo := fmt.Sprintf("[arch] actual: %s, expect: %s", actual.arch, c.item.arch)
	c.result += archInfo

	if !diff.SubContains(c.item.arch, actual.arch) {
		pass = false
		c.suggestionOnFail += "[arch] change to a compatible server"
	}

	return pass
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
