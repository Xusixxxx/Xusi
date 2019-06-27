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
	"reflect"
	"strconv"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/net/context"
	"xusi-projects/xusi-framework/core/net/httplibs"
	basic2 "xusi-projects/xusi-framework/core/net/server/basic"
	"xusi-projects/xusi-framework/core/net/static"
	"xusi-projects/xusi-framework/core/util/xurl"
	"xusi-projects/xusi-framework/xnet"
)

// 请求处理者
type requestHandler struct {
	*context.Context
}

// 请求初始化
func (requestHandler *requestHandler) init(responseWriter http.ResponseWriter, request *http.Request) {
	requestHandler.Context = &context.Context{}
	requestHandler.ResponseWriter = responseWriter
	requestHandler.Request = request
	requestHandler.StatusCode = httplibs.CODE_200
	if request.ParseForm() != nil {
		logger.Warn("parse request form unsuccessful, maybe form unusable!")
	}
}

// HTTP请求招待
func (requestHandler *requestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// 对请求进行初始化
	requestHandler.init(responseWriter, request)
	// 调用路由器找到路由处理函数
	functions := xnet.Load().Router.Find(requestHandler.Context)
	// 如果路由未找到
	if len(functions) == 0 {
		requestHandler.StatusCode = httplibs.CODE_404
	}
	// 执行路由处理函数
	for _, function := range functions {
		work := Job{serload: Serload{pri: reflect.TypeOf(function).String()}}
		function(requestHandler.Context)
		JobQueue <- work
	}
	// 处理请求结果信息
	requestHandler.requestEndInfo(len(functions))
}

// 处理请求结束信息
func (requestHandler *requestHandler) requestEndInfo(funcCount int) {
	// 是否需要输出欢迎页面(理应说欢迎页面也应该属于一个路由，但是为了保证注册"/"时不会叠加，所以单独设置返回)
	if requestHandler.URL.String() == "/" && funcCount == 0 {
		requestHandler.XWriteString(static.PAGE_WELCOME)
		goto PrintColor
	}
	// 是否需要输出错误页面
	switch requestHandler.StatusCode {
	case httplibs.CODE_404:
		requestHandler.XWriteString(static.PAGE_404)
	case httplibs.CODE_500:
		requestHandler.XWriteString(static.PAGE_500)
	}

PrintColor:
	// 定义输出信息存储变量
	// 针对不同内容输出不同色彩
	var statusColor string
	switch requestHandler.StatusCode {
	case httplibs.CODE_200:
		statusColor = logger.GreenBg
	case httplibs.CODE_301, httplibs.CODE_302, httplibs.CODE_304, httplibs.CODE_307:
		statusColor = logger.YellowBg
	default:
		statusColor = logger.RedBg
	}
	logger.Info(
		logger.MagentaBg, requestHandler.Method, logger.Reset,
		statusColor, strconv.Itoa(requestHandler.StatusCode), logger.Reset,
		logger.Yellow, requestHandler.Host, logger.Reset,
		logger.Blue, xurl.UrlDecoder(requestHandler.URL.String()), logger.Reset,
	)

	// 如果为dev模式，输出请求详细
	if xnet.RunMode() == basic2.RUN_MODE_DEV {
		var headerInfo, fromInfo string
		if len(requestHandler.Request.Header) > 0 {
			for k, v := range requestHandler.Request.Header {
				headerInfo += "\t" + k + " = " + v[0] + "\n"
			}
		}
		if len(requestHandler.Request.Form) > 0 {
			for k, v := range requestHandler.Request.Form {
				fromInfo += "\t" + k + " = " + v[0] + "\n"
			}
		}
		logger.Debug(
			logger.Yellow, "xusi http request info :\nHeader >>\n"+headerInfo, logger.Reset,
			logger.Yellow, "\nFrom >>\n"+fromInfo, logger.Reset,
		)
	}
}
