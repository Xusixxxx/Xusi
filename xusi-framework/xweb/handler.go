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

// xweb请求处理者
package xweb

import (
	"net/http"
	"strconv"
	"strings"
	"xusi-projects/xusi-framework/core/logger"
	"xusi-projects/xusi-framework/core/util"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
)

// 请求处理者
type requestHandler struct {
	*context.Context
	realRoute string // 真实路由，解析后访问的最终路由
}

// 统一函数处理
// 所有函数都将经过这个
func (r *requestHandler) serveHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// 如果请求的方法不被允许或不存在
	// 解析url，避免被拒
	key := util.UrlDecoder(request.URL.String())
	// 如果末尾有斜杠，则移除
	if string([]byte(key)[len(key)-1]) == "/" && key != "/" {
		key = string([]byte(key)[0 : len(key)-1])
	}
	// 是否为携带参数
	if strings.Contains(key, "?") {
		key = strings.Split(key, "?")[0]
	}

	// 如果未携带任何参数，但是绑定了带参数的路由
	// 那么应该请求到空参数的路由函数
	// 否则则视为有参数，进入解析
	if _, ok := xrouterInstance.routerTable.Table[key]; ok {
		r.realRoute = key
	} else {
		// 解析路由参数
		// 取到每一段切片，并排除第一段空切片
		slice := strings.Split(key, "/")
		slice = slice[1:len(slice)]
		// 遍历所有路由进行比对
		for k, v := range xrouterInstance.routerTable.Table {
			// 对遍历的pattern切片
			// 属性拼接再切片
			kTemp := k
			for _, v := range v.Params {
				kTemp += "/{" + v[1] + "}"
			}
			rSlice := strings.Split(kTemp, "/")
			rSlice = rSlice[1:len(rSlice)]
			// 如果传入的请求url切片量大于了该路由切片量不等，则不对应，进入下一个
			if len(slice) != len(rSlice) {
				continue
			} else {
				// 如果第一段切片不通，则不对应，进入下一个
				if slice[0] != rSlice[0] {
					continue
				} else {
					// 切片1对比切片1，切片2对比切片2
					// 属性段的索引需要由i减去非属性段数量
					paramI := 0
					for i, rs := range rSlice {
						// 检测当前是否为属性段
						isParamSlice := false
						for _, param := range v.Params {
							if param[0] == strconv.Itoa(i-paramI) && (i < len(slice)) && (rs != slice[i]) {
								isParamSlice = true
								break
							}
						}
						// 如果为属性段，这段则通过，开始匹配下一段，并且去掉属性段
						// 如果不为属性段，则匹配是否相同
						if isParamSlice {
							paramI++
							key = strings.ReplaceAll(key, "/"+slice[i], "")
						} else {
							if slice[i] != rs {
								break
							}
						}
					}
				}
			}
		}
		r.realRoute = key
	}
	logger.Debug("route params parse url : " + key)
	logger.Debug("parse url >> " + key)
	if _, ok := xrouterInstance.routerTable.Table[key]; ok {
		// 是否通过
		isPass := false
		logger.Debug("check route : " + util.UrlDecoder(key))
		// 遍历所有该路由所支持的方法
		for _, value := range xrouterInstance.routerTable.Table[key].Method {
			logger.Debug("router take in method : ", value, ", input method : ", request.Method)
			// 如果找到对应被允许的方法，那么则通过验证
			if value == request.Method {
				isPass = true
				logger.Debug("verification by")
				break
			}
		}
		if isPass == false {
			logger.Debug("fail verification")
			r.StatusCode = httplib.CODE_415
		}
	} else {
		logger.Debug("not found route : " + util.UrlDecoder(request.URL.String()))
		r.StatusCode = httplib.CODE_404
	}

	// 解析From表单
	request.ParseForm()

}
