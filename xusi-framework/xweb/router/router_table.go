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
	"strings"
	"sync"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util"
	"xusi-projects/xusi-framework/xweb/context"
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
}

// 创建一个路由表
func CreateRouterTable() RouterTable {
	return RouterTable{
		Table: map[string]RouterTableItem{},
	}
}

// 添加项(允许包含http协议外的method)
func (routerTable RouterTable) Add(pattern string, method []string, function func(ctx *context.Context)) {
	for _, value := range method {
		logger.Debug(fmt.Sprintf("router <- %7s %s%s%s", value, logger.YellowBg, pattern, logger.Reset))
	}
	routerTable.RLock()
	defer routerTable.RUnlock()
	routerTable.Table[pattern] = RouterTableItem{
		Method:   method,
		Pattern:  pattern,
		Function: function,
		Exclude:  "",
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
