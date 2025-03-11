package main

import (
	"fmt"

	"github.com/Arslanodev/fl_struct/internal"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select functions to execute:",
		Items: []string{
			"List files",
			"Group files",
			"Search files and folders",
			"Analyze Dir (Under development)",
			"Index files (for quicker search) (Under development)",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch index {
	case 0:
		dirPath, option := internal.PromptDirPathWithOptions("list files")
		internal.ListFiles(dirPath, option)
	case 1:
		dirPath, _ := internal.PromptDirPathWithOptions("group files")
		// Prompt additional validation message
		if internal.PromptYesOrNo(dirPath) {
			internal.Structurize(dirPath)
		}
	case 2:
		keyword := internal.PromptSearchKeyword()
		dirPath, _ := internal.PromptDirPathWithOptions("")

		filePaths, err := internal.SearchFileOrFolder(dirPath, keyword)
		if err != nil {
			fmt.Println("Failed to search files: ", err)
			return
		}

		for _, filepath := range filePaths {
			fmt.Println(filepath)
		}

	}
}
