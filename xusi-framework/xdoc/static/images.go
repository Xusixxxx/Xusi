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

package static

import (
	"xusi-projects/xusi-framework/core/asset"
	"xusi-projects/xusi-framework/xweb"
	"xusi-projects/xusi-framework/xweb/context"
	"xusi-projects/xusi-framework/xweb/httplib"
)

func init() {
	xweb.Get("/logo", func(ctx *context.Context) {
		// 找到Logo
		isOK := false
		for _, value := range asset.AssetsMenu {
			if value.Name == "logo-xdoc.png" {
				logo, err := value.GetContext()
				if err != nil {
					ctx.StatusCode = httplib.CODE_500
				} else {
					ctx.Http.ResponseWriter.Header().Set("Content-Type", "image/gif")
					ctx.Http.ResponseWriter.Write(logo)
				}
				isOK = true
				break
			}
		}
		if !isOK {
			ctx.StatusCode = httplib.CODE_404
		}
	})
}
