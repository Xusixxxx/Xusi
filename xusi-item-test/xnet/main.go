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

package main

import (
	"xusi-projects/xusi-framework/core/net/context"
	"xusi-projects/xusi-framework/core/net/router"
	"xusi-projects/xusi-framework/core/net/router/product/audrey"
	"xusi-projects/xusi-framework/core/net/server"
	"xusi-projects/xusi-framework/core/net/server/product/amanda"
	"xusi-projects/xusi-framework/xnet"
)

func init() {
	xnet.Init(server.Blueprint(amanda.Load()), router.Blueprint(audrey.Load()))
}

func main() {
	xnet.Load().Get("/hello", func(context *context.Context) {
		context.XWriteString("Hello")
	})
	xnet.Load().Get("/hello", func(context *context.Context) {
		context.XWriteString("Hello1")
	})
	xnet.Load().Get("/test", testXNet)
	xnet.Load().Get("/test2", func(context *context.Context) {
		context.XWriteString("test")
	})
	xnet.Load().Get("/", func(i *context.Context) {
		i.XWriteString("nihao")
	})

	xnet.Load().Run("80")

}

func testXNet(context *context.Context) {
	context.XWriteJSON(struct {
		Name string `json:"name"`
		age  int    `json:"age"`
	}{
		Name: "小明",
		age:  15,
	})
}
