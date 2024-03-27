package tests

import (
	"fl_struct/cmd"
	"testing"
)

func Test_GetFileExtension(t *testing.T) {
	arg := "another.mp4"
	function := cmd.GetFileExtension(arg)
	expected := "mp4"
	if function != expected {
        t.Errorf("got %q, wanted %q", function, expected)
    }
}
