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

/* XusiPackage ->
    @describe Network Server蓝图，可根据该蓝图打造各式各样的Network Server
<- End */
package server

type Blueprint interface {
	// 运行Network Server
	Run([]string)
	RunMode([]string) string
}

func Run(blueprint Blueprint, params ...string) {
	blueprint.Run(params)
}

func RunMode(blueprint Blueprint, mode ...string) string {
	return blueprint.RunMode(mode)
}
