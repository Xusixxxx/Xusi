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
    @describe 路由器的根基，包含了一个符合Xusi设计概念的路由器的通用基础功能
<- End */
package basic

import (
	"fmt"
	"reflect"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/net/context"
	"xusi-projects/xusi-framework/core/net/httplibs"
)

// 添加路由到路由器
// 路由器允许几种情况
// 多对多对多
// 一个路由可以允许多种请求方法，可以指向多个处理函数(指向的处理函数可拓展权重)
// 多个路由可以允许多种请求方法，并且指向一个或多个处理函数(指向的处理函数可拓展权重)
/* XusiFunc ->
    @describe 将路由添加到路由器中，该路由允许多个路由地址多赢多个允许的请求方法类型，并且允许拥有多个处理函数
    @param patterns Patterns 路由地址数组
    @param methods Methods 路由允许的方法类型数组
    @param functions Functions 路由处理函数数组
<- End */
func (router *Router) AddRoute(patterns Patterns, methods Methods, functions Functions) {
	for _, pattern := range patterns {
		router.RLock()
		// 检测路由是否已存在，如果存在则添加，否则新增
		routerTableItem := router.Table[pattern]
		if _, ok := router.Table[pattern]; ok {
			// 附加允许的请求方法
			for _, method := range methods {
				// 检测是否已有存在的请求方法，有则忽略
				isEx := false
				for _, value := range router.Table[pattern].Methods {
					if method == value {
						isEx = true
						break
					}
				}
				if !isEx {
					routerTableItem.Methods = append(routerTableItem.Methods, method)
					logger.Debug(fmt.Sprintf("addtion route method by %s%s%s , %s%s%s", logger.CyanBg, pattern, logger.Reset, logger.Cyan, method, logger.Reset))
				}
			}
			// 附加处理函数
			for _, function := range functions {
				// 检测是否已有存在的处理函数，有则忽略
				isEx := false
				for _, value := range router.Table[pattern].Functions {
					if reflect.ValueOf(function) == reflect.ValueOf(value) {
						isEx = true
						break
					}
				}
				if !isEx {
					routerTableItem.Functions = append(routerTableItem.Functions, function)
					logger.Debug(fmt.Sprintf("addtion route function by %s%s%s , %s%s%s", logger.CyanBg, pattern, logger.Reset, logger.Cyan, reflect.ValueOf(function).String(), logger.Reset))
				}
			}
			router.Table[pattern] = routerTableItem
		} else {
			router.Table[pattern] = RouteTableItem{
				Patterns:  patterns,
				Methods:   methods,
				Functions: functions,
			}
			logger.Debug("add new route", pattern, methods, reflect.ValueOf(functions).String())
		}
		router.RUnlock()
	}
}

func (router *Router) Add(patterns Patterns, methods Methods, function func(*context.Context)) {
	router.AddRoute(patterns, methods, Functions{function})
}
func (router *Router) GetArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_GET}, functions)
}
func (router *Router) Get(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_GET}, Functions{function})
}
func (router *Router) PostArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_POST}, functions)
}
func (router *Router) Post(pattern string, function func(ctx *context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_POST}, Functions{function})
}
func (router *Router) PutArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_PUT}, functions)
}
func (router *Router) Put(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_PUT}, Functions{function})
}
func (router *Router) DeleteArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_DELETE}, functions)
}
func (router *Router) Delete(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_DELETE}, Functions{function})
}
func (router *Router) HeadArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_HEAD}, functions)
}
func (router *Router) Head(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_HEAD}, Functions{function})
}
func (router *Router) TraceArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_TRACE}, functions)
}
func (router *Router) Trace(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_TRACE}, Functions{function})
}
func (router *Router) ConnectArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_CONNECT}, functions)
}
func (router *Router) Connect(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_CONNECT}, Functions{function})
}
func (router *Router) OptionsArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{httplibs.METHOD_OPTIONS}, functions)
}
func (router *Router) Options(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{httplibs.METHOD_OPTIONS}, Functions{function})
}
func (router *Router) AllArr(patterns Patterns, functions Functions) {
	router.AddRoute(patterns, Methods{
		httplibs.METHOD_GET,
		httplibs.METHOD_POST,
		httplibs.METHOD_PUT,
		httplibs.METHOD_DELETE,
		httplibs.METHOD_HEAD,
		httplibs.METHOD_TRACE,
		httplibs.METHOD_CONNECT,
		httplibs.METHOD_OPTIONS,
	}, functions)
}
func (router *Router) All(pattern string, function func(*context.Context)) {
	router.AddRoute(Patterns{pattern}, Methods{
		httplibs.METHOD_GET,
		httplibs.METHOD_POST,
		httplibs.METHOD_PUT,
		httplibs.METHOD_DELETE,
		httplibs.METHOD_HEAD,
		httplibs.METHOD_TRACE,
		httplibs.METHOD_CONNECT,
		httplibs.METHOD_OPTIONS,
	}, Functions{function})
}
