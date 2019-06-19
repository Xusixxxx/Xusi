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

package logger

// 日志配置器
var Conf loggerConfigurator

// 初始化日志配置器，赋予一些默认的设置
func init() {
	Conf = loggerConfigurator{
		Mode:    MODE_DEV,
		Disable: false,
	}
}

// 日志配置器结构体
type loggerConfigurator struct {
	Mode    string // 运行环境
	Disable bool   // 是否禁用日志
}
