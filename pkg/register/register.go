package register

import (
	"fmt"

	v1 "github.com/njuptlzf/servercheck/api/check/v1"
)

var Checks []v1.Checker

// RegisterCheck function is used to register capability check functions
func RegisterCheck(check v1.Checker) {
	fmt.Println("register checker:", check.Name())
	Checks = append(Checks, check)
}
