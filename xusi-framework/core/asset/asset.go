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
    @describe 静态资产管理
<- End */
package asset

import "xusi-projects/xusi-framework/core/util"

// 资产菜单
var AssetsMenu map[string]Assets

func init() {
	AssetsMenu = map[string]Assets{}
}

/* XusiStrcut ->
   @describe 静态资产模型
*/
type Assets struct {
	Name     string // $describe 资产名称
	Content  string // $describe 资产内容 （Base64加密）
	FileName string // $describe 文件名称
	FileType string // $describe 文件类型
	FullName string // $describe 完整路径
} // -< End

/* XusiFunc ->
    @describe 添加静态资产（该函数通常由xusi-framework调用）
    @param key string 键
    @param assets Assets 要添加的静态资产模型
<- End */
func Add(key string, assets Assets) {
	AssetsMenu[key] = assets
}

/* XusiFunc ->
    @describe 获取静态资产内容，返回 <kbd>[]byte</kbd> 和 <kbd>error</kbd>
<- End */
func (assets Assets) GetContext() ([]byte, error) {
	base64Data, err := util.DecodedBase64(assets.Content)
	if err != nil {
		return nil, err
	}
	return util.UnGZipCompress(base64Data)
}
