package readme

import "regexp"

var titleRegexp = regexp.MustCompile(`#\s+[^\n]+\n`)

// TitleEndIndex gets the index of the end of the title.
func TitleEndIndex(b []byte) int {
	loc := titleRegexp.FindIndex(b)
	if loc == nil {
		// NOTE Just add badges to the beginning.
		return 0
	}

	return loc[1]
}
