package check

import (
	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	// "github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/juju/errors"
)

type XXXXChecker struct {
	// Name
	name string
	// Specific check item
	item *XXXXOption
	// Detailed description
	description string
	// Suggestion on failure
	suggestionOnFail string
	// Return code: fail, warn, or ok
	rc v1.ReturnCode
	// Actual check result
	result string
	// Dedicated retrieval interface
	retriever XXXXRetriever
}

// Dedicated check item
type XXXXOption struct {
	// to complete
}

var _ v1.Checker = &XXXXChecker{}

func init() {
	// to complete
	register.RegisterCheck(newXXXXChecker(&XXXXOption{}, &RealXXXXRetriever{}))
}

type XXXXRetriever interface {
	Get() (*XXXXOption, error)
}

type RealXXXXRetriever struct{}

var _ XXXXRetriever = &RealXXXXRetriever{}

func (r *RealXXXXRetriever) Get() (*XXXXOption, error) {
	// to complete
	return &XXXXOption{}, nil
}

func newXXXXChecker(item *XXXXOption, retriever XXXXRetriever) *XXXXChecker {
	return &XXXXChecker{
		name:      "XXXX",
		item:      item,
		description: "check XXXX",
		retriever: retriever,
	}
}

func (c *XXXXChecker) Check() error {
	actual, err := c.retriever.Get()
	if err != nil {
		return errors.Trace(err)
	}

	// default rc: WARN or FAIL
	c.rc = v1.FAIL
	// c.rc = v1.WARN

	if c.diff(actual) {
		c.rc = v1.PASS
	}
	return nil
}

func (c *XXXXChecker) diff(actual *XXXXOption) bool {
	pass := true

	// to complete

	return pass
}

func (c *XXXXChecker) Name() string {
	return c.name
}

func (c *XXXXChecker) Description() string {
	return c.description
}

func (c *XXXXChecker) ReturnCode() v1.ReturnCode {
	return c.rc
}

func (c *XXXXChecker) Result() string {
	return c.result
}

func (c *XXXXChecker) SuggestionOnFail() string {
	return c.suggestionOnFail
}
