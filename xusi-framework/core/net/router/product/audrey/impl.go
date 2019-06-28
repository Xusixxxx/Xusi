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

package audrey

import (
	"xusi-projects/xusi-framework/core/net/context"
	"xusi-projects/xusi-framework/core/util/xurl"
)

// 该路由器将直接返回所有找到的处理函数
func (audrey *Audrey) Find(ctx *context.Context) []func(*context.Context) {
	routeTableItem := audrey.Router.Table[xurl.UrlDecoder(ctx.URL.String())]
	if !routeTableItem.IsNil() {
		return routeTableItem.Functions
	} else {
		return nil
	}
}

// 如果找到的函数仅有一个，直接返回
// 如果找到的函数有多个，则返回其中一个
func (audrey *Audrey) FindTODO(ctx *context.Context) []func(*context.Context) {
	routeTableItem := audrey.Router.Table[xurl.UrlDecoder(ctx.URL.String())]
	if !routeTableItem.IsNil() {
		if len(routeTableItem.Functions) == 1 {
			return []func(*context.Context){routeTableItem.Functions[0]}
		} else {
			// TODO 权重算法

		}
	} else {
		return nil
	}
}