package check

import (
	"github.com/juju/errors"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
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
	register.RegisterCheck(newXXXXChecker(&XXXXOption{}, &RealXXXXRetriever{option.Opt}))
}

type XXXXRetriever interface {
	Get() (*XXXXOption, error)
}

type RealXXXXRetriever struct{
	*optionv1.Option
}

var _ XXXXRetriever = &RealXXXXRetriever{}

func (r *RealXXXXRetriever) Get() (actual *XXXXOption,err error) {
	// to complete
	return
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

	ok, err := c.diff(actual)
	if err != nil {
		return errors.Trace(err)
	}
	if ok {
		c.rc = v1.PASS
	}
	return nil
}

func (c *XXXXChecker) diff(actual *XXXXOption) (bool, error) {
	pass := true

	// to complete

	return pass, nil
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
