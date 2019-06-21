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
    @describe 针对map操作的相关工具包
<- End */
package xmap

import (
	"encoding/json"
	"sort"
)

/* XusiFunc ->
    @describe 对以string作为key的map进行排序，返回所有key和错误信息
    @param mapType interface{} xmap[string]xxx...
<- End */
func SortMapByStringKey(mapType interface{}) ([]string, error) {
	strArr, err := GetKeys(mapType)
	if err != nil {
		return strArr, err
	}
	sort.Strings(strArr)
	return strArr, err
}

/* XusiFunc ->
    @describe 返回以string作为key的map的所有key和错误信息
    @param mapType interface{} xmap[string]xxx...
<- End */
func GetKeys(mapType interface{}) ([]string, error) {
	// 先将传入的内容json化
	jsonByte, err := json.Marshal(mapType)
	if err != nil {
		return nil, err
	}
	// 之后再将内容转为map
	data := make(map[string]interface{})
	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		return nil, err
	}
	// 遍历获取map的key
	strArr := []string{}
	for key, _ := range data {
		strArr = append(strArr, key)
	}
	return strArr, nil
}
