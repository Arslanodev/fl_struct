package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func Structurize(dirPath string) {
	fullPath, err := filepath.Abs(filepath.Clean(dirPath))
	if err != nil {
		fmt.Println("Failed to get absolute path: ", err)
		return
	}

	files, _ := os.ReadDir(fullPath)

	for _, item := range files {
		if !item.IsDir() {
			src_path := filepath.Join(fullPath, item.Name())
			fl_ext := GetFileExtension(src_path)

			dest_dir := filepath.Join(fullPath, fl_ext)

			os.Mkdir(dest_dir, os.ModePerm)

			dest_fl_path := filepath.Join(dest_dir, item.Name())
			MoveFile(src_path, dest_fl_path)

			result := src_path + " -> " + dest_fl_path
			fmt.Println(result)
		}
	}
}

func ListFiles(dirPath string, option string) {
	fullpath, err := filepath.Abs(filepath.Clean(dirPath))
	if err != nil {
		fmt.Println("Failed to get absolute path: ", err)
		return
	}

	entries, err := os.ReadDir(fullpath)
	if err != nil {
		fmt.Println("Failed to read base directory path")
	}

	lengths := DetermineColumnLengths(entries)

	// Header
	format := "%s" + lengths["count"] + " " + lengths["size"] + " " + lengths["filename"] + " " + lengths["kind"] + " " + lengths["date"] + " " + "%s\n"
	fmt.Printf(format, boldCyan, "count", "size", "name", "kind", "date", reset)

	var filesInfo []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		fileInfo := FileInfo{
			Name:      entry.Name(),
			Kind:      filepath.Ext(entry.Name()),
			DateAdded: info.ModTime().Format("2006-01-02 15:04:05"),
		}

		if entry.IsDir() {
			path := filepath.Join(fullpath, entry.Name())
			size, err := GetDirSize(path)
			if err != nil {
				fmt.Println("failed to get dir size: ", err)
				return
			}
			fileInfo.Size = FormatBytes(uint64(size))
			fileInfo.ByteSize = size
		} else {
			fileInfo.Size = FormatBytes(uint64(info.Size()))
			fileInfo.ByteSize = info.Size()
		}

		filesInfo = append(filesInfo, fileInfo)
	}

	// Sorting
	switch option {
	case "-s":
		SortBySize(filesInfo)
	case "-k":
		SortByFileKind(filesInfo)
	case "-d":
		SortByDateAdded(filesInfo)
	default:
		SortByFileName(filesInfo)
	}

	for index, info := range filesInfo {
		info.Count = int64(index) + 1
		PrintFileInfo(info, lengths)
	}
}
