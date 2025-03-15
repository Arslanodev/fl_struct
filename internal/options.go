package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
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
			// Check if the file name does not start with a dot
			if item.Name()[0] != '.' {
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

	if option == "-l" {
		var filesInfo []FileInfo
		WalkThroughFolder(fullpath, &filesInfo)
		SortBySize(filesInfo)
		for index, info := range filesInfo {
			info.Count = int64(index) + 1
			PrintFileInfo(info, lengths)
		}
		return
	}

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

// It performs a partial match (substring search) on the file or folder name.
func SearchFileOrFolder(dirPath string) {
	var filepaths []string

	directory, _ := filepath.Abs(filepath.Clean(dirPath))

	// Walk through the directory and its subdirectories
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		filepaths = append(filepaths, path)
		return nil
	})

	if err != nil {
		fmt.Println("Failed to walk through file: ", err)
		return
	}

	prompt := promptui.Select{
		Label:             "Search for a fruit",
		Items:             filepaths,
		Size:              7, // Number of items to display at once
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			// Perform a case-insensitive substring search
			fruit := filepaths[index]
			return strings.Contains(strings.ToLower(fruit), strings.ToLower(input))
		},
	}

	// Run the prompt
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	// Display the selected result
	OpenLocation(result)
}

func WalkThroughFolder(fullpath string, fileInfo *[]FileInfo) {
	entries, err := os.ReadDir(fullpath)
	if err != nil {
		fmt.Println("Failed to read base directory path")
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		if entry.IsDir() {
			path := filepath.Join(fullpath, entry.Name())
			WalkThroughFolder(path, fileInfo)
		} else {
			*fileInfo = append(*fileInfo, FileInfo{
				Name:      entry.Name(),
				Kind:      filepath.Ext(entry.Name()),
				DateAdded: info.ModTime().Format("2006-01-02 15:04:05"),
				Size:      FormatBytes(uint64(info.Size())),
				ByteSize:  info.Size(),
			})
		}
	}
}
