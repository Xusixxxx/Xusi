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
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util/xmap"
	"xusi-projects/xusi-framework/core/util/xstring"
	"xusi-projects/xusi-framework/xdoc/config"
	"xusi-projects/xusi-framework/xdoc/model"
	"xusi-projects/xusi-framework/xdoc/static"
)

// 渲染总览菜单目录
func RenderSidebar(rootPage string) string {
	// 初始化内容
	content := ""
	// 对map进行排序
	keys, err := xmap.SortMapByStringKey(Docs)
	if err != nil {
		logger.Error(err)
	}
	// 添加菜单html内容
	for _, key := range keys {
		pkg := Docs[key]
		if !pkg.IsNull() {
			content += `<ul><li><a href="#` + pkg.GetPackagePath() + `">` + pkg.GetPackagePath() + `</a></li></ul>`
		}
	}
	return strings.ReplaceAll(rootPage, static.PAGE_SLDEBAR, content)
}

// 渲染内容区的目录
func RenderContent(rootPage string) string {
	// 初始化内容
	content := ""
	// 对package模型的存储map进行排序，确保不会因为刷新导致乱序
	keys, err := xmap.SortMapByStringKey(Docs)
	if err != nil {
		logger.Error(err)
	}
	// 遍历所有package模型，开始生产文档内容
	for _, key := range keys {
		packageModel := Docs[key]
		// 如果这个package不包含任何需要公开在文档内的内容，那么则不显示，直接跳过
		if packageModel.IsNull() {
			continue
		}
		// 如果package描述未填写则默认为none
		if xstring.IsEmptyString(packageModel.Describe) {
			packageModel.Describe = "none"
		}
		// 生产package的Html内容
		packageHtml := `
			<h1 style="line-height:80px" id="` + packageModel.GetPackagePath() + `">&nbsp;</h1>
			<ul style="list-style:none;">
				<blockquote>
					<a href="#` + packageModel.GetPackagePath() + `"><p><span class="keyword">package</span> : <span class="package-name">` + packageModel.Name + `</span></p></a>
					<p style="margin-top:5px;font-size:12px;color:#FF6600;">import "` + config.Config.PackageHeader + "/" + string([]byte(packageModel.GetPackagePath())[1:len(packageModel.GetPackagePath())]) + `"</p>
					<footer>` + packageModel.Describe + `</footer></br>
					{const}
					{struct}
					{function}
				</blockquote>
			</ul>
		`
		// 生产各项子成员Html
		packageHtml = renderStruct(packageHtml, packageModel)
		packageHtml = renderConst(packageHtml, packageModel)
		packageHtml = renderFunc(packageHtml, packageModel)

		content += packageHtml
	}
	page := strings.ReplaceAll(rootPage, static.PAGE_MENU_CONTENT, content)
	// 配置项调整
	page = strings.ReplaceAll(page, static.PAGE_NAME, config.Config.Name)
	page = strings.ReplaceAll(page, static.PAGE_DESCRIBE, config.Config.Describe)
	page = strings.ReplaceAll(page, static.PAGE_DESCRIBE_LABEL, config.Config.DescribeLabel)

	return page
}

// 渲染结构体
func renderStruct(content string, packageModel model.PackageModel) string {
	// 对结构体集合进行排序
	structKeys, err := xmap.SortMapByStringKey(packageModel.Struct)
	if err != nil {
		logger.Error(err)
	}

	// 初始化结构体模板
	// 遍历所有结构体集合成员，生成Html内容
	structStrTemp := ""
	for _, structKey := range structKeys {
		struct4 := packageModel.Struct[structKey]
		structStrTemp += `
			<blockquote style="font-size:13px;">
				<a data-toggle="collapse" href="#` + strings.ReplaceAll(packageModel.GetPackagePath(), "/", "") + struct4.Name + `" aria-expanded="false" aria-controls="` + packageModel.GetPackagePath() + struct4.Name + `">
					<li role="presentation">
						<span class="keyword">type</span> <span class="type-name">` + struct4.Name + `</span>
						<span style="color:#999999;font-size:12px;"> - ` + struct4.Describe + `</span>
					</li>
				</a>
				<div style="margin-top:10px" class="collapse" id="` + strings.ReplaceAll(packageModel.GetPackagePath(), "/", "") + struct4.Name + `">
					<div class="well">
						<table style="table-layout:fixed;" border="0">
							{struct-attr}
						</table>
					</div>
				</div>
			</blockquote>
		`
		// 生成结构体内的属性Html信息
		structStrTemp = renderStructAttr(structStrTemp, struct4)
	}
	return strings.ReplaceAll(content, "{struct}", structStrTemp)
}

// 渲染结构体属性
func renderStructAttr(content string, structModel model.StructModel) string {
	// 对结构体集合进行排序
	structAttrKeys, err := xmap.SortMapByStringKey(structModel.Attr)
	if err != nil {
		logger.Error(err)
	}

	// 初始化结构体属性模板
	// 遍历所有结构体属性成员，生成Html内容
	structAttrStrTemp := ""
	for _, structAttrKey := range structAttrKeys {
		attr := structModel.Attr[structAttrKey]
		structAttrStrTemp += `
			<tr>
				<td style="padding:5px;"><span>` + attr.Name + `</span></td>
				<td style="padding:5px;"><kbd>` + attr.Type + `</kbd></td>
				<td style="padding:5px;color:#3333CC;"> ` + attr.Addition + ` </td>
				<td style="padding:5px;"><span style="color:#999999;font-size:12px;"> - ` + attr.Describe + `</span></td>
			</tr>
		`
	}
	return strings.ReplaceAll(content, "{struct-attr}", structAttrStrTemp)
}

// 渲染结构体常量
func renderConst(content string, packageModel model.PackageModel) string {
	// 对结构体常量集合进行排序
	constKeys, err := xmap.SortMapByStringKey(packageModel.Const)
	if err != nil {
		logger.Error(err)
	}

	// 初始化常量模板
	// 遍历所有结构体常量成员生成Html内容
	constStrTemp := ""
	for _, constKey := range constKeys {
		const4 := packageModel.Const[constKey]
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
	return strings.ReplaceAll(content, "{const}", constStrTemp)
}

// 渲染结构体函数
func renderFunc(content string, packageModel model.PackageModel) string {
	// 对结构体的函数集合进行排序
	funcKeys, err := xmap.SortMapByStringKey(packageModel.Func)
	if err != nil {
		logger.Error(err)
	}
	// 初始化结构体函数模板
	// 遍历所有结构体内的函数生成Html内容
	functionStrTemp := ""
	for _, funcKey := range funcKeys {
		funcModel := packageModel.Func[funcKey]
		functionStrTemp += `
			<blockquote style="font-size:13px;">
				<li role="presentation">
					<span class="keyword">func</span> <span class="func-name">` + funcModel.Name + `</span>
						({func-parames})
					<span style="color:#999999;font-size:12px;"> - ` + funcModel.Describe + `</p>
				</li>
			</blockquote>
		`
		functionStrTemp = renderFuncParam(functionStrTemp, funcModel)
	}
	return strings.ReplaceAll(content, "{function}", functionStrTemp)
}

// 渲染结构体函数内的参数
func renderFuncParam(content string, funcModel model.FuncModel) string {
	// 对结构体函数内的参数进行排序
	var paramKeys []string
	for paramKey := range funcModel.Params {
		// 为了保证参数的顺序
		// 对key进行拼接，以参数位置作为开头
		paramKeys = append(paramKeys, strconv.Itoa(funcModel.Params[paramKey].Index)+"|"+paramKey)
	}
	sort.Strings(paramKeys)
	// 初始化参数模板
	// 遍历所有参数生成Html内容
	paramStrTemp := ""
	paramIndex := 1
	for _, paramKey := range paramKeys {
		param := funcModel.Params[strings.Split(paramKey, "|")[1]]
		// 检测是否为最后一个参数
		// 如果不是，那么结尾不需要逗号
		if paramIndex < len(funcModel.Params) {
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
	return strings.ReplaceAll(content, "{func-parames}", paramStrTemp)
}
