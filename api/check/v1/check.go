package v1

type ReturnCode int

const (
	PASS ReturnCode = iota
	WARN
	FAIL
)

// Common checker interface
type Checker interface {
	// Check item name
	Name() string
	// Detailed description of the check item
	Description() string
	// Perform check
	Check() error
	// Check return code
	ReturnCode() ReturnCode
	// Actual check result
	Result() string
	// Suggestions when the check fails
	SuggestionOnFail() string
}
