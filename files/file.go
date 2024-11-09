package files

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(path string) ([]string, error) {
	var lines []string
	file, err := os.OpenFile(path, os.O_RDONLY, 0444) // r--r--r--
	defer file.Close()
	if err != nil {
		return lines, err
	}
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func AppendFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666) // rw-rw-rw-
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	file.Sync()
	return nil
}

func OverwriteFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666) // rw-rw-rw-
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	file.Sync()
	return nil
}
