package internal

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

func DetermineColumnLengths(files []fs.DirEntry) map[string]string {
	var lengths FileColumnLengths
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

// OpenLocation opens the parent directory of a file or folder in the system's file explorer.
func OpenLocation(path string) error {
	var cmd *exec.Cmd

	// Determine the OS and construct the appropriate command
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", "-R", path)
	case "linux": // Linux
		cmd = exec.Command("xdg-open", path)
	case "windows": // Windows
		cmd = exec.Command("explorer", "/select", path)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// Run the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open location: %v", err)
	}

	return nil
}
