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

package xweb

import (
	"net/http"
	"strconv"
	"strings"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util/xurl"
	"xusi-projects/xusi-framework/xweb/httplibs"
	"xusi-projects/xusi-framework/xweb/static"
)

// 请求处理者
type requestHandler struct {
	*context                // 请求上下文
	RealRouteAddress string // 真实路由地址
}

// 初始化请求处理者
func (requestHandler *requestHandler) init(responseWriter http.ResponseWriter, request *http.Request) {
	requestHandler.context = &context{
		Http: struct {
			*http.Request
			http.ResponseWriter
		}{request, responseWriter},
		StatusCode:   httplibs.CODE_200,
		RouterParams: map[string]string{},
	}
}

// HTTP招待
// 所有的HTTP均经过此处进行处理
func (requestHandler *requestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// 初始化请求处理者
	requestHandler.init(responseWriter, request)
	// 对请求的URL进行解析
	// 获取URL中携带的路由参数
	// 对路由进行校验
	// 构建请求者
	url := requestHandler.analysisURL(request)
	requestHandler.analysisRouteParames(url)
	if requestHandler.routeRuleCheck(url, request) == httplibs.CODE_200 {
		// 解析From表单
		request.ParseForm()
		xWeb.router.Table[requestHandler.RealRouteAddress].Function(requestHandler.context)
	} else {
		switch requestHandler.context.StatusCode {
		case httplibs.CODE_404:
			requestHandler.context.Http.ResponseWriter.Write([]byte(static.PAGE_404))
		case httplibs.CODE_500:
			requestHandler.context.Http.ResponseWriter.Write([]byte(static.PAGE_500))
		}
	}
	requestHandler.printRequestInfo(requestHandler.context, request)
}

// 请求结果打印
func (requestHandler *requestHandler) printRequestInfo(ctx *context, request *http.Request) {
	var statusColor string
	switch ctx.StatusCode {
	case httplibs.CODE_200:
		statusColor = logger.GreenBg
	case httplibs.CODE_301, httplibs.CODE_302, httplibs.CODE_304, httplibs.CODE_307:
		statusColor = logger.YellowBg
	default:
		statusColor = logger.RedBg
	}
	logger.Info(
		logger.MagentaBg, request.Method, logger.Reset,
		statusColor, strconv.Itoa(ctx.StatusCode), logger.Reset,
		logger.Yellow, request.Host, logger.Reset,
		logger.Blue, xurl.UrlDecoder(request.URL.String()), logger.Reset,
	)

	// 如果为dev模式，输出请求详细
	if Conf.RunMode == httplibs.RUNMODE_DEV {
		var headerInfo, fromInfo string
		if len(request.Header) > 0 {
			for k, v := range request.Header {
				headerInfo += "\t" + k + " = " + v[0] + "\n"
			}
		}
		request.ParseForm()
		if len(request.Form) > 0 {
			for k, v := range request.Form {
				fromInfo += "\t" + k + " = " + v[0] + "\n"
			}
		}
		logger.Debug(
			logger.Yellow, "xusi http request info :\nHeader >>\n"+headerInfo, logger.Reset,
			logger.Yellow, "\nFrom >>\n"+fromInfo, logger.Reset,
		)
	}
}

// 路由规则校验，返回状态码
func (requestHandler *requestHandler) routeRuleCheck(url string, request *http.Request) int {
	if _, ok := xWeb.router.Table[url]; ok {
		// 是否通过
		isPass := false
		logger.Debug("check route : " + xurl.UrlDecoder(url))
		// 遍历所有该路由所支持的方法
		for _, value := range xWeb.router.Table[url].Method {
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
			requestHandler.StatusCode = httplibs.CODE_415
		}
	} else {
		logger.Debug("not found route : " + xurl.UrlDecoder(request.URL.String()))
		requestHandler.StatusCode = httplibs.CODE_404
	}
	return requestHandler.StatusCode
}

// 解析URL，防止请求被拒
func (requestHandler *requestHandler) analysisURL(request *http.Request) string {
	// 如果请求的方法不被允许或不存在
	// 解析url，避免被拒
	key := xurl.UrlDecoder(request.URL.String())
	// 如果末尾有斜杠，则移除
	if string([]byte(key)[len(key)-1]) == "/" && key != "/" {
		key = string([]byte(key)[0 : len(key)-1])
	}
	// 是否为携带参数
	if strings.Contains(key, "?") {
		key = strings.Split(key, "?")[0]
	}
	return key
}

// 解析路由参数
func (requestHandler *requestHandler) analysisRouteParames(url string) {
	// 首先检查是否为正常url
	// 如果未携带任何参数，但是绑定了带参数的路由
	// 那么应该请求到空参数的路由函数
	// 否则则视为有参数，进入解析
	if _, ok := xWeb.router.Table[url]; ok {
		requestHandler.RealRouteAddress = url
	} else {
		// 解析路由参数
		// 取到每一段切片，并排除第一段空切片
		slice := strings.Split(url, "/")
		slice = slice[1:len(slice)]
		// 遍历所有路由进行比对
		for k, v := range xWeb.router.Table {
			// 对遍历的pattern切片
			// 属性拼接再切片
			kTemp := k
			for _, v := range v.Params {
				kTemp += "/{" + v[1] + "}"
			}
			rSlice := strings.Split(kTemp, "/")
			rSlice = rSlice[1:len(rSlice)]
			// 如果传入的请求url切片量大于了该路由切片量不等，则不对应，进入下一个
			if len(slice) != len(rSlice) {
				continue
			} else {
				// 如果第一段切片不通，则不对应，进入下一个
				if slice[0] != rSlice[0] {
					continue
				} else {
					// 切片1对比切片1，切片2对比切片2
					// 属性段的索引需要由i减去非属性段数量
					paramI := 0
					for i, rs := range rSlice {
						// 检测当前是否为属性段
						isParamSlice := false
						for _, param := range v.Params {
							if param[0] == strconv.Itoa(i-paramI) && (i < len(slice)) && (rs != slice[i]) {
								isParamSlice = true
								break
							}
						}
						// 如果为属性段，这段则通过，开始匹配下一段，并且去掉属性段
						// 如果不为属性段，则匹配是否相同
						if isParamSlice {
							paramI++
							url = strings.ReplaceAll(url, "/"+slice[i], "")
						} else {
							if slice[i] != rs {
								break
							}
						}
					}
				}
			}
		}
		requestHandler.RealRouteAddress = url
		logger.Debug("parse url >> " + url)
	}
}
