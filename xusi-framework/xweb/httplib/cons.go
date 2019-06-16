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

// xweb 所有使用到的常量
package httplib

const (
	// 默认参数
	DEFAULT_ADDRESS = "8080"

	// 运行模式
	RUNMODE_PROD = "prod"
	RUNMODE_DEV  = "dev"

	// 网络类型
	NETWORK_TCP        = "tcp"
	NETWORK_TCP4       = "tcp4"
	NETWORK_TCP6       = "tcp6"
	NETWORK_UNIX       = "unix"
	NETWORK_UNIXPACKET = "unixpacket"

	// 请求类型
	METHOD_GET     = "GET"
	METHOD_POST    = "POST"
	METHOD_PUT     = "PUT"
	METHOD_DELETE  = "DELETE"
	METHOD_HEAD    = "HEAD"
	METHOD_TRACE   = "TRACE"
	METHOD_CONNECT = "CONNECT"
	METHOD_OPTIONS = "OPTIONS"

	// HTTP状态吗
	CODE_200 = 200 // 成功的请求

	CODE_301 = 301 // 已永久移动
	CODE_302 = 302 // 对象已移动
	CODE_304 = 304 // 未修改
	CODE_307 = 307 // 临时重定向

	CODE_404 = 404 // 找不到请求内容
	CODE_415 = 415 // 服务器拒绝服务，请求的格式不支持

	CODE_500 = 500 // 服务器内部异常

)
