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
    @describe 对xusi-framework-core-net的封装，快捷搭建net server。
<- End */
package xnet

import (
	"reflect"
	"sync"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/net/router"
	"xusi-projects/xusi-framework/core/net/server"
)

// xnet实例
var xnet *XNet

func init() {
	xnet = &XNet{}
}

type XNet struct {
	Server server.Blueprint // Network Server
	Router router.Blueprint // 路由器
}

var once sync.Once

// 初始化XNet
// 初始化时选择需要装载的路由器和HTTP Server
// 根据不同的装载内容产生不同的特性
/* XusiFunc ->
    @describe 初始化XNet并装载HTTP Server和Router，允许使用不同特性的符合Xusi蓝图设计概念的路由器和HTTP Server
    @param server server.Blueprint 符合Xusi设计概念的server实例
    @param router router.Blueprint 符合Xusi设计概念的router实例
<- End */
func Init(server server.Blueprint, router router.Blueprint) {
	once.Do(func() {
		logger.Debug("init xnet server and router...")
		xnet.Server = server
		xnet.Router = router
		logger.Debug("xnet server :", reflect.TypeOf(server).String())
		logger.Debug("xnet router :", reflect.TypeOf(router).String())
	})
}

// 获取XNet实例
/* XusiFunc ->
    @describe 获取到XNet实例，返回*XNet
<- End */
func Load() *XNet {
	return xnet
}

// 设置或取到运行模式
/* XusiFunc ->
    @describe 获取或设置XNet的运行模式，如果不传入参数，则标识为获取当前运行模式
    @param mode ...string 运行模式
<- End */
func RunMode(mode ...string) string {
	return server.RunMode(xnet.Server, mode...)
}

// 运行XNet
/* XusiFunc ->
    @describe 运行XNet实例
    @param params ...string 运行参数
<- End */
func (xnet *XNet) Run(params ...string) {
	server.Run(xnet.Server, params...)
}
