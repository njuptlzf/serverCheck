package diff

// Returns true if parent contains all elements in sub, otherwise returns false.
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
