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

import (
	"net/http"
	"net/url"
)

type Context struct {
	Http struct {
		*http.Request
		http.ResponseWriter
	}
	// diy属性
	StatusCode   int               // 状态码
	routerParams map[string]string // 路由参数，存在于url之中
}

// 将字符串写入响应体
// 发送响应体及状态码
func (ctx Context) WirteString(content string) {
	ctx.Http.ResponseWriter.Write([]byte(content))
}

// 获取头部
func (ctx Context) GetHeader() http.Header {
	return ctx.Http.Request.Header
}

// 获取表单
func (ctx Context) GetFrom() url.Values {
	return ctx.Http.Request.Form
}
