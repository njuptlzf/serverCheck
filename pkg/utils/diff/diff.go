package diff

// 如果 parent 包含 sub 中的所有元素，返回 true, 否则返回 false
func SubContains(parent []string, sub []string) bool {
	for _, s := range sub {
		if !contains(parent, s) {
			return false
		}
	}
	return true
}

func contains(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
