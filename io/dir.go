package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// if directory not exit, create it.
func MakeDirAll(dirPath string) (string, error) {

	exists, err := PathExists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		return dirPath, nil
	}
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return "", fmt.Errorf("error while creating directory,err:%s", err.Error())
	}
	return dirPath, nil
}

// check path exist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// try to delete empty dir. true: delete an empty dir, false: delete nothing.
func DeleteEmptyDir(dirPath string) (bool, error) {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		if strings.Contains(err.Error(), "The system cannot find") {
			return false, nil
		} else {
			return false, fmt.Errorf("occur error while reading %s %s", dirPath, err.Error())
		}
	}
	if len(dir) == 0 {
		//empty dir
		err = os.Remove(dirPath)
		if err != nil {
			return false, fmt.Errorf("occur error while deleting %s %s", dirPath, err.Error())
		}
		return true, nil
	}

	return false, nil
}

// delete empty dir recursive, delete until not empty.
func DeleteEmptyDirRecursive(dirPath string) (bool, error) {

	tmpPath := dirPath
	for {
		ok, err := DeleteEmptyDir(tmpPath)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
		dir := GetDirOfPath(tmpPath)
		tmpPath = dir
	}
}

// eg /var/www/xx.log -> /var/www
func GetDirOfPath(fullPath string) string {

	index1 := strings.LastIndex(fullPath, "/")
	//maybe windows environment
	index2 := strings.LastIndex(fullPath, "\\")
	index := index1
	if index2 > index1 {
		index = index2
	}

	return fullPath[:index]
}

// 判断 targetPath 是否是 basePath 的子目录
func IsSubPath(basePath, targetPath string) (bool, error) {
	absBase, err := filepath.Abs(basePath)
	if err != nil {
		return false, err
	}

	absTarget, err := filepath.Abs(targetPath)
	if err != nil {
		return false, err
	}

	rel, err := filepath.Rel(absBase, absTarget)
	if err != nil {
		return false, err
	}

	// rel 开头是 .. 表示不在 basePath 内部
	if strings.HasPrefix(rel, "..") || rel == "." {
		return false, nil
	}

	return true, nil
}
