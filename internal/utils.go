package internal

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
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

func Structurize(root_dir string) {

	files, _ := os.ReadDir(root_dir)

	for _, item := range files {
		if !item.IsDir() {
			src_path := filepath.Join(root_dir, item.Name())
			fl_ext := GetFileExtension(src_path)

			dest_dir := filepath.Join(root_dir, fl_ext)

			os.Mkdir(dest_dir, os.ModePerm)

			dest_fl_path := filepath.Join(dest_dir, item.Name())
			MoveFile(src_path, dest_fl_path)

			result := src_path + " -> " + dest_fl_path
			fmt.Println(result)
		}
	}
}
