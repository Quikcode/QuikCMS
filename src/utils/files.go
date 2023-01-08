package utils

import (
	"fmt"
	"os"
)

type Files struct {
	Path string
}

func (i Files) Exist() bool {
	if _, err := os.Stat(i.Path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func (i Files) Write(data string) {
	file, _ := os.Create(i.Path)
	defer file.Close()
	file.WriteString(data)
}

func (i Files) Read() *os.File {
	file, err := os.Open(i.Path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return file
}
