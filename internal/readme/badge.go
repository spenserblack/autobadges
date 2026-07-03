package readme

import (
	"fmt"
	"io"
)

// AddBadges adds badges to the README.
func AddBadges(readme io.ReadWriteSeeker, badges []string) error {
	bytes, err := io.ReadAll(readme)
	if err != nil {
		return err
	}
	index := TitleEndIndex(bytes)
	suffix := bytes[index:]
	_, err = readme.Seek(int64(index), io.SeekStart)
	if err != nil {
		return err
	}

	// NOTE We add an empty newline after the title (if it exists) to make the formatting prettier.
	if index != 0 {
		_, err = fmt.Fprintln(readme)
		if err != nil {
			return err
		}
	}

	for _, badge := range badges {
		_, err := fmt.Fprintln(readme, badge)
		if err != nil {
			return err
		}
	}

	// NOTE If there was no title and we added the badges to the top, we add an extra newline after
	//		the badges to separate the badges from the following paragraph (or whatever syntax).
	if index == 0 {
		_, err = fmt.Fprintln(readme)
		if err != nil {
			return err
		}
	}

	_, err = readme.Write(suffix)
	return err
}
