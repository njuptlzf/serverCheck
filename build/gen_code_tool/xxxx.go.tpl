package check

import (
	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	// "github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/juju/errors"
)

type XXXXChecker struct {
	// 名称
	name string
	// 具体检查项
	item *XXXXOption
	// 详细描述
	description string
	// 失败建议
	suggestionOnFail string
	// 返回码: fail 失败, warn 警告, ok 成功
	rc v1.ReturnCode
	// 检查实际结果
	result string
	// 专用的获取接口
	retriever XXXXRetriever
}

// 专用的检查项, 需要按需定义
type XXXXOption struct {
	// to complete
}

// 1. hard code is used to describe the Checker interface that must be implemented. 2. for automatic registration
var _ v1.Checker = &XXXXChecker{}

func init() {
	// to complete
	register.RegisterCheck(newXXXXChecker(&XXXXOption{}, &RealXXXXRetriever{}))
}

// Special interface needs to be implemented on demand
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
		description: "",
		retriever: retriever,
	}
}

func (c *XXXXChecker) Check() error {
	actual, err := c.retriever.Get()
	if err != nil {
		return errors.Trace(err)
	}

	if c.diff(actual) {
		c.rc = v1.PASS
	} else {
		// Returns WARN or FAIL when failed, defined as needed
		c.rc = v1.FAIL
		// c.rc = v1.WARN
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
