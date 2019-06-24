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
	"testing"
	"xusi-projects/xusi-framework/xnet/router/basic"
)

func TestRouteTableItem(t *testing.T) {
	table := map[string]basic.RouteTableItem{}
	table["1"] = basic.RouteTableItem{}
	table["3"] = basic.RouteTableItem{Patterns: []string{"123"}}
	t.Log(table["1"])
	t.Log(table["1"].Methods, len(table["1"].Methods))
	t.Log(table["2"])
	t.Log(table["1"].IsNil(), table["2"].IsNil(), table["3"].IsNil())
}
