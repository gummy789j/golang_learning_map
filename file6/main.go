package main

import (
	"io/ioutil"
	"os"
)

// 判斷文件是否存在
func PathExists(path string) (bool, error) {
	_,err := os.Stat(path)
	// 若存在
	if err == nil {
		return true, nil
	}
	// 檢查其error 代表的是否是因為檔案不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 將file讀進來 再寫到其他file去
func main() {
	
	filePath1 := "./data/kkk.txt"
	b1, err := PathExists(filePath1)
	filePath2 := "./jjj.txt"
	b2, err := PathExists(filePath2)

	if err != nil {
		panic("Checking files does not Success !")
	}

	if !b1 || !b2 {
		panic("Files not Exist")
	}
	
	data, err := ioutil.ReadFile(filePath2)
	if err != nil {
		panic("Read file error !")
	}

	err = ioutil.WriteFile(filePath1,data,7555)
	if err != nil {
		panic("Write file error !")
	}

}