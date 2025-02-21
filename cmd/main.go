package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type FileInfo struct {
	Count     int64
	Name      string
	Size      int64
	Kind      string
	DateAdded string
}

type FileColumnLengths struct {
	Count    int
	Filename int
	Size     int
	Kind     int
	Date     int
}

func main() {
	prompt := promptui.Select{
		Label: "Select functions to execute:",
		Items: []string{
			"List files (sort by: name, size, kind, date added)",
			"Search files",
			"Group files",
			"Index files (for quicker search)",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch index {
	case 0:
		dirPath, option := promptDirPathWithOptions()
		listFiles(dirPath, option)
	}
}

func promptDirPathWithOptions() (string, string) {
	fmt.Println(`
Select directory path

Options:
-s sort by file size
-k sort by file kind
-d sort by file date added

Example: ./Users/arslan/Desktop -s
	`)
	prompt := promptui.Prompt{
		Label: "Enter directory path",
	}

	dirPath, err := prompt.Run()
	if err != nil {
		fmt.Println("Error")
	}

	params := strings.Split(dirPath, " ")

	if len(params) != 2 {
		params = append(params, "")
	}

	return params[0], params[1]
}

func listFiles(dirPath string, option string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Failed to read base directory path")
	}

	// Sorting
	switch option {
	case "-s":
		SortBySize(entries)
	case "-k":
		SortByFileKind(entries)
	case "-d":
		SortByDateAdded(entries)
	}

	lengths := determineColumnLengths(entries)
	format := fmt.Sprintf("%s %s %s %s %s\n", lengths["count"], lengths["size"], lengths["filename"], lengths["kind"], lengths["date"])
	fmt.Printf(format, "count", "size", "name", "kind", "date")

	for index, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		printFileInfo(FileInfo{
			Count:     int64(index + 1),
			Name:      entry.Name(),
			Size:      info.Size(),
			Kind:      filepath.Ext(entry.Name()),
			DateAdded: info.ModTime().Format("2006-01-02 15:04:05"),
		}, format)
	}

}

func printFileInfo(info FileInfo, format string) {
	if info.Kind == "" {
		info.Kind = "folder"
	}
	fmt.Printf(format,
		strconv.Itoa(int(info.Count)),
		strconv.Itoa(int(info.Size)),
		info.Name,
		info.Kind,
		info.DateAdded,
	)
}

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

func determineColumnLengths(files []fs.DirEntry) map[string]string {
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
		if len(strconv.Itoa(int(info.Size()))) > lengths.Size {
			lengths.Size = len(strconv.Itoa(int(info.Size())))
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
