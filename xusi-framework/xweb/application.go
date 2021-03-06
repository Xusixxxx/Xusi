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

// xweb应用程序
// xweb的应用程序
// 对xweb进行统一的管理
package xweb

import (
	"net"
	"net/http"
	"xusi-projects/xusi-framework/core/logger"
)

// xweb应用程序示例
var app *application

// 初始化xweb应用程序
func init() {
	app = &application{
		Server:   &http.Server{},
		Handlers: &requestHandler{},
	}
}

// xweb 应用程序
type application struct {
	Server   *http.Server    // 网络请求接收中心
	Handlers *requestHandler // 网络请求处理者
}

// 监听网络
// network：监听的网络类型 tcp ...
// address：监听的网络端口
func listen(params ...string) {
	endRunning := make(chan bool, 1)

	go func() {

		if len(params) != 0 {
			conf.network = params[0]
			conf.address = params[1]
		}

		//启动一个 http server
		listener, err := net.Listen(conf.network, ":"+conf.address)
		logger.Info("xusi http server running...")
		logger.Info("address in http server is ->", logger.Cyan, conf.address+"("+conf.network+")", logger.Reset)
		if err != nil {
			endRunning <- true
			return
		}
		err = app.Server.Serve(listener)
		if err != nil {
			endRunning <- true
			return
		}
	}()

	// 阻塞宿主程序
	<-endRunning
}
