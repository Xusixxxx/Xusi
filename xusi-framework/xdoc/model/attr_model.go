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
    @describe 支撑xdoc运行的基本模型，包含了一个package内应该正常拥有的内容
<- End */
package model

/* XusiStrcut ->
   @describe 用于支撑XDoc运行的结构体或函数参数属性模型
*/
type AttrModel struct {
	Name     string // $describe 属性名称
	Type     string // $describe 属性类型
	Addition string // $describe 属性的附加内容（如json指定等）
	Describe string // $describe 属性描述
	Index    int    // $describe 所在的位置
} // -< End
