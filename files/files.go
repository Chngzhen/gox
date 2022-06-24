package files

import (
	"io/ioutil"
	"os"
	"strings"
)

var pathSeparator = string(os.PathSeparator)

// ExtendsIn 检查文件是否以指定后缀中的某一个结尾。
func ExtendsIn(fileName string, extends []string) bool {
	for _, ext := range extends {
		if strings.HasSuffix(fileName, "."+ext) {
			return true
		}
	}
	return false
}

// CheckDir 检查目录是否存在。若create为true，则会创建不存在的目录。
func CheckDir(dirPath string, create bool) (bool, error) {
	if _, ok := IsExisted(dirPath); !ok {
		if !create {
			return false, nil
		}

		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return false, err
		}
	}
	return true, nil
}

// IsExisted 判断文件是否存在。注意，若无权访问目标文件，依旧会返回true。
func IsExisted(filePath string) (*os.FileInfo, bool) {
	fileInfo, err := os.Lstat(filePath)
	return &fileInfo, !os.IsNotExist(err)
}

// RetrieveFiles 检索指定目录下的子孙文件。
func RetrieveFiles(dir string, filePaths chan<- string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := dir + pathSeparator + file.Name()
		if file.IsDir() {
			err = RetrieveFiles(filePath, filePaths)
			if err != nil {
				return err
			}
		} else {
			filePaths <- filePath
		}
	}
	return nil
}
