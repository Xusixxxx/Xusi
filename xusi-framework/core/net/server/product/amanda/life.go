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
	"net/http"
	"xusi-projects/xusi-framework/core/net/httplibs"
	basic2 "xusi-projects/xusi-framework/core/net/server/basic"
)

var server *Amanda

func init() {
	server = &Amanda{
		Server: &basic2.Server{},
		Config: &Config{
			Config: &basic2.Config{
				Network: httplibs.NETWORK_TCP4,
				Address: DEFAULT_ADDRESS,
				Port:    DEFAULT_PORT,
				RunMode: basic2.DEFAULT_RUN_MODE,
			},
		},
		HttpServer: &http.Server{},
		Handlers:   &requestHandler{},
	}
}

func Load() *Amanda {
	return server
}