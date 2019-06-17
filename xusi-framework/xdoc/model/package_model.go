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

/* XusiPackage ->
    @describe xdoc模型
<- End */
package model

/* XusiStrcut ->
   @describe 包模型
*/
type PackageModel struct {
	Name     string                 // $describe 名称
	Describe string                 // $describe 描述
	Const    map[string]ConstModel  // $describe 常量
	Struct   map[string]StructModel // $describe 结构体
	Func     map[string]FuncModel   // $describe 函数
} // -< End

/* XusiFunc ->
    @describe 是否为一个不包含任何公开内容的包
<- End */
func (this PackageModel) IsNull() bool {
	if len(this.Func) == 0 && len(this.Struct) == 0 && len(this.Const) == 0 {
		return true
	} else {
		return false
	}
}
