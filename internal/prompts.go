package internal

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func PromptDirPathWithOptions(option string) (string, string) {
	switch option {
	case "list files":
		ListFilesMessage()
	case "group files":
		GroupFilesMessage()
	}

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

func PromptYesOrNo(selectedDirPath string) bool {
	fullPath, err := filepath.Abs(filepath.Clean(selectedDirPath))
	if err != nil {
		fmt.Println("Failed to get fullpath")
		return false
	}

	prompt := promptui.Select{
		Label: fmt.Sprintf(`Are you sure you want to group files in %s? Select [Yes/No]`, fullPath),
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Error")
	}

	if result == "Yes" {
		return true
	}

	return false
}

func ListFilesMessage() {
	fmt.Println(`
List files with various filtering options

Options:
-s sort by file size
-k sort by file kind
-d sort by file date added

Example: /Users/user/Desktop -s
	`)
}

func GroupFilesMessage() {
	fmt.Println(`
Group files into folders according to their kind

Options:
-e group by file extension (.img, .png, .dmg, etc.)
-k group by file kind (videos, audio, documents, etc.)

Example: /Users/user/Desktop
	`)
}
