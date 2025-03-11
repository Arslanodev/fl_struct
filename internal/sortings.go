package internal

import (
	"sort"
)

// SortBySize sorts files by their size in a descending order
func SortBySize(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].ByteSize > files[j].ByteSize
	})
}

// SortByFileKind sorts files by their kind in descending order
func SortByFileKind(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Kind > files[j].Kind
	})
}

// SortByDateAdded sorts files by their date-added in a descending order
func SortByDateAdded(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].DateAdded > files[j].DateAdded
	})
}

// SortByFilename sorts files by their names in a ascending order
func SortByFileName(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})
}
