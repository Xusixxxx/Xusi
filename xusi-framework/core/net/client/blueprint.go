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

package client

// 网络客户端设计蓝图
// 可根据蓝图拓展出HTTP Client、Socket Client等客户端
// Xusixxxx希望的所有作为Client的这种工具内容都不应该为一次性使用，那么蓝图的设计就应该围绕几点
// 应该保持工具的可复用，而不是用一次购买一次，应该是我买十个工具，那么在一个工具被使用时，我可以选择其他的工具来使用，使用结束又放回
// 在工具使用前，应该可以自行创建一些插件来改变工具的使用状态
type Blueprint interface {
}
