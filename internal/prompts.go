package internal

import (
	"fmt"
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

Example: /Users/user/Desktop
	`)
}
