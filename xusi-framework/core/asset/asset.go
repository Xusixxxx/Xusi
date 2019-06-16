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

// 静态资产文件
package asset

import "xusi-projects/xusi-framework/core/util"

// 资产菜单
var AssetsMenu map[string]Assets

func init() {
	AssetsMenu = map[string]Assets{}
}

// 资产信息
type Assets struct {
	Name     string // 资产名称
	Content  string // 资产内容 （Base64加密）
	FileName string // 文件名称
	FileType string // 文件类型
	FullName string // 完整路径
}

// 添加资产
func Add(key string, assets Assets) {
	AssetsMenu[key] = assets
}

// 根据Key获取资产
func GetAssets(key string) Assets {
	return AssetsMenu[key]
}

// 根据Key获取资产内容
func GetContext(key string) ([]byte, error) {
	base64Data, err := util.DecodedBase64(AssetsMenu[key].Content)
	if err != nil {
		return nil, err
	}
	return util.UnGZipCompress(base64Data)
}

// 获取资产内容
func (assets Assets) GetContext() ([]byte, error) {
	base64Data, err := util.DecodedBase64(assets.Content)
	if err != nil {
		return nil, err
	}
	return util.UnGZipCompress(base64Data)
}
