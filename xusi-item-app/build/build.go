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

// Xusi 项目构建工具
// 需要将所有源码作为静态资源打包一份到编译结果的机器码文件中
// 依赖于xusi的asset包
// 构建完成的静态文件统一由xusi-app在main包下建立xusi_build.go文件向asset包添加
package build

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"xusi-projects/xusi-framework/core/util"
)

// Xusi 构建者接口实现
type Builder interface {
	Build(projectPath string)
}

// 开始构建工作
func Build(projectPath string) error {
	// 先获取本项目中所有包含的文件
	// 记录每个文件的信息
	// 记录完成后对文件进行加密
	xFiles := util.GetFilesAll(projectPath)
	// 存储资产的go文件内容
	goTempFile := `
	package main
	
	import (
		"xusi-projects/xusi-framework/core/asset"
	)
	
	func init(){
	
	`
	for _, xFile := range xFiles {
		// 排除文件夹
		if xFile.IsDir {
			continue
		}
		// 排除指定后缀
		if strings.Contains(xFile.Name, ".exe") {
			continue
		}
		// 输出构建文件信息
		fmt.Println("building file -> " + strings.ReplaceAll(xFile.FullName, "/", "\\"))
		// 读取文件内容
		b, err := util.ReadFileContent(xFile.FullName)
		if err != nil {
			return err
		}
		// 并加密压缩文件内容
		gzip, err := util.GZipCompress(b)
		if err != nil {
			return err
		}
		// 拼接资产go
		goTempFile += `
		asset.Add("` + strings.ReplaceAll(xFile.FullName, "\\", "/") + `",asset.Assets{
			Name:		"` + strings.ReplaceAll(xFile.Name, "\\", "/") + `",
			Content:	"` + util.EncryptBase64(gzip.Next(gzip.Len())) + `",
			FileName:	"` + xFile.Name + `",
			FileType:	"` + strings.ReplaceAll(xFile.Ext, ".", "") + `",
			FullName:	"` + strings.ReplaceAll(xFile.FullName, "\\", "/") + `",
		})

		`
	}

	// 输出资产go文件
	goTempFile += `}`
	err := util.WriteToFile(projectPath+"/xusi_build.go", []byte(goTempFile))
	if err != nil {
		return err
	}
	// 构建
	cmd := exec.Command("go", "build", "xdoc.go", "xusi_build.go")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	cmd.Wait()

	// 移除
	time.Sleep(1 * time.Second)
	err = os.Remove(projectPath + "/xusi_build.go")
	if err != nil {
		return err
	}

	return nil
}
