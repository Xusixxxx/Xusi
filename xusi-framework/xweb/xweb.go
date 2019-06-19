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

func Add(pattern string, method []string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, method, function)
}

func Get(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_GET}, function)
}

func Post(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_POST}, function)
}

func Put(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_PUT}, function)
}

func Delete(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_DELETE}, function)
}

func Head(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_HEAD}, function)
}

func Trace(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_TRACE}, function)
}

func Connect(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_CONNECT}, function)
}

func Options(pattern string, function func(ctx *context.Context)) {
	xrouterInstance.routerTable.Add(pattern, []string{httplib.METHOD_OPTIONS}, function)
}

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
func SetRunMode(mode string) {
	conf.mode = mode
	logger.Conf.Mode = mode
}
