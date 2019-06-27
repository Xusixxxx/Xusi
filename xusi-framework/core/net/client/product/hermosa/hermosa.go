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

package hermosa

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"xusi-projects/xusi-framework/core/net/client/basic"
	"xusi-projects/xusi-framework/core/net/client/util"
	"xusi-projects/xusi-framework/core/net/httplibs"
)

type Hermosa struct {
	*basic.Client
	HttpClient *http.Client // HTTP客户端
}

// 发起HTTP请求
func (hermosa *Hermosa) HTTP(method string, xurl string, header map[string]string, data map[string]string) ([]byte, error) {
	// method转大写
	method = strings.ToUpper(method)
	// 变量初始化
	var response *http.Response
	var reader *strings.Reader
	// 配置请求信息
	switch method {
	case httplibs.METHOD_GET:
		response, err := http.Get(xurl + "?" + util.HttpDataMapAnalysis(data))
		if err != nil {
			return nil, err
		}
	default:

	}

	if err != nil {
		// handle error
	}

	response.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response.Header.Set("Cookie", "name=anny")

	resp, err := hermosa.HttpClient.Do(response)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
