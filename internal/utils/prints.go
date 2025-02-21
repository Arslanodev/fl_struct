package utils

import (
	"fmt"
	"strconv"

	"github.com/Arslanodev/fl_struct/internal"
)

const (
	reset    = "\033[0m"
	boldCyan = "\033[1;36m"
	boldRed  = "\033[1;31m"
	dimWhite = "\033[2;37m"
	yellow   = "\033[0;33m"
	green    = "\033[0;32m"
	red      = "\033[0;31m"
	white    = "\033[0;37m"
	blue     = "\033[0;34m"
	magenta  = "\033[0;35m"
)

func PrintFileInfo(info internal.FileInfo, format map[string]string) {
	if info.Kind == "" {
		info.Kind = "folder"
	}

	// Count (dim white)
	fmt.Printf("%s"+format["count"]+"%s ", dimWhite, strconv.Itoa(int(info.Count)), reset)

	// Size (color based on size)
	sizeColor := white
	switch {
	case info.Size[len(info.Size)-2:] == "GB":
		sizeColor = boldRed
	case info.Size[len(info.Size)-2:] == "KB":
		sizeColor = green
	case info.Size[len(info.Size)-2:] == "MB":
		sizeColor = red
	case info.Size[len(info.Size)-1:] == "B":
		sizeColor = yellow
	}
	fmt.Printf("%s"+format["size"]+"%s ", sizeColor, info.Size, reset)

	// Name (white)
	fmt.Printf("%s"+format["filename"]+"%s ", white, info.Name, reset)

	// Kind (color based on type)
	kindColor := white
	switch info.Kind {
	case "folder":
		kindColor = blue
	case ".mp3":
		kindColor = magenta
	default:
		kindColor = yellow
	}
	fmt.Printf("%s"+format["kind"]+"%s ", kindColor, info.Kind, reset)

	// Date (dim white)
	fmt.Printf("%s"+format["date"]+"%s\n", dimWhite, info.DateAdded, reset)
}

func BuildFormat(color, column, reset string) string {
	return color + column + reset
}
