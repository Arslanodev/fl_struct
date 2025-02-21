package internal

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func PromptDirPathWithOptions() (string, string) {
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
