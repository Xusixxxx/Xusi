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
	"strings"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util"
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
	// 解析url，避免被拒
	key := util.UrlDecoder(request.URL.String())
	// 是否为携带参数
	if strings.Contains(key, "/{") {
		//		/hello/{id}/{name}
		//		/hello/{id}/{name}/xusi/{age}

	} else {
		key = strings.Split(key, "?")[0]
	}
	logger.Debug("parse url >> " + key)
	if _, ok := xrouterInstance.routerTable.Table[key]; ok {
		// 是否通过
		isPass := false
		logger.Debug("check route : " + util.UrlDecoder(request.URL.String()))
		// 遍历所有该路由所支持的方法
		for _, value := range xrouterInstance.routerTable.Table[key].Method {
			logger.Debug("router take in method : ", value, ", input method : ", request.Method)
			// 如果找到对应被允许的方法，那么则通过验证
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
		logger.Debug("not found route : " + util.UrlDecoder(request.URL.String()))
		r.StatusCode = httplib.CODE_404
	}

	// 解析From表单
	request.ParseForm()

}
