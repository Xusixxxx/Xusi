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

package context

import "net/http"

type Context struct {
	Http struct {
		*http.Request
		http.ResponseWriter
	}
	StatusCode int // 状态码
}

// 输出字符串
func (context Context) Write(str string) {
	context.Http.ResponseWriter.WriteHeader(context.StatusCode)
	context.Http.ResponseWriter.Write([]byte(str))
}

// 输出字符串
func (context Context) WriteByte(byte []byte) {
	context.Http.ResponseWriter.Write(byte)
}
