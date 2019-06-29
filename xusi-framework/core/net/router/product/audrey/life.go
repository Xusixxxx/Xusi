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
	basic2 "xusi-projects/xusi-framework/core/net/router/basic"
)

// 路由器实例
var router *Audrey

func init() {
	router = &Audrey{
		&basic2.Router{
			Table: map[string]basic2.RouteTableItem{},
		},
	}
}

/* XusiFunc ->
    @describe 获取到Audrey路由器的实例
<- End */
func Load() *Audrey {
	return router
}
