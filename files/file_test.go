package files_test

import (
	"fmt"
	"github.com/liukunc9/go-utils/files"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	file := "C:\\Users\\Admin\\Desktop\\1.txt"
	lines, _ := files.ReadFile(file)
	fmt.Println(lines)
}

func TestOverwriteFile(t *testing.T) {
	file := "C:\\Users\\Admin\\Desktop\\1.txt"
	lines := make([]string, 0)
	lines = append(lines, "5")
	lines = append(lines, "6")
	lines = append(lines, "7")
	files.OverwriteFile(file, strings.Join(lines, "\n"))
}

func TestAppendFile(t *testing.T) {
	file := "C:\\Users\\Admin\\Desktop\\1.txt"
	lines := make([]string, 0)
	lines = append(lines, "8")
	lines = append(lines, "9")
	lines = append(lines, "10")
	files.AppendFile(file, strings.Join(lines, "\n"))
}
