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

package xdoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"xusi-projects/xusi-framework/core/util"
	"xusi-projects/xusi-framework/xweb/context"
)

// 渲染内容区的目录
func RenderMenu(rootPage string, ctx *context.Context) {
	menu := ""
	// 排序
	var keys []string
	for key := range Docs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		v := Docs[key]
		if v.IsNull() {
			continue
		}
		if util.IsEmptyString(v.Describe) {
			v.Describe = "none"
		}
		ul := `
				<ul style="list-style:none;">
					<a href="#` + v.Name + `">
						<blockquote>
							<p><span class="keyword">package</span> : <span class="package-name">` + v.Name + `</span></p>
							<footer>` + v.Describe + `</footer></br>
							{struct}
							{function}
						</blockquote>
					</a>

				</ul>
			`

		// 添加结构体
		// 排序
		var structKeys []string
		for structKey := range v.Struct {
			structKeys = append(structKeys, structKey)
		}
		sort.Strings(structKeys)

		structStrTemp := ""
		for _, structKey := range structKeys {
			struct4 := v.Struct[structKey]
			structStrTemp += `
					<blockquote style="font-size:13px;">
						<li role="presentation">
							<a href="#"><span class="keyword">type</span> <span class="type-name">` + struct4.Name + `</span></a>
							<span style="color:#999999;font-size:12px;"> - ` + struct4.Describe + `</p>
						</li>
					</blockquote>
				`
		}

		// 添加函数
		// 排序
		var funcKeys []string
		for funKey := range v.Func {
			funcKeys = append(funcKeys, funKey)
		}
		sort.Strings(funcKeys)

		functionStrTemp := ""
		for _, funcKey := range funcKeys {
			function := v.Func[funcKey]
			functionStrTemp += `
					<blockquote style="font-size:13px;">
						<li role="presentation">
							<a href="#" ` + function.Name + "-" + function.Name + `"><span class="keyword">func</span> <span class="func-name">` + function.Name + `</span></a>
								({func-parames})
							<span style="color:#999999;font-size:12px;"> - ` + function.Describe + `</p>

						</li>
					</blockquote>
				`
			// 添加函数参数
			// 排序
			var paramKeys []string
			for paramKey := range function.Params {
				paramKeys = append(paramKeys, strconv.Itoa(function.Params[paramKey].Index)+"|"+paramKey)
			}
			sort.Strings(paramKeys)

			paramStrTemp := ""
			paramIndex := 1
			for _, paramKey := range paramKeys {
				param := function.Params[strings.Split(paramKey, "|")[1]]
				fmt.Println(param.Name, param.Type)
				if paramIndex < len(function.Params) {
					paramStrTemp += `
							<span class="badge param-tip" data-placement="top" title="` + param.Describe + `">
								` + param.Name + " " + param.Type + `
							</span>,
						`
				} else {
					paramStrTemp += `
							<span class="badge param-tip" data-placement="top" title="` + param.Describe + `">
								` + param.Name + " " + param.Type + `
							</span>
						`
				}
				paramIndex++
			}
			// 去除最后一个逗号
			functionStrTemp = strings.ReplaceAll(functionStrTemp, "{func-parames}", paramStrTemp)
		}
		ul = strings.ReplaceAll(ul, "{function}", functionStrTemp)
		ul = strings.ReplaceAll(ul, "{struct}", structStrTemp)
		menu += ul
	}
	ctx.WirteString(strings.ReplaceAll(rootPage, "{menu-content}", menu))
}
