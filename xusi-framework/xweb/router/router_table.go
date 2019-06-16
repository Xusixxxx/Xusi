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

// 路由池
package router

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
)

// 路由表
// 存储所有路由信息
// 带有读写锁，保证线程安全
// 允许一个映射地址拥有多种方法统一处理
type RouterTable struct {
	sync.RWMutex
	Table map[string]RouterTableItem
}

// 路由表项
type RouterTableItem struct {
	Method   []string                   // 接受的方法类型
	Pattern  string                     // 映射的地址
	Function func(ctx *context.Context) // 路由处理的函数
	Exclude  string                     // 被排除的方法链串
	Params   map[string][]string        // 路由中携带的参数 map[name]{[]string{index, name}}
}

// 创建一个路由表
func CreateRouterTable() RouterTable {
	routerTable := RouterTable{
		Table: map[string]RouterTableItem{},
	}
	// 初始化“/”根路由，确保任何请求都能被接收
	routerTable.Table["/"] = RouterTableItem{
		Method: []string{
			httplib.METHOD_GET,
			httplib.METHOD_POST,
			httplib.METHOD_PUT,
			httplib.METHOD_DELETE,
			httplib.METHOD_CONNECT,
			httplib.METHOD_HEAD,
			httplib.METHOD_OPTIONS,
			httplib.METHOD_TRACE,
		},
		Pattern: "/",
		Function: func(ctx *context.Context) {
			ctx.WirteString("Welcome Xusi!")
		},
		Exclude: "",
		Params:  map[string][]string{},
	}
	return routerTable
}

// 添加项(允许包含http协议外的method)
func (routerTable RouterTable) Add(patternTemp string, method []string, function func(ctx *context.Context)) {
	// copy
	pattern := patternTemp
	// 如果开头没有“/”则自动补上
	if string([]byte(pattern)[0]) != "/" {
		pattern = "/" + pattern
	}
	// 如果末尾为“/”则移除
	if string([]byte(pattern)[len(pattern)-1]) == "/" {
		pattern = string([]byte(pattern)[0 : len(pattern)-1])
	}
	// 如果为“\”则替换为“/”
	pattern = strings.ReplaceAll(pattern, "\\", "/")
	// 如果有多余空格移除
	pattern = strings.ReplaceAll(pattern, " ", "")
	// 路由参数解析
	// 切片分析每一段
	slice := strings.Split(pattern, "/")
	// 移除首个空切片
	slice = slice[1:len(slice)]
	params := ""
	pattern = ""
	for _, v := range slice {
		// 如果存在空段，则表示存在 “//”
		if util.IsEmptyString(v) {
			logger.Warn("route exception(exist null slice), may not work : ", patternTemp)
		}
		// 获取参数名
		if strings.Contains(v, "{") && strings.Contains(v, "}") {

			result := util.GetBetweenStr(v, "{", "}")
			ok := true
			// 如果为空，则表示存在不完整的标识符，可能仅有{，可能仅有}
			if util.IsEmptyString(result) {
				ok = false
				logger.Warn("route exception(incomplete identifier), may not work : ", patternTemp)
			}
			// 如果取值中还包含标识符
			if strings.Contains(result, "{") || strings.Contains(result, "}") {
				ok = false
				logger.Warn("route exception(redundant identifiers), may not work : ", patternTemp)
			}
			if ok {
				params += result + ","
			}
		} else {
			if !util.IsEmptyString(v) {
				pattern += "/" + v
			}
		}

	}
	// 移除最后的逗号
	if (len(params) > len(params)-1) && len(params) != 0 {
		params = string([]byte(params)[0 : len(params)-1])
	}
	// 遍历所有支持的方法，打印Debug日志
	for _, value := range method {
		logger.Debug(fmt.Sprintf("router <- %7s %s%s%s", value, logger.YellowBg, pattern, logger.Reset))
	}
	// 读写锁加锁
	// 之后添加路由项
	routerTable.RLock()
	defer routerTable.RUnlock()

	// 转map
	paramsMap := make(map[string][]string)
	for i, param := range strings.Split(params, ",") {
		paramsMap[param] = []string{strconv.Itoa(i), param}
	}
	routerTable.Table[pattern] = RouterTableItem{
		Method:   method,
		Pattern:  pattern,
		Function: function,
		Exclude:  "",
		Params:   paramsMap,
	}
}

// 根据处理函数去除所有该处理函数注册的路由项
func (routerTable RouterTable) RemoveFunc(function func(ctx context.Context)) {
	for key, value := range routerTable.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			routerTable.remove(key)
		}
	}
}

// 根据处理函数和请求方法排除所有该处理函数注册的路由项
func (routerTable RouterTable) ExcludeFuncByMethod(method string, function func(ctx context.Context)) {
	for _, value := range routerTable.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			if !strings.Contains(value.Exclude, method) {
				value.Exclude = util.MoreSpaceToOnce(strings.TrimSpace(value.Exclude + method + ";"))
			}
		}
	}
}

// 去除指定函数的方法排除
func (routerTable RouterTable) UNExcludeFuncByMethod(method string, function func(ctx context.Context)) {
	for _, value := range routerTable.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			if strings.Contains(value.Exclude, method) {
				value.Exclude = util.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(value.Exclude, method+";", "")))
			}
		}
	}
}

// 去除指定函数的所有方法排除
func (routerTable RouterTable) UNExcludeFunc(function func(ctx context.Context)) {
	for _, value := range routerTable.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			value.Exclude = ""
		}
	}
}

// 移除
func (routerTable RouterTable) remove(key string) {
	routerTable.RLock()
	delete(routerTable.Table, key)
	routerTable.RUnlock()
}
