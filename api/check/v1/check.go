package v1

type ReturnCode int

const (
	PASS ReturnCode = iota
	WARN
	FAIL
)

// 通用的检查器接口
type Checker interface {
	// 检查项名称
	Name() string
	// 检查项详细描述信息
	Description() string
	// 执行检查
	Check() error
	// 检查返回码
	ReturnCode() ReturnCode
	// 检查实际结果
	Result() string
	// 检查不通过时候的建议
	SuggestionOnFail() string
}
