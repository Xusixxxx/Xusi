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
	"net/url"
	"xusi-projects/xusi-framework/core/logger"
)

/* XusiFunc ->
    @describe 对字符串进行URL编码，返回编码结果
    @param content string 需要编码的内容
<- End */
func UrlEncoder(content string) string {
	return url.QueryEscape(content)
}

/* XusiFunc ->
    @describe 对已进行URL编码的字符串进行解码，返回解码内容
    @param content string 需要解码的内容
<- End */
func UrlDecoder(content string) string {
	result, err := url.QueryUnescape(content)
	if err != nil {
		logger.Error(err)
	}
	return result
}
