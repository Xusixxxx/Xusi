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

package amanda

import (
	"net"
	"net/http"
	"strconv"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xnet/server/basic"
)

type Amanda struct {
	*basic.Server
	*Config
	HttpServer *http.Server    // 网络请求接收中心
	Handlers   *requestHandler // 网络请求处理者
}

// 启动监听
func (amanda Amanda) listen() {
	endRunning := make(chan bool, 1)
	go func() {
		// 注册路由器函数
		http.Handle("/", amanda.Handlers)
		// 启动一个 http server
		address := amanda.Config.Address + ":" + strconv.Itoa(amanda.Config.Port)
		listener, err := net.Listen(amanda.Config.Network, address)
		logger.Info("xusi http server running...")
		logger.Info("address in http server is ->", logger.Cyan, address+"("+amanda.Config.Network+")", logger.Reset)
		if err != nil {
			endRunning <- true
			return
		}
		err = amanda.HttpServer.Serve(listener)
		if err != nil {
			endRunning <- true
			return
		}
	}()
	// 阻塞宿主程序
	<-endRunning
}

// 启动适配
// 8080
// localhost, 8080
// tcp4, localhost, 8080
func (amanda *Amanda) runAdapter(params []string) {
	// 检查入参
	// 启动http server
	switch len(params) {
	case 1:
		port, err := strconv.Atoi(params[0])
		if err != nil {
			panic(err)
		}
		amanda.Config.Port = port
	case 2:
		port, err := strconv.Atoi(params[1])
		if err != nil {
			panic(err)
		}
		amanda.Config.Address = params[0]
		amanda.Config.Port = port
	case 3:
		port, err := strconv.Atoi(params[2])
		if err != nil {
			panic(err)
		}
		amanda.Config.Network = params[0]
		amanda.Config.Address = params[1]
		amanda.Config.Port = port
	}
}
