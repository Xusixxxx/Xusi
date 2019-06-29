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
    @describe HTTP Library（图书馆），存放net通用内容
<- End */
package httplibs

const (

	// 网络类型
	NETWORK_TCP        = "tcp"        // &describe 网络类型/tcp
	NETWORK_TCP4       = "tcp4"       // &describe 网络类型/tcp4
	NETWORK_TCP6       = "tcp6"       // &describe 网络类型/tcp6
	NETWORK_UNIX       = "unix"       // &describe 网络类型/unix
	NETWORK_UNIXPACKET = "unixpacket" // &describe 网络类型/unixpacket

	// 请求类型
	METHOD_GET     = "GET"     // &describe 请求类型/GET
	METHOD_POST    = "POST"    // &describe 请求类型/POST
	METHOD_PUT     = "PUT"     // &describe 请求类型/PUT
	METHOD_DELETE  = "DELETE"  // &describe 请求类型/DELETE
	METHOD_HEAD    = "HEAD"    // &describe 请求类型/HEAD
	METHOD_TRACE   = "TRACE"   // &describe 请求类型/TRACE
	METHOD_CONNECT = "CONNECT" // &describe 请求类型/CONNECT
	METHOD_OPTIONS = "OPTIONS" // &describe 请求类型/OPTIONS

	// HTTP状态吗
	CODE_200 = 200 // &describe 成功的请求

	CODE_301 = 301 // &describe 已永久移动
	CODE_302 = 302 // &describe 对象已移动
	CODE_304 = 304 // &describe 未修改
	CODE_307 = 307 // &describe 临时重定向

	CODE_404 = 404 // &describe 找不到请求内容
	CODE_415 = 415 // &describe 服务器拒绝服务，请求的格式不支持

	CODE_500 = 500 // &describe 服务器内部异常

) // -< End
