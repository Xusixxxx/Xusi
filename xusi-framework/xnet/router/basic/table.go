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

package basic

// 路由器表成员结构
type RouteTableItem struct {
	Patterns  Patterns  // 映射的地址
	Methods   Methods   // 接受的方法类型
	Functions Functions // 路由处理的函数
}

/* XusiFunc ->
    @describe 检查该路由成员是否为空
<- End */
func (routeTableItem RouteTableItem) IsNil() bool {
	return len(routeTableItem.Methods) == 0 && len(routeTableItem.Functions) == 0 && len(routeTableItem.Patterns) == 0
}