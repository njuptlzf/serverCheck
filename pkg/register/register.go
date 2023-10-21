package register

import (
	"fmt"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
)

var Checks []v1.Checker

// RegisterCheck 函数用于注册能力检查函数
func RegisterCheck(check v1.Checker) {
	fmt.Println("register checker:", check.Name())
	Checks = append(Checks, check)
}
