// Package badges provides utilities for generating Markdown badge text.
package badges

// Badges gets all of the badges for a project.
func Badges(root string) []string {
	// NOTE 2 Cargo badges
	badges := make([]string, 0, 2)
	cargo := Cargo(root)
	badges = append(badges, cargo...)
	badges = append(badges, GoMod(root))

	return badges
}
