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

// xweb请求处理者
package xweb

import (
	"net/http"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
)

// 请求处理者
type requestHandler struct {
	*context.Context
}

// 统一函数处理
// 所有函数都将经过这个
func (r *requestHandler) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// 如果请求的方法不被允许或不存在
	if _, ok := xrouterInstance.routerTable.Table[request.URL.String()]; ok {
		isPass := false
		logger.Debug("check route : " + request.URL.String())
		for _, value := range xrouterInstance.routerTable.Table[request.URL.String()].Method {
			logger.Debug("router take in method : ", value, ", input method : ", request.Method)
			if value == request.Method {
				isPass = true
				logger.Debug("verification by")
				break
			}
		}
		if isPass == false {
			logger.Debug("fail verification")
			r.StatusCode = httplib.CODE_415
		}
	} else {
		logger.Debug("not found route : " + request.URL.String())
		r.StatusCode = httplib.CODE_404
	}

	request.ParseForm()
}
