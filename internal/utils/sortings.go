package utils

import (
	"io/fs"
	"path/filepath"
	"sort"
)

// SortBySize sorts files by their size in a descending order
func SortBySize(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1, _ := files[i].Info()
		info2, _ := files[j].Info()
		return info1.Size() > info2.Size()
	})
}

// SortByFileKind sorts files by their kind in descending order
func SortByFileKind(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1 := filepath.Ext(files[i].Name())
		info2 := filepath.Ext(files[j].Name())
		return info1 > info2
	})
}

// SortByDateAdded sorts files by their date-added in a descending order
func SortByDateAdded(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1, _ := files[i].Info()
		info2, _ := files[j].Info()
		return info1.ModTime().String() > info2.ModTime().String()
	})
}

// SortByFilename sorts files by their names in a ascending order
func SortByFileName(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
}
