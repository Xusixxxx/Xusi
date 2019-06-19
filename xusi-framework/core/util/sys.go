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

// 获取工作路径
func GetWorkPath() string {
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return ""
	}
	return path
}

// 获取程序当前运行目录
func GetCurrentPath() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path, nil
}

// 获取程序当前运行路径（不抛出error）
func GetCurrentPathNoErr() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

// 获取当前程序执行路径
func GetRunPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}
