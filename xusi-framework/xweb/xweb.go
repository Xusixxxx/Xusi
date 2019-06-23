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
	"strconv"
	"sync"
	"xusi-projects/xusi-framework/xweb/httplibs"
	"xusi-projects/xusi-framework/xweb/static"
)

// 路由路由器装载，保证全局仅有一个路由器
var buildRouterOnce sync.Once

var xWeb *xweb

func init() {
	xWeb = &xweb{}
	// 装载路由器
	xWeb.loadRouter()
}

type xweb struct {
	router *router // 路由器
}

// 装载路由器到xweb，确保仅能装载一次
func (xweb *xweb) loadRouter() {
	buildRouterOnce.Do(func() {
		xweb.router = &router{
			Table: map[string]routerTableItem{},
		}
		// 初始化“/”根路由，确保任何请求都能被接收
		xweb.router.Table["/"] = routerTableItem{
			Method: []string{
				httplibs.METHOD_GET,
				httplibs.METHOD_POST,
				httplibs.METHOD_PUT,
				httplibs.METHOD_DELETE,
				httplibs.METHOD_CONNECT,
				httplibs.METHOD_HEAD,
				httplibs.METHOD_OPTIONS,
				httplibs.METHOD_TRACE,
			},
			Pattern: "/",
			Function: func(ctx *context) {
				ctx.WirteString(static.PAGE_WELCOME)
			},
			Exclude: "",
			Params:  map[string][]string{},
		}
	})
}

// 8080
// localhost, 8080
// tcp4, localhost, 8080
func Run(params ...string) {
	// 注册所有路由
	//router.registryRouters()
	// 启动http server
	switch len(params) {
	case 1:
		port, err := strconv.Atoi(params[0])
		if err != nil {
			panic(err)
		}
		Conf.Port = port
	case 2:
		port, err := strconv.Atoi(params[1])
		if err != nil {
			panic(err)
		}
		Conf.Address = params[0]
		Conf.Port = port
	case 3:
		port, err := strconv.Atoi(params[2])
		if err != nil {
			panic(err)
		}
		Conf.Network = params[0]
		Conf.Address = params[1]
		Conf.Port = port
	}
	listen()
}
