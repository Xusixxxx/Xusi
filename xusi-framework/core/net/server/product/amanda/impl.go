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

package amanda

import (
	"errors"
	"xusi-projects/xusi-framework/core/logger"
	basic2 "xusi-projects/xusi-framework/core/net/server/basic"
)

/* XusiFunc ->
    @describe 运行Amanda
    @param params []string 运行参数
<- End */
func (amanda *Amanda) Run(params []string) {
	amanda.runAdapter(params)
	amanda.listen()
}

/* XusiFunc ->
    @describe 获取或设置Amanda的运行模式，如果不传参则为获取
    @param mode []string 运行模式
<- End */
func (amanda *Config) RunMode(mode []string) string {
	if len(mode) > 0 {
		if mode[0] == basic2.RUN_MODE_DEV || mode[0] == basic2.RUN_MODE_PROD {
			amanda.Config.RunMode = mode[0]
		} else {
			logger.Error(errors.New("error run mode type : " + mode[0] + "\n" + "only allow dev(basic.RUN_MODE_DEV) or prod(basic.RUN_MODE_PROD)"))
		}
	}
	return amanda.Config.RunMode
}
