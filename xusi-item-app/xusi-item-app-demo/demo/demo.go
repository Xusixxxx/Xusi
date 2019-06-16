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
    @describe 演示的包
<- End */
package demo

const name string = "测试=包含类型" // &describe 包含等于符号和类型的常量
const MODE_PROD = "prod"      // &describe 生产环境
const mODE_PROD2 = "prod"     // &describe 生产环境2
const MODE_DEV string = "dev" // &describe 开发环境

const (
	CONS1     = "con1" // &describe 测试con1
	CONS2 int = 1      // &describe 测试整型
) // -< End

/* XusiStrcut ->
   @describe 测试描述
*/
type Application struct {
	Host   string `json:"host"` // $describe 主机
	Port   int    // $describe 端口
	Mode   string `json:"mode"` // $describe 运行模式
	Debug  bool   // $describe 是否为Debug模式 测试的
	test   bool   // $describe 测试参数
	string        // $describe 测试参数2
	int           // $describe 测试参数3
} // -< End

/* XusiFunc ->
    @describe 开始执行应用程序
    @param mode 运行模式
    @param isDebug 是否为Debug模式
    @return Application 应用生命周期
<- End */
func Init(mode string, isDebug bool) Application {
	return Application{
		Host:  "localhost",
		Port:  80,
		Mode:  mode,
		Debug: isDebug,
	}
}
