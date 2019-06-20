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

const (
	// 运行类型
	MODE_DEV  = "dev"  // &describe 开发环境
	MODE_PROD = "prod" // &describe 生产环境
	// 日志级别
	LEVEL_FATAL string = "Fatal" // &describe 致命的错误
	LEVEL_ERROR string = "Error" // &describe 运行期间产生的错误
	LEVEL_WARN  string = "Warn"  // &describe 警告
	LEVEL_INFO  string = "Info"  // &describe 有意义的事件信息
	LEVEL_DEBUG string = "Debug" // &describe 调试信息
	LEVEL_TRACE string = "Trace" // &describe 更详细的跟踪信息
) // -< End
