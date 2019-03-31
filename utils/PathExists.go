package utils

import (
	"fmt"
	"os"
)

//PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Printf("has dir![%v]\n", path)
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Printf("no dir![%v]\n", path)
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
		return false, nil
	}
	return false, err
}
