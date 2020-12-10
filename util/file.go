package util

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func OpenFileOrCreate(path string) *os.File {
	if !PathExists(path) {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		_ = os.Chmod(path, 0660)
		return file
	} else {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		return file
	}
}
