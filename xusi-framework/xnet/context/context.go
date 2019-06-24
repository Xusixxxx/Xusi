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
	"encoding/json"
	"net/http"
	"strconv"
	"xusi-projects/xusi-framework/core/logger"
)

/* XusiStrcut ->
   @describe 请求上下文，包含了对Request和ResponseWriter的封装，以及一些特殊属性
*/
type Context struct {
	*http.Request
	http.ResponseWriter
	StatusCode int // 请求状态码
} // -< End

/* XusiFunc ->
    @describe 将字符串写入响应体
    @param content string 字符串
<- End */
func (ctx *Context) XWriteString(content string) {
	ctx.ResponseWriter.WriteHeader(ctx.StatusCode)
	i, err := ctx.ResponseWriter.Write([]byte(content))
	if err == nil {
		logger.Debug("write data for " + strconv.Itoa(i) + " byte")
	} else {
		logger.Error(err)
	}
}

/* XusiFunc ->
    @describe 将结构转为json串写入响应体
    @param content string 字符串
<- End */
func (ctx *Context) XWriteJSON(value interface{}) {
	ctx.ResponseWriter.Header().Set("Content-type", "application/json;charset=UTF-8")
	ctx.ResponseWriter.WriteHeader(ctx.StatusCode)
	json, err := json.Marshal(value)
	if err == nil {
		i, err := ctx.ResponseWriter.Write(json)
		logger.Debug("write data for " + strconv.Itoa(i) + "byte")
		if err != nil {
			logger.Error(err)
		}
	} else {
		logger.Error(err)
	}
}
