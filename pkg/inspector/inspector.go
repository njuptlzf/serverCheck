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

// Inspector 是一个检查器，用于检查环境各种检查项是否符合要求
type Inspector struct {
	// 命令行选项, 用于覆盖检查项默认值, 判断检查项开关
	option *optionv1.Option
	// 所有检查项
	checkers []checkv1.Checker
	// 最终检查结果, 默认为 success, 如果有检查失败或者警告，会被设置为 fail 或者 warn
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

// 如果为 fail, 不通过; 否则通过
func (i *Inspector) ZeroRc() bool {
	if i.option.Strict {
		return i.rc == checkv1.PASS
	}
	return i.rc != checkv1.FAIL
}

// 设置检查器的结果, 优先级 fail > warn > success, 意味着如果传进来是fail, 直接设置成fail, 如果传进来是warn, 但是当前结果是fail, 返回，以此类推
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
	result := "| Ability | Details | result | Passed | Recommendation |\n| --- | --- | --- | --- |\n"
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
