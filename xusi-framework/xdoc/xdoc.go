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
	"xusi-projects/xusi-framework/core/asset"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/xdoc/model"
	"xusi-projects/xusi-framework/xdoc/static"
	"xusi-projects/xusi-framework/xdoc/xdoc_util"
	"xusi-projects/xusi-framework/xnet"
)

// 文档字典
var Docs = map[string]model.PackageModel{}

/* XusiFunc ->
    @describe 运行xdoc文档web服务器
    @param port string 监听端口
<- End */
func Run(port string) {
	// 日志配置
	logger.Conf.Mode = logger.MODE_PROD
	// 遍历所有静态资产
	for _, assetFile := range asset.AssetsMenu {
		// 格式化静态资产文件内容
		content, err := xdoc_util.FormatAssets(assetFile)
		if err != nil {
			panic(err)
		}

		// 开始解析
		logger.Info("analysis >> " + assetFile.Name)
		startAnalysis(content, assetFile)
	}
	// 全部文件解析完成后排除多余路径

	// 路由解析
	router()
	// 启动web服务
	xnet.Run(port)
}

// 生成文档路由
func router() {
	// 根路由
	root := "/"

	// 加载目录
	xnet.Get(root, func(ctx *xnet.Context) {
		page := RenderContent(static.PAGE_DOC)
		page = RenderSidebar(page)

		ctx.WirteString(page)
	})
}
