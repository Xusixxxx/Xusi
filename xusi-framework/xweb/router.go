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

// xweb 路由
package xweb

import (
	"fmt"
	"net/http"
	"strconv"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
	"xusi-projects/xusi-framework/xweb/router"
)

var xrouterInstance xrouter

func init() {
	xrouterInstance = xrouter{
		routerTable: router.CreateRouterTable(),
	}
}

// 路由器
type xrouter struct {
	routerTable router.RouterTable // 路由表
}

// 注册所有路由
// 默认状态码为200
func registryRouters() {
	logger.Info("xusi http server router registration...")
	for _, itemTemp := range xrouterInstance.routerTable.Table {
		// 如果没有这行代码，不知道为什么会导致这样的情况
		// 不管访问任一路由，调用的都为最后一个注册的路由处理函数
		item := itemTemp
		http.HandleFunc(item.Pattern, func(responseWriter http.ResponseWriter, request *http.Request) {
			ctx := &context.Context{
				Http: struct {
					*http.Request
					http.ResponseWriter
				}{request, responseWriter},
				StatusCode: httplib.CODE_200,
			}
			// 执行全局处理函数
			rHandler := &requestHandler{ctx}
			rHandler.serveHTTP(responseWriter, request)
			// 如果在全局处理函数检测完毕后状态码仍是200，那么执行路由处理函数
			if ctx.StatusCode == httplib.CODE_200 {
				logger.Debug("run route handler -> ", item.Function)
				item.Function(ctx)
			} else {
				ctx.Http.ResponseWriter.Write([]byte("xusi failed request : " + strconv.Itoa(ctx.StatusCode)))
			}
			// 请求结果打印
			// 打印请求结果
			var statusColor string
			switch ctx.StatusCode {
			case httplib.CODE_200:
				statusColor = logger.GreenBg
			case httplib.CODE_301, httplib.CODE_302, httplib.CODE_304, httplib.CODE_307:
				statusColor = logger.YellowBg
			default:
				statusColor = logger.RedBg
			}
			logger.Info(
				logger.MagentaBg, request.Method, logger.Reset,
				statusColor, strconv.Itoa(ctx.StatusCode), logger.Reset,
				logger.Yellow, request.Host, logger.Reset,
				logger.Blue, request.URL.String(), logger.Reset,
			)
		})
		logger.Debug(fmt.Sprintf("router registry <- %s%s%s", logger.YellowBg, item.Pattern, logger.Reset))
	}
}

// 实现接口
type routerImplement interface {
	// 特定类型添加路由
	Add(pattern string, method []string, function func(ctx *context.Context))
	// 添加各类型路由
	Get(pattern string, function func(ctx *context.Context))
	Post(pattern string, function func(ctx *context.Context))
	Put(pattern string, function func(ctx *context.Context))
	Delete(pattern string, function func(ctx *context.Context))
	Head(pattern string, function func(ctx *context.Context))
	Trace(pattern string, function func(ctx *context.Context))
	Connect(pattern string, function func(ctx *context.Context))
	Options(pattern string, function func(ctx *context.Context))
	// 添加所有类型路由
	All(pattern string, function func(ctx *context.Context))
}

func (router xrouter) Add(pattern string, method []string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, method, function)
}

func (router xrouter) Get(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_GET}, function)
}

func (router xrouter) Post(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_POST}, function)
}

func (router xrouter) Put(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_PUT}, function)
}

func (router xrouter) Delete(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_DELETE}, function)
}

func (router xrouter) Head(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_HEAD}, function)
}

func (router xrouter) Trace(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_TRACE}, function)
}

func (router xrouter) Connect(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_CONNECT}, function)
}

func (router xrouter) Options(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{httplib.METHOD_OPTIONS}, function)
}

func (router xrouter) All(pattern string, function func(ctx *context.Context)) {
	router.routerTable.Add(pattern, []string{
		httplib.METHOD_GET,
		httplib.METHOD_POST,
		httplib.METHOD_PUT,
		httplib.METHOD_DELETE,
		httplib.METHOD_CONNECT,
		httplib.METHOD_HEAD,
		httplib.METHOD_OPTIONS,
		httplib.METHOD_TRACE,
	}, function)
}
