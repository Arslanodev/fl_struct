package cmd

import (
	"path"
	"strings"
)

func GetFileExtension(filepath string) string {
	extension := path.Ext(filepath)
	ext := strings.Replace(extension, ".", "", 1)

	return ext
}
