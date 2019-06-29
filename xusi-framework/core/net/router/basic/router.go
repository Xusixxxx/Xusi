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

import (
	"sync"
)

/* XusiStrcut ->
   @describe 符合Xusi设计概念的路由器结构模型
*/
type Router struct {
	sync.RWMutex                           // $describe 读写锁(通常用于守护路由器内部读写安全)
	Table        map[string]RouteTableItem // $describe 通用基础路由表
} // -< End
