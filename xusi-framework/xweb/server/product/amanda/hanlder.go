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

package amanda

import (
	"net/http"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xweb/httplibs"
)

// 请求处理者
type requestHandler struct {
	http.ResponseWriter     // 响应流
	*http.Request           // HTTP请求
	StatusCode          int // 请求状态码
}

// 请求初始化
func (requestHandler *requestHandler) init(responseWriter http.ResponseWriter, request *http.Request) {
	requestHandler.ResponseWriter = responseWriter
	requestHandler.Request = request
	requestHandler.StatusCode = httplibs.CODE_200
	if request.ParseForm() != nil {
		logger.Warn("parse request form unsuccessful, maybe form unusable!")
	}
}

// HTTP请求招待
func (requestHandler *requestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestHandler.init(responseWriter, request)

}
