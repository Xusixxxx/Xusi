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

package api

import (
	"xusi-projects/xusi-framework/xweb"
	"xusi-projects/xusi-framework/xweb/context"
)

func init() {

	xweb.Get("/hello/{id}/{name}", func(ctx *context.Context) {
		ctx.WirteString("hello, [" + ctx.RouterParams["id"] + "]" + ctx.RouterParams["name"] + "!")
	})

	xweb.Get("/hello1", func(ctx *context.Context) {
		ctx.WirteString("ces1")
	})

	xweb.Get("/hello2", func(ctx *context.Context) {
		ctx.WirteString("ces2")
	})
}
