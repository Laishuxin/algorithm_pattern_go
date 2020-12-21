package StackPractice

import (
	"code/Stack/ArrayListStack"
	"io/ioutil"
	"path"
)

// 使用栈读取指定路径下的所有文件

func GetAllFile(aPath string, files []string) ([]string, error) {
	dir, err := ioutil.ReadDir(aPath)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if file.IsDir() {
			aPath = path.Join(aPath, file.Name())
			files = append(files, aPath)
			files, err = GetAllFile(aPath, files)
		} else {
			files = append(files, path.Join(aPath, file.Name()))
		}
	}
	return files, err
}

func GetAllFileUsingStack(aPath string, files []string) ([]string, error) {

	var err error
	stack := ArrayListStack.NewArrayListStack(40)
	stack.Push(aPath)

	for !stack.IsEmpty() {
		aPath = stack.Pop().(string)
		//files = append(files, aPath)
		dir, _ := ioutil.ReadDir(aPath)
		for _, file := range dir {
			filePath := path.Join(aPath, file.Name())
			files = append(files, filePath)
			if file.IsDir() {
				stack.Push(filePath)
			}
		}
	}

	return files, err
}
