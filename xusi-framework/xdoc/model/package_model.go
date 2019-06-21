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

package model

import (
	"strings"
	"xusi-projects/xusi-framework/core/asset"
)

/* XusiStrcut ->
   @describe 用于支撑XDoc运行的包模型
*/
type PackageModel struct {
	Name     string                 // $describe 包名
	IsRoot   bool                   // $describe 是否为顶级包
	DirPath  string                 // $describe 包路径
	Describe string                 // $describe 包描述
	Const    map[string]ConstModel  // $describe 包内所包含的常量
	Struct   map[string]StructModel // $describe 包内所包含的结构体
	Func     map[string]FuncModel   // $describe 包内所包含的函数
	Exclude  []string               // $describe 被排除的会影响包真实引用路径的包名
} // -< End

/* XusiFunc ->
    @describe 是否为一个不包含任何对外公开内容的package
<- End */
func (this PackageModel) IsNull() bool {
	if len(this.Func) == 0 && len(this.Struct) == 0 && len(this.Const) == 0 {
		return true
	} else {
		return false
	}
}

/* XusiFunc ->
    @describe 获取package完整路径
<- End */
func (this PackageModel) GetPackagePath() string {
	// 去除文件名称
	// 并将文件夹名称改为包名
	slice := strings.Split(strings.ReplaceAll(this.DirPath, "/", "/\\"), "\\")
	slice = slice[0 : len(slice)-1]
	result := ""
	for _, value := range slice {
		result += value
	}
	result = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(result+this.Name, asset.ROOT_PATH, ""), "xusi-projects", ""), "xusi-projects/xusi-framework", "")
	// 排除空包
	for _, packageName := range this.Exclude {
		result = strings.ReplaceAll(result, "/"+packageName, "")
	}
	return result
}
