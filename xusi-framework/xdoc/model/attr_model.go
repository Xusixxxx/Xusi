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

package model

// Attr文档模型
/*
 Name string 'json:"name"'  // $describe 示例
 名称   类型     附加内容             描述
*/
type AttrModel struct {
	Name     string // 名称
	Type     string // 类型
	Addition string // 附加内容
	Describe string // 描述
}
