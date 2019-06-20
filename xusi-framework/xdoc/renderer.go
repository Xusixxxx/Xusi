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
	"sort"
	"strconv"
	"strings"
	"xusi-projects/xusi-framework/core/util"
)

// 渲染总览目录
func RenderSidebar(rootPage string) string {
	menu := ""
	// 排序
	var keys []string
	for key := range Docs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		pkg := Docs[key]
		if !pkg.IsNull() {
			menu += `<ul><li><a href="#` + pkg.GetPackagePath() + `">` + pkg.GetPackagePath() + `</a></li></ul>`
		}
	}
	return strings.ReplaceAll(rootPage, PAGE_SLDEBAR, menu)
}

// 渲染内容区的目录
func RenderMenu(rootPage string) string {
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
				<h1 style="line-height:80px" id="` + v.GetPackagePath() + `">&nbsp;</h1>
				<ul style="list-style:none;">
						<blockquote>
							<a href="#` + v.GetPackagePath() + `"><p><span class="keyword">package</span> : <span class="package-name">` + v.Name + `</span></p></a>
							<p style="margin-top:5px;font-size:12px;color:#FF6600;">import ` + string([]byte(v.GetPackagePath())[1:len(v.GetPackagePath())]) + `</p>
							<footer>` + v.Describe + `</footer></br>
							{const}
							{struct}
							{function}
						</blockquote>
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
						<a data-toggle="collapse" href="#` + strings.ReplaceAll(v.GetPackagePath(), "/", "") + struct4.Name + `" aria-expanded="false" aria-controls="` + v.GetPackagePath() + struct4.Name + `">
							<li role="presentation">
								<span class="keyword">type</span> <span class="type-name">` + struct4.Name + `</span>
								<span style="color:#999999;font-size:12px;"> - ` + struct4.Describe + `</span>
							</li>
						</a>
						<div style="margin-top:10px" class="collapse" id="` + strings.ReplaceAll(v.GetPackagePath(), "/", "") + struct4.Name + `">
							<div class="well">
								<table style="table-layout:fixed;" border="0">
									{struct-attr}
								</table>
							</div>
						</div>
					</blockquote>
				`

			// 添加结构体信息
			var structAttrKeys []string
			for structAttrKey := range struct4.Attr {
				structAttrKeys = append(structAttrKeys, structAttrKey)
			}
			sort.Strings(structKeys)

			structAttrStrTemp := ""
			for _, structAttrKey := range structAttrKeys {
				attr := struct4.Attr[structAttrKey]
				structAttrStrTemp += `
					<tr>
						<td style="padding:5px;"><span>` + attr.Name + `</span></td>
						<td style="padding:5px;"><kbd>` + attr.Type + `</kbd></td>
						<td style="padding:5px;color:#3333CC;"> ` + attr.Addition + ` </td>
						<td style="padding:5px;"><span style="color:#999999;font-size:12px;"> - ` + attr.Describe + `</span></td>
					</tr>
				`
			}
			structStrTemp = strings.ReplaceAll(structStrTemp, "{struct-attr}", structAttrStrTemp)
		}

		// 添加常量
		// 排序
		var constKeys []string
		for constKey := range v.Const {
			constKeys = append(constKeys, constKey)
		}
		sort.Strings(constKeys)

		constStrTemp := ""
		for _, constKey := range constKeys {
			const4 := v.Const[constKey]
			if const4.Type != "unknown" {
				const4.Type = `<kbd>` + const4.Type + `</kbd>`
			} else {
				const4.Type = ""
			}
			constStrTemp += `
					<blockquote style="font-size:13px;">
						<li role="presentation">
							<span style="color:#CC0000;margin-right:10px;">const</span><span class="type-name" style="color:#FF6600;">` + const4.Name + `</span>
							 ` + const4.Type + ` = ` + const4.Value + `
							<span style="color:#999999;font-size:12px;"> - ` + const4.Describe + `</p>
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
							<span class="keyword">func</span> <span class="func-name">` + function.Name + `</span>
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
				if paramIndex < len(function.Params) {
					paramStrTemp += `
							<span class="badge param-tip" data-placement="top" title="` + param.Describe + `">
								` + param.Name + " <kbd>" + param.Type + `</kbd>
							</span>,
						`
				} else {
					paramStrTemp += `
							<span class="badge param-tip" data-placement="top" title="` + param.Describe + `">
								` + param.Name + " <kbd>" + param.Type + `</kbd>
							</span>
						`
				}
				paramIndex++
			}
			// 去除最后一个逗号
			functionStrTemp = strings.ReplaceAll(functionStrTemp, "{func-parames}", paramStrTemp)
		}
		ul = strings.ReplaceAll(ul, "{function}", functionStrTemp)
		ul = strings.ReplaceAll(ul, "{const}", constStrTemp)
		ul = strings.ReplaceAll(ul, "{struct}", structStrTemp)
		menu += ul
	}
	return strings.ReplaceAll(rootPage, PAGE_MENU_CONTENT, menu)
}
