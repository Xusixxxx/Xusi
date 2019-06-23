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
	"net/http"
	"xusi-projects/xusi-framework/xweb/router/basic"
)

/* XusiFunc ->
    @describe 检查在路由器中方法是否被允许
    @param router basic.Router 路由器
    @param pattern string 路由地址
    @param request *http.Request HTTP请求
<- End */
func CheckMethod(router *basic.Router, pattern string, request *http.Request) bool {
	isPass := false
	for _, method := range router.Table[pattern].Methods {
		if method == request.Method {
			isPass = true
			break
		}
	}
	return isPass
}
