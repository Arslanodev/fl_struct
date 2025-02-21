package utils

import (
	"io/fs"
	"path/filepath"
	"sort"
)

func SortBySize(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1, _ := files[i].Info()
		info2, _ := files[j].Info()
		return info1.Size() > info2.Size()
	})
}

func SortByFileKind(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1 := filepath.Ext(files[i].Name())
		info2 := filepath.Ext(files[j].Name())
		return info1 > info2
	})
}

func SortByDateAdded(files []fs.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		info1, _ := files[i].Info()
		info2, _ := files[j].Info()
		return info1.ModTime().String() > info2.ModTime().String()
	})
}
