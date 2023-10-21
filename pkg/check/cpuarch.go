package check

import (
	"fmt"
	"runtime"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	// "github.com/njuptlzf/servercheck/pkg/option"
	"github.com/njuptlzf/servercheck/pkg/register"
	"github.com/njuptlzf/servercheck/pkg/utils/diff"
	"github.com/juju/errors"
)

type CPUArchChecker struct {
	// 名称
	name string
	// 具体检查项
	item *CPUArchOption
	// 详细描述
	description string
	// 失败建议
	suggestionOnFail string
	// 返回码: fail 失败, warn 警告, ok 成功
	rc v1.ReturnCode
	// 检查实际结果
	result string
	// 专用的获取接口
	retriever CPUArchRetriever
}

// 专用的检查项, 需要按需定义
type CPUArchOption struct {
	// 架构
	arch []string
}

// 1. hard code is used to describe the Checker interface that must be implemented. 2. for automatic registration
var _ v1.Checker = &CPUArchChecker{}

func init() {
	register.RegisterCheck(newCPUArchChecker(&CPUArchOption{
		arch: []string{
			"amd64",
			"arm64",
		},
	}, &RealCPUArchRetriever{}))
}

// Special interface needs to be implemented on demand
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
