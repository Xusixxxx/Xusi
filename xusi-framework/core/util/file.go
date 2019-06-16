// Copyright 2019 Xusixxxx Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

// Xusi文件信息结构
type XFile struct {
	Name       string      // 文件名
	Ext        string      // 文件后缀名
	FullName   string      // 绝对路径
	IsDir      bool        // 是否为文件夹
	Size       int64       // 文件大小
	Mode       os.FileMode // 文件的模式和权限位
	ModifyTime time.Time   // 文件修改时间
	Sys        interface{} // 底层数据来源（可以返回nil）
}

// 获取文件的XFile信息
func GetXFile(fInfo os.FileInfo) (XFile, error) {
	fullName, e := filepath.Abs(filepath.Dir(fInfo.Name()))
	if e != nil {
		return XFile{}, e
	} else {
		return XFile{
			Name:       fInfo.Name(),
			Ext:        path.Ext(fullName + "/" + fInfo.Name()),
			FullName:   fullName + "/" + fInfo.Name(),
			IsDir:      fInfo.IsDir(),
			Size:       fInfo.Size(),
			Mode:       fInfo.Mode(),
			ModifyTime: fInfo.ModTime(),
			Sys:        fInfo.Sys(),
		}, nil
	}
}

// 获取路径下的所有文件
func GetFilesAll(dirPath string) []XFile {
	result := ""
	getAllFiles(&result, dirPath)
	xfiles := []XFile{}
	json.Unmarshal([]byte("["+result+"]"), &xfiles)
	return xfiles
}

// 获取指定路径下的文件
func GetFiles(dirPath string) []XFile {
	result := ""
	files, _ := ioutil.ReadDir(dirPath)
	for _, v := range files {
		xFile, _ := GetXFile(v)
		json, _ := json.Marshal(xFile)
		if result == "" {
			result = result + string(json)
		} else {
			result = result + "," + string(json)
		}
	}
	xfiles := []XFile{}
	json.Unmarshal([]byte("["+result+"]"), &xfiles)
	return xfiles
}

// 递归获取
func getAllFiles(result *string, dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	for _, v := range files {
		xFile := getXFile(dirPath, v)
		json, _ := json.Marshal(xFile)
		if *result == "" {
			*result = *result + string(json)
		} else {
			*result = *result + "," + string(json)
		}

		if v.IsDir() {
			getAllFiles(result, dirPath+"/"+v.Name())
		}
	}
}

// 获取递归文件的XFile信息
func getXFile(rootPath string, fInfo os.FileInfo) XFile {
	return XFile{
		Name:       fInfo.Name(),
		Ext:        path.Ext(rootPath + "/" + fInfo.Name()),
		FullName:   rootPath + "/" + fInfo.Name(),
		IsDir:      fInfo.IsDir(),
		Size:       fInfo.Size(),
		Mode:       fInfo.Mode(),
		ModifyTime: fInfo.ModTime(),
		Sys:        fInfo.Sys(),
	}
}

// 读取文件内容
func ReadFileContent(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// 写入数据到文件中
// 不存在则会创建新文件
func WriteToFile(path string, data []byte) error {
	var f *os.File
	var err error

	if FileIsExist(path) { //如果文件存在
		f, err = os.OpenFile(path, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(path) //创建文件
	}
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	_, err = w.WriteString(string(data))
	if err != nil {
		return err
	}
	w.Flush()
	f.Close()
	return nil
}

// 文件是否存在
func FileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
