package check

import (
	"fmt"
	"runtime"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
	// "github.com/njuptlzf/servercheck/pkg/option"
	"github.com/juju/errors"
	"github.com/njuptlzf/servercheck/pkg/register"
)

type CPUCoreChecker struct {
	// 名称
	name string
	// 具体检查项
	item *CPUCoreOption
	// 详细描述
	description string
	// 失败建议
	suggestionOnFail string
	// 返回码: fail 失败, warn 警告, ok 成功
	rc v1.ReturnCode
	// 检查实际结果
	result string
	// 专用的获取接口
	retriever CPURetriever
}

// 专用的检查项, 需要按需定义
type CPUCoreOption struct {
	// 核心数
	number int
	// todo: 支持该功能
	// cpu利用率
	// usage float64
}

// 1.hard code 描述需要实现通用 Checker 2.自动注册的标识
var _ v1.Checker = &CPUCoreChecker{}

func init() {
	register.RegisterCheck(newCPUChecker(&CPUCoreOption{
		number: 4,
	}, &RealCPURetriever{}))
}

// 专用的检查项接口，需要按需实现
type CPURetriever interface {
	Get() (*CPUCoreOption, error)
}

type RealCPURetriever struct{}

var _ CPURetriever = &RealCPURetriever{}

func (r *RealCPURetriever) Get() (*CPUCoreOption, error) {
	return &CPUCoreOption{
		number: runtime.NumCPU(),
	}, nil
}

func newCPUChecker(item *CPUCoreOption, retriever CPURetriever) *CPUCoreChecker {
	return &CPUCoreChecker{
		name:        "CPU",
		description: "check CPU core",
		item:        item,
		retriever:   retriever,
	}
}

func (c *CPUCoreChecker) Check() error {
	actual, err := c.retriever.Get()
	if err != nil {
		return errors.Trace(err)
	}

	if c.diff(actual) {
		c.rc = v1.PASS
	} else {
		c.rc = v1.WARN
	}
	return nil
}

func (c *CPUCoreChecker) diff(actual *CPUCoreOption) bool {
	pass := true
	coreNumInfo := fmt.Sprintf("[number of cores] acutal: %d, expect: %d", actual.number, c.item.number)
	c.result += coreNumInfo

	if actual.number < c.item.number {
		pass = false
		c.suggestionOnFail += "[number of cores] increase server's CPU"
	}
	return pass
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
