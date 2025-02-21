package utils

import (
	"fmt"
	"strconv"

	"github.com/Arslanodev/fl_struct/internal"
)

func PrintFileInfo(info internal.FileInfo, format string) {
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
