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

// xweb 配置文件
// application的执行需要依靠config
package xweb

import (
	"xusi-projects/xusi-framework/xweb/httplib"
)

var conf config

func init() {
	conf = config{
		network: httplib.NETWORK_TCP4,
		address: httplib.DEFAULT_ADDRESS,
	}
}

type config struct {
	network string // 使用网络类型
	address string // 监听的地址
}
