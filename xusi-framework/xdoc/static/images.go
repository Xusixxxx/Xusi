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
	"xusi-projects/xusi-framework/core/net/context"
	"xusi-projects/xusi-framework/core/net/httplibs"
	"xusi-projects/xusi-framework/xnet"
)

// 初始化图片资源
func InitImage() {
	asset.Add("C:/Users/Xusixxxx/Documents/Sources/golang/src/xusi-projects/xusi-framework/xdoc/static/logo-xdoc.png", asset.Assets{
		Name:     "logo-xdoc.png",
		Content:  "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",
		FileName: "logo-xdoc.png",
		FileType: "png",
		FullName: "C:/Users/Xusixxxx/Documents/Sources/golang/src/xusi-projects/xusi-framework/xdoc/static/logo-xdoc.png",
	})

	asset.Add("C:/Users/Xusixxxx/Documents/Sources/golang/src/xusi-projects/xusi-framework/xdoc/static/favicon.ico", asset.Assets{
		Name:     "favicon.ico",
		Content:  "H4sIAAAAAAAA/6RTSUwTYRT+SYgXL8SD8chNj549mhA8kXhSg0tMjRhLCGgw1ahocKtsBjSgQjRRIAaiwYCiImspIN1o6UhpLRDoRikOS+0y889nZiu2AWPi1zRv5s173/u+f94QkkWySE6OGHPJhWxC9hJCDhBCcgghuUTOS8gmZM9u+a8CO0CAIEXPRgDh+JqS+zeodQvJOI57RtDvG0zj/NtM8ccJVLrXh9zI/1aL5/a6DOb0nhifABW2nnGUk2KV6Ta05hrUeOX5vMCDZnDEBAFN80O4OHgeT5zN4AFMBMfR5e1E40wbjtlb0ebpkGqpokvVqeK9fxyXRrQ4Z23GCesrGEMmHO3Oww3mLSptjai06FE19QjPnE/h2AwimeFF1Ptm5gWKmQ6UTL9G9eRNaPpPo9hQhnp7PUrnjCieasG1MR0OObtRM/cVXHId9I8zEGGJzOKO+SH0TBva5z5CN1qGId8Aro+V4+RELW65e9Bgb4CmrxB6W11qvurr3cInlBtKcNfZiq4Qg5WoX8qP+odRNFSE/d+HscvahVMzvTg4WIn77t7Uebriv1Dg6sO9iavItzSjlV0WmbG46cPPBAtKE/gRY1E4b0LpkgPjqx60+ybBK++tM+TCGXMTXtoe4PDkY1SHFwE+ioWNJUxHphGMBrfdIVV7i/czrhh10A5ocLZfI1XFuCgcK3bZQ8CAgMIhzhT3hRcoqMBLuWpTBbTWBuT1FEBnvCzl4nwclmWzdJ2kSbjZWbAJNn2+oudLmME+Rx8qbPXwsq5UD7PKyLoVn/JubnlQ/X9YC+GI25DmUayNxCLb7lwmopTDOi9zq9+CIPmkO/aIIP+J3wEAAP//7laG8X4EAAA=",
		FileName: "favicon.ico",
		FileType: "ico",
		FullName: "C:/Users/Xusixxxx/Documents/Sources/golang/src/xusi-projects/xusi-framework/xdoc/static/favicon.ico",
	})

	xnet.Load().Get("/logo", func(ctx *context.Context) {
		// 找到Logo
		isOK := false
		for _, value := range asset.AssetsMenu {
			if value.Name == "logo-xdoc.png" {
				logo, err := value.GetContext()
				if err != nil {
					ctx.StatusCode = httplibs.CODE_500
				} else {
					ctx.ResponseWriter.Header().Set("Content-Type", "image/gif")
					ctx.ResponseWriter.Write(logo)
				}
				isOK = true
				break
			}
		}
		if !isOK {
			ctx.StatusCode = httplibs.CODE_404
		}
	})

	xnet.Load().Get("/favicon.ico", func(ctx *context.Context) {
		// 找到favicon
		isOK := false
		for _, value := range asset.AssetsMenu {
			if value.Name == "favicon.ico" {
				logo, err := value.GetContext()
				if err != nil {
					ctx.StatusCode = httplibs.CODE_500
				} else {
					ctx.ResponseWriter.Header().Set("Content-Type", "image/x-icon")
					ctx.ResponseWriter.Write(logo)
				}
				isOK = true
				break
			}
		}
		if !isOK {
			ctx.StatusCode = httplibs.CODE_404
		}
	})
}
