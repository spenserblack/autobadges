// Package stringutil provides utilities for strings.
package stringutil

// Unquote remotes quotes from around a string. Returns the original string if it could not be
// unquoted.
func Unquote(s string) string {
	if len(s) < 2 {
		return s
	}
	left := s[0]
	end := len(s) - 1
	right := s[end]
	if left != right {
		return s
	}
	switch left {
	case '\'', '"', '`':
		return s[1:end]
	}
	return s
}
