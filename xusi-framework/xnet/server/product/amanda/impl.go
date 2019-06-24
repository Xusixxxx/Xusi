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
	"xusi-projects/xusi-framework/xnet/server/basic"
)

func (amanda *Amanda) Run(params []string) {
	amanda.runAdapter(params)
	amanda.listen()
}

func (amanda *Config) RunMode(mode []string) string {
	if len(mode) > 0 {
		if mode[0] == basic.RUN_MODE_DEV || mode[0] == basic.RUN_MODE_PROD {
			amanda.Config.RunMode = mode[0]
		} else {
			logger.Error(errors.New("error run mode type : " + mode[0] + "\n" + "only allow dev(basic.RUN_MODE_DEV) or prod(basic.RUN_MODE_PROD)"))
		}
	}
	return amanda.Config.RunMode
}
