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

package xdoc_util

import (
	"strings"
	"xusi-projects/xusi-framework/core/asset"
)

// 资产文件内容格式化
// 保证每一行首尾没有空格，移除多余空行
func FormatAssets(assetFile asset.Assets) (string, error) {
	assetContent, err := assetFile.GetContext()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(assetContent), "\n")
	content := ""
	for _, line := range lines {
		if strings.ReplaceAll(line, " ", "") == "" {
			continue
		}
		content += strings.TrimSpace(line) + "\n"
	}
	return content, nil
}
