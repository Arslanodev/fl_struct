package internal

import "testing"

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		name     string // Test case name
		filepath string // Input file path
		expected string // Expected output
	}{
		{
			name:     "Standard file with extension",
			filepath: "example.txt",
			expected: "txt",
		},
		{
			name:     "File with multiple dots",
			filepath: "archive.tar.gz",
			expected: "gz",
		},
		{
			name:     "File with no extension",
			filepath: "README",
			expected: "",
		},
		{
			name:     "File with dot in directory name",
			filepath: "path/to.some/file",
			expected: "",
		},
		{
			name:     "File with dot at the end",
			filepath: "file.",
			expected: "",
		},
		{
			name:     "File with multiple extensions",
			filepath: "image.jpeg.png",
			expected: "png",
		},
		{
			name:     "Hidden file with extension",
			filepath: ".gitignore",
			expected: "gitignore",
		},
		{
			name:     "Empty file path",
			filepath: "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFileExtension(tt.filepath)
			if result != tt.expected {
				t.Errorf("GetFileExtension(%q) = %q; expected %q", tt.filepath, result, tt.expected)
			}
		})
	}
}
