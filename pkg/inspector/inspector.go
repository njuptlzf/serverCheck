package inspector

import (
	"fmt"

	"github.com/juju/errors"
	checkv1 "github.com/njuptlzf/servercheck/api/check/v1"
	optionv1 "github.com/njuptlzf/servercheck/api/option/v1"
	"github.com/njuptlzf/servercheck/pkg/check"
	"github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
)

// Inspector is a checker used to check whether various environmental check items meet the requirements
type Inspector struct {
	// Command line options, used to override default values of check items and determine check item switches
	option *optionv1.Option
	// All check items
	checkers []checkv1.Checker
	// Final check result, default is success, if there is a check failure or warning, it will be set to fail or warn
	rc checkv1.ReturnCode
}

func NewInspector() *Inspector {
	return &Inspector{
		option:   option.Opt,
		checkers: register.Checks,
		rc:       checkv1.PASS,
	}
}

func (i *Inspector) Check() error {
	for _, c := range i.checkers {
		enable, err := checkerEnabled(c)
		if err != nil {
			return errors.Trace(err)
		}
		if !enable {
			return nil
		}
		fmt.Println("check: ", c.Name())
		if err := c.Check(); err != nil {
			return errors.Trace(err)
		}
		i.setRC(c.ReturnCode())
	}
	return nil
}

// If it is fail, it does not pass; otherwise it passes
func (i *Inspector) ZeroRc() bool {
	if i.option.Strict {
		return i.rc == checkv1.PASS
	}
	return i.rc != checkv1.FAIL
}

// Set the result of the checker, priority fail > warn > success, which means that if it is fail, it is directly set to fail, if it is warn, but the current result is fail, return, and so on
func (i *Inspector) setRC(rc checkv1.ReturnCode) {
	if rc == checkv1.FAIL || i.rc == checkv1.FAIL {
		i.rc = checkv1.FAIL
		return
	}
	if rc == checkv1.WARN || i.rc == checkv1.WARN {
		i.rc = checkv1.WARN
		return
	}
	i.rc = checkv1.PASS
}

func (i *Inspector) PrintResult() error {
	result := "| Ability | Details | Result | Passed | Recommendation |\n| --- | --- | --- | --- | --- |\n"
	for _, checker := range i.checkers {
		enable, err := checkerEnabled(checker)
		if err != nil {
			return errors.Trace(err)
		}
		if enable {
			result += formatResult(checker)
		}
	}
	print(result)
	return nil
}

func formatResult(checker checkv1.Checker) string {
	suggestion := ""
	if checker.ReturnCode() != checkv1.PASS {
		suggestion = checker.SuggestionOnFail()
	}
	return "| " + checker.Name() + " | " + checker.Description() + " | " + checker.Result() + " | " + returnCodeToString(checker.ReturnCode()) + " | " + suggestion + " |\n"
}

func returnCodeToString(rc checkv1.ReturnCode) string {
	switch rc {
	case checkv1.PASS:
		return "pass"
	case checkv1.WARN:
		return "warn"
	case checkv1.FAIL:
		return "fail"
	default:
		return ""
	}
}

func checkerEnabled(checker checkv1.Checker) (bool, error) {
	switch checker.(type) {
	case *check.CPUCoreChecker:
		return option.Opt.CPUCore, nil
	case *check.CPUArchChecker:
		return option.Opt.CPUArch, nil
	default:
		return false, errors.Errorf("unknown checker %s is not supported, type: %T", checker.Name(), checker)
	}
}
