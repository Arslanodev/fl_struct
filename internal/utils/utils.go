package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Arslanodev/fl_struct/internal"
)

func GetFileExtension(filepath string) string {
	extension := path.Ext(filepath)
	ext := strings.Replace(extension, ".", "", 1)

	return ext
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to dest from source: %v", err)
	}

	inputFile.Close()

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}
	return nil
}

func Structurize(root_dir string) {

	files, _ := os.ReadDir(root_dir)

	for _, item := range files {
		if !item.IsDir() {
			src_path := filepath.Join(root_dir, item.Name())
			fl_ext := GetFileExtension(src_path)

			dest_dir := filepath.Join(root_dir, fl_ext)

			os.Mkdir(dest_dir, os.ModePerm)

			dest_fl_path := filepath.Join(dest_dir, item.Name())
			MoveFile(src_path, dest_fl_path)

			result := src_path + " -> " + dest_fl_path
			fmt.Println(result)
		}
	}
}

func DetermineColumnLengths(files []fs.DirEntry) map[string]string {
	var lengths internal.FileColumnLengths
	for count, file := range files {
		// Count column length
		if len(strconv.Itoa(count+1)) > lengths.Count {
			lengths.Count = len(strconv.Itoa(count + 1))
		}

		// Filename column length
		if len(file.Name()) > lengths.Filename {
			lengths.Filename = len(file.Name())
		}

		// Size column length
		info, _ := file.Info()
		if len(FormatBytes(uint64(info.Size()))) > lengths.Size {
			lengths.Size = len(FormatBytes(uint64(info.Size())))
		}

		// Kind column length
		if len(filepath.Ext(file.Name())) > lengths.Kind {
			lengths.Kind = len(filepath.Ext(file.Name()))
		}

		// Date column length
		if len(info.ModTime().Format("2025-02-18 22:05:19")) > lengths.Date {
			lengths.Date = len(info.ModTime().String())
		}
	}

	lengths.Count = len(strconv.Itoa(lengths.Count))

	data := make(map[string]string)
	data["count"] = "%-" + strconv.Itoa(lengths.Count+7) + "s"
	data["filename"] = "%-" + strconv.Itoa(lengths.Filename+2) + "s"
	data["size"] = "%-" + strconv.Itoa(lengths.Size+2) + "s"
	data["kind"] = "%-" + strconv.Itoa(lengths.Kind+2) + "s"
	data["date"] = "%-" + strconv.Itoa(lengths.Date+2) + "s"

	return data
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

	var filesInfo []internal.FileInfo
	for index, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		fileInfo := internal.FileInfo{
			Count:     int64(index + 1),
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

	for _, info := range filesInfo {
		PrintFileInfo(info, lengths)
	}
}

func FormatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func GetDirSize(path string) (int64, error) {
	var size int64

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Failed to read directory")
		return 0, err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			// Recursively get the size of the subdirectory
			subDirSize, err := GetDirSize(fullPath)
			if err != nil {
				fmt.Println("Failed to get dir size")
				return 0, err
			}
			size += subDirSize
		} else {
			// Get the file size
			info, err := entry.Info()
			if err != nil {
				fmt.Println("Failed to get entry info")
				return 0, err
			}
			size += info.Size()
		}
	}

	return size, nil
}
