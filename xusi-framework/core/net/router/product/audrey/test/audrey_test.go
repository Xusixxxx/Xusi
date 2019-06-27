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

package test

import (
	"net/http"
	"testing"
	"xusi-projects/xusi-framework/core/net/httplibs"
	basic2 "xusi-projects/xusi-framework/core/net/router/basic"
	"xusi-projects/xusi-framework/core/net/router/product/audrey"
)

func TestAdd(t *testing.T) {
	fun1 := func(responseWriter http.ResponseWriter, r *http.Request) {}
	fun2 := func(responseWriter http.ResponseWriter, r *http.Request) {}
	fun3 := func(responseWriter http.ResponseWriter, r *http.Request) {}

	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_GET}, func(responseWriter http.ResponseWriter, r *http.Request) {})
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_POST}, func(responseWriter http.ResponseWriter, r *http.Request) {})
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_GET}, fun1)
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_POST}, fun1)
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_GET}, fun1)
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_DELETE}, fun2)
	audrey.Load().Add(basic2.Patterns{"/xusi"}, basic2.Methods{httplibs.METHOD_GET}, fun3)
	audrey.Load().Add(basic2.Patterns{"/xusi2"}, basic2.Methods{httplibs.METHOD_GET}, func(responseWriter http.ResponseWriter, r *http.Request) {})
	for _, value := range audrey.Load().Router.Table {
		t.Log(value)
	}
}
