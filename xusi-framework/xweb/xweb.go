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

// xweb framework
// 用于提供给外部使用的功能
package xweb

import (
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
)

/* XusiFunc ->
    @describe 启动XWeb HTTP服务器
    @param params ...string 启动的参数，通常仅传入一个端口
<- End */
func Run(params ...string) {
	// 注册所有路由
	registryRouters()
	// 启动http server
	if len(params) > 0 && params[0] != "" {
		listen(conf.network, params[0])
	} else {
		listen()
	}
}

/* XusiFunc ->
    @describe 添加指定允许的方法类型的路由
    @param pattern string 路由地址
    @param method []string 允许的方法类型
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Add(pattern string, method []string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, method, function)
}

/* XusiFunc ->
    @describe 添加Get方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Get(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_GET}, function)
}

/* XusiFunc ->
    @describe 添加Post方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Post(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_POST}, function)
}

/* XusiFunc ->
    @describe 添加Put方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Put(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_PUT}, function)
}

/* XusiFunc ->
    @describe 添加Delete方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Delete(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_DELETE}, function)
}

/* XusiFunc ->
    @describe 添加Head方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Head(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_HEAD}, function)
}

/* XusiFunc ->
    @describe 添加Trace方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Trace(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_TRACE}, function)
}

/* XusiFunc ->
    @describe 添加Connect方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Connect(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_CONNECT}, function)
}

/* XusiFunc ->
    @describe 添加Options方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func Options(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_OPTIONS}, function)
}

/* XusiFunc ->
    @describe 添加允许所有方法类型的路由
    @param pattern string 路由地址
    @param function func(ctx *context.Context) 路由处理函数
<- End */
func All(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{
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

// 设置运行模式
/* XusiFunc ->
    @describe 设置XWeb HTTP服务器的运行模式
    @param mode string 运行的模式
<- End */
func SetRunMode(mode string) {
	conf.mode = mode
	logger.Conf.Mode = mode
}
