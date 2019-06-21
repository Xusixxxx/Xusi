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

package config

import "time"

// 配置器实例
var Config Configurator

// 初始化配置器
func init() {
	Config = Configurator{
		Name:          "Xusi General Development Suite",
		PackageHeader: "Xusi/xusi-framework",
		Describe: `
			<h3>欢迎查看 <kbd>xusi-framework</kbd> 项目文档</h3>
			<footer>该文档由 <kbd>xusi-framework-xdoc</kbd> 自动化生成，为最近一次编译的最新版本文档。</footer>
		`,
		DescribeLabel: `
			<p style="font-size:13px;">@Xusixxxx ( xusixxxx@gmail.com )</p>
			<p style="font-size:13px;">最后更新时间：` + time.Now().String()[:19] + `</p>
		`,
	}
}

/* XusiStrcut ->
   @describe 控制xdoc一些状态的配置器
*/
type Configurator struct {
	Name          string // $describe 文档名称
	PackageHeader string // $describe 文档包头，会对import路径造成影响
	Describe      string // $describe 文档描述，可使用HTML编码
	DescribeLabel string // $describe 文档描述标签，可使用HTML编码
} // -< End
