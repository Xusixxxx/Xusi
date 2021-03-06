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
    @describe 针对package操作的相关工具包
<- End */
package xpackage

import "reflect"

/* XusiFunc ->
    @describe 根据传入对象获取指定package路径，返回string类型的路径信息
    @param obj interface{} 需要获取的package中的任一对象
<- End */
func GetPackagePath(obj interface{}) string {
	return reflect.TypeOf(obj).PkgPath()
}
