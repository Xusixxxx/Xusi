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
    @describe 一个符合Xusi设计概念的路由器(奥德丽)，完美主义的她希望任何事都做到极致，她会找到指定路由地址下所有允许的内容
<- End */
package audrey

import (
	basic2 "xusi-projects/xusi-framework/core/net/router/basic"
)

/* XusiStrcut ->
   @describe 路由器(奥德丽)结构
*/
type Audrey struct {
	*basic2.Router
}
