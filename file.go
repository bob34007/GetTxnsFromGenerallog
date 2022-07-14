package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readAllDataFromFile(filename string) (string, error) {
	var data []byte
	var err error

	data, err = os.ReadFile(filename)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
	//return string(data),nil
}

func writeFile(fileName string, data string) error {
	/******************* 使用 io.WriteString 写入文件 **********************/
	f1, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open write file fail ,", fileName, err)
		return err
	}
	defer f1.Close()
	w := bufio.NewWriter(f1)
	n, err := w.WriteString(data) //写入文件(字符串)
	if err != nil {
		fmt.Println("write file fail ,", fileName, err)
		return err
	}
	fmt.Printf("写入 %v %v 个字节\n", fileName, n)
	return nil
}

func getFileListFromDir(path string) ([]string, error) {
	fs := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fs = append(fs, f.Name())
	}
	return fs, err
}

func GetFileNumFromPath(path string) (int64, error) {
	var fileNum int64 = 0
	if len(path) == 0 {
		return 0, nil
	}

	return fileNum, nil
}
