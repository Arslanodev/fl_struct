package internal

import (
	"fmt"
	"strconv"
)

func PrintFileInfo(info FileInfo, format map[string]string) {
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
