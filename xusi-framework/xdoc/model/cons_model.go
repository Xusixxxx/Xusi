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

/* XusiStrcut ->
   @describe 用于支撑XDoc运行的常量模型
*/
type ConstModel struct {
	Name     string // $describe 常量名称
	Type     string // $describe 常量类型
	Value    string // $describe 常量值
	Describe string // $describe 常量描述
} // -< End
