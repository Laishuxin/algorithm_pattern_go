package QueuePractice

import (
	"code/Queue/ArrayQueue"
	"errors"
	"io/ioutil"
	"path"
)

var (
	ErrNotADir = errors.New("It is not a directory")
)

// 采用队列获取文件夹中所有文件（广度优先）
func GetAllFiles(aPath string, files []string) ([]string, error) {
	queue := ArrayQueue.NewArrayQueue(20)

	var _err error
	_ = queue.EnQueue(aPath)
	for !queue.IsEmpty() {
		apath, _ := queue.DeQueue()
		aPath = apath.(string)

		dir, err := ioutil.ReadDir(aPath)
		if err != nil {
			break
		}
		_err = err
		for _, file := range dir {
			filePath := path.Join(aPath, file.Name())
			files = append(files, filePath)
			if file.IsDir() {
				_ = queue.EnQueue(filePath)
			}
		}

	}
	return files, _err
}
