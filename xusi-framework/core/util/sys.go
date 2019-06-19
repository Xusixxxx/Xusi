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
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"xusi-projects/xusi-framework/core/logger"
)

/* XusiFunc ->
    @describe 获取项目的工作路径，返回string
<- End */
func GetWorkPath() string {
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return ""
	}
	return path
}

/* XusiFunc ->
    @describe 获取程序当前运行的目录，返回目录路径和错误信息
<- End */
func GetCurrentPath() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path, nil
}

/* XusiFunc ->
    @describe 获取程序当前运行的目录，返回目录路径，不抛出错误信息
<- End */
func GetCurrentPathNoErr() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

/* XusiFunc ->
    @describe 获取当前程序的运行路径，返回目录路径及错误信息
<- End */
func GetRunPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}
