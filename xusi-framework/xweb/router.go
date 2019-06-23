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
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util/xstring"
)

type router struct {
	sync.RWMutex
	Table map[string]routerTableItem // $describe 带读写锁的路由表
}

// 路由表项
/* XusiStrcut ->
   @describe 路由表子项，记录了一个路由完整的所包含的信息
*/
type routerTableItem struct {
	Method   []string            // $describe 接受的方法类型
	Pattern  string              // $describe 映射的地址
	Function func(ctx *context)  // $describe 路由处理的函数
	Exclude  string              // $describe 被排除的方法链串
	Params   map[string][]string // $describe 路由中携带的参数 map[name]{[]string{index, name}}
} // -< End

// 添加路由项(允许包含http协议外的method)
/* XusiFunc ->
    @describe 向路由表内添加子项
    @param patternTemp string 路由地址
    @param method []string 接受的方法类型集合
    @param function func(ctx *context.Context) 对应的路由处理函数
<- End */
func (routerTable router) Add(patternTemp string, method []string, function func(ctx *context)) {
	// copy
	pattern := patternTemp
	// 如果开头没有“/”则自动补上
	if string([]byte(pattern)[0]) != "/" {
		pattern = "/" + pattern
	}
	// 如果末尾为“/”则移除
	if string([]byte(pattern)[len(pattern)-1]) == "/" && pattern != "/" {
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
	// 如果路由为“/”则直接添加
	if pattern == "/" {
		goto Jump
	}
	pattern = ""
	for _, v := range slice {
		// 如果存在空段，则表示存在 “//”
		if xstring.IsEmptyString(v) {
			logger.Warn("route exception(exist null slice), may not work : ", patternTemp)
		}
		// 获取参数名
		if strings.Contains(v, "{") && strings.Contains(v, "}") {

			result := xstring.GetBetweenStr(v, "{", "}")
			ok := true
			// 如果为空，则表示存在不完整的标识符，可能仅有{，可能仅有}
			if xstring.IsEmptyString(result) {
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
			if !xstring.IsEmptyString(v) {
				pattern += "/" + v
			}
		}

	}
	// 移除最后的逗号
	if (len(params) > len(params)-1) && len(params) != 0 {
		params = string([]byte(params)[0 : len(params)-1])
	}
	// 路由为"/"
Jump:
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
	routerTable.Table[pattern] = routerTableItem{
		Method:   method,
		Pattern:  pattern,
		Function: function,
		Exclude:  "",
		Params:   paramsMap,
	}
}

// 根据处理函数去除所有该处理函数注册的路由项
/* XusiFunc ->
    @describe 根据处理函数去除所有该处理函数所注册的路由项
    @param function func(ctx context.Context) 路由处理函数
<- End */
func (router router) RemoveFunc(function func(ctx context)) {
	for key, value := range router.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			router.remove(key)
		}
	}
}

// 根据处理函数和请求方法排除所有该处理函数注册的路由项
/* XusiFunc ->
    @describe 根据处理函数的请求方法排除所有该处理函数注册的路由项
    @param method string 请求的类型
    @param function func(ctx context.Context) 路由处理函数
<- End */
func (router router) ExcludeFuncByMethod(method string, function func(ctx context)) {
	for _, value := range router.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			if !strings.Contains(value.Exclude, method) {
				value.Exclude = xstring.MoreSpaceToOnce(strings.TrimSpace(value.Exclude + method + ";"))
			}
		}
	}
}

// 去除指定函数的方法排除
/* XusiFunc ->
    @describe 去除指定函数的方法排除
    @param method string 请求的类型
    @param function func(ctx context.Context) 路由处理函数
<- End */
func (router router) UNExcludeFuncByMethod(method string, function func(ctx context)) {
	for _, value := range router.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			if strings.Contains(value.Exclude, method) {
				value.Exclude = xstring.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(value.Exclude, method+";", "")))
			}
		}
	}
}

// 去除指定函数的所有方法排除
/* XusiFunc ->
    @describe 去除指定函数的所有方法排除
    @param function func(ctx context.Context) 路由处理函数
<- End */
func (router router) UNExcludeFunc(function func(ctx context)) {
	for _, value := range router.Table {
		if reflect.ValueOf(function).Pointer() == reflect.ValueOf(value.Function).Pointer() {
			value.Exclude = ""
		}
	}
}

// 移除路由表项
func (router router) remove(key string) {
	router.RLock()
	delete(router.Table, key)
	router.RUnlock()
}
