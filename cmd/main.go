package main

import (
	"fmt"

	"github.com/Arslanodev/fl_struct/internal"
	"github.com/Arslanodev/fl_struct/internal/utils"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select functions to execute:",
		Items: []string{
			"List files",
			"Group files (Under development)",
			"Search files (Under development)",
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
		dirPath, option := internal.PromptDirPathWithOptions()
		utils.ListFiles(dirPath, option)
	}
}
