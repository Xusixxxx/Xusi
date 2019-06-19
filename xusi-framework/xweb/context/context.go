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
    @describe 对XWeb上下文的封装，允许更为快捷的对上下文进行操作
<- End */
package context

import (
	"net/http"
	"net/url"
)

/* XusiStrcut ->
   @describe 请求上下文，包含了对Request和ResponseWriter的封装，以及一些特殊属性
*/
type Context struct {
	Http struct {
		*http.Request
		http.ResponseWriter
	}
	// diy属性
	StatusCode   int               // $describe 请求的结果状态码
	RouterParams map[string]string // $describe 存在于url之中的路由参数，如：/xusi/{id}/{name}
} // -< End

/* XusiFunc ->
    @describe 将字符串写入响应体
    @param content string 字符串
<- End */
func (ctx Context) WirteString(content string) {
	ctx.Http.ResponseWriter.Write([]byte(content))
}

/* XusiFunc ->
    @describe 获取请求的Header信息
<- End */
func (ctx Context) GetHeader() http.Header {
	return ctx.Http.Request.Header
}

/* XusiFunc ->
    @describe 获取请求的From信息
<- End */
func (ctx Context) GetFrom() url.Values {
	return ctx.Http.Request.Form
}
