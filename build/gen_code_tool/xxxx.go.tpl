package check

import (
	"github.com/juju/errors"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
)

var _ v1.Checker = &XXXXChecker{}

type XXXXChecker struct {
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
	retriever XXXXRetriever
}

func init() {
	register.RegisterCheck(newXXXXChecker(&RealXXXXRetriever{exp: &expXXXXOption{Option: option.Opt}}))
}

func newXXXXChecker(retriever XXXXRetriever) *XXXXChecker {
	return &XXXXChecker{
		name:        "XXXX",
		description: "check XXXX",
		retriever:   retriever,
	}
}

func (c *XXXXChecker) Check() error {
	exp, act, err := c.retriever.Collect()
	if err != nil {
		return errors.Trace(err)
	}

	// c.rc = v1.WARN
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

func (c *XXXXChecker) diff(exp *expXXXXOption, act *actXXXXOption) (bool, error) {
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

// XXXXOption is a dedicated check item
type RealXXXXRetriever struct {
	// expect option value
	exp *expXXXXOption

	// actual option value
	act *actXXXXOption
}

type expXXXXOption struct {
	*optionv1.Option
}

type actXXXXOption struct {
	// to complete
}

type XXXXRetriever interface {
	Collect() (*expXXXXOption, *actXXXXOption, error)
}

var _ XXXXRetriever = &RealXXXXRetriever{}

func (r *RealXXXXRetriever) Collect() (*expXXXXOption, *actXXXXOption, error) {
	// to complete

	return r.exp, r.act, nil
}
