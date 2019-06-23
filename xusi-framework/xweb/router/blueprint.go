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
    @describe 路由器蓝图，可根据该蓝图打造各式各样的路由器
<- End */
package router

import (
	"net/http"
	"xusi-projects/xusi-framework/xweb/router/basic"
)

type Blueprint interface {
	// 查找路由
	Find(string, *http.Request) basic.RouteTableItem
	// 呼叫路由
	Call(basic.RouteTableItem) []func(responseWriter http.ResponseWriter, r *http.Request)
}

func Find(blueprint Blueprint, pattern string, request *http.Request) basic.RouteTableItem {
	return blueprint.Find(pattern, request)
}

func Call(blueprint Blueprint, routeTableItem basic.RouteTableItem) []func(responseWriter http.ResponseWriter, r *http.Request) {
	return blueprint.Call(routeTableItem)
}
