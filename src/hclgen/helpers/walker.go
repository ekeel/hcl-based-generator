package helpers

import (
	"os"
	"path/filepath"
	"regexp"
)

func walkMatch(root string, pattern string) ([]string, error) {
	var matches []string
	regex := regexp.MustCompile(pattern)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matched := regex.MatchString(path)

		if matched {
			matches = append(matches, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}
