package utils

import (
	"sort"

	"github.com/Arslanodev/fl_struct/internal"
)

// SortBySize sorts files by their size in a descending order
func SortBySize(files []internal.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ByteSize > files[j].ByteSize
	})
}

// SortByFileKind sorts files by their kind in descending order
func SortByFileKind(files []internal.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Kind > files[j].Kind
	})
}

// SortByDateAdded sorts files by their date-added in a descending order
func SortByDateAdded(files []internal.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].DateAdded > files[j].DateAdded
	})
}

// SortByFilename sorts files by their names in a ascending order
func SortByFileName(files []internal.FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
}
