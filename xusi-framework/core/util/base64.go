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
    @describe Xusi-framework-core中所包含的通用开发工具包
<- End */
package util

import "encoding/base64"

/* XusiFunc ->
    @describe 对数据进行Base64加密，返回一个string类型的加密结果
    @param data []byte 需要加密的数据
<- End */
func EncryptBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

/* XusiFunc ->
    @describe 对已经进行Base64加密的数据进行解密，返回字节数组和错误信息
    @param data string 已被进行Base64加密的数据
<- End */
func DecodedBase64(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
