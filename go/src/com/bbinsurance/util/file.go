package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// delete file
func DeleteFile(file string) error {
	return os.Remove(file)
}

// rename file name
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// put string to file
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

// it returns false when it's a directory or does not exist.
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// get string from text file
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

func FileCreate(path string) (*os.File, error) {
	parentFolder := filepath.Dir(path)
	err := os.MkdirAll(parentFolder, 0777)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("FileCreate permission err %s", err)
		}
		return nil, err
	} else {
		file, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		return file, nil
	}
}
