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
    @describe 随机内容处理工具包
<- End */
package xrand

import (
	"math/rand"
	"strconv"
	"time"
)

/* XusiFunc ->
    @describe 置随机种子
<- End */
func Seed() {
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(int(RandInt64(0, 999)+RandInt64(0, 999)))))
}

/* XusiFunc ->
    @describe 创建一个指定范围内的int64随机数
    @param min int64 最小值
    @param max int64 最大值
<- End */
func RandInt64(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(int(max+min))))
	return min + rand.Int63n(max-min)
}

/* XusiFunc ->
    @describe 创建一个指定范围内的int随机数
    @param min int 最小值
    @param max int 最大值
<- End */
func RandInt(min int, max int) int {
	return int(RandInt64(int64(min), int64(max)))
}

/* XusiFunc ->
    @describe 创建一个指定范围内的随机数字字符串
    @param min int 最小值
    @param max int 最大值
<- End */
func RandNumber(min int, max int) string {
	return strconv.Itoa(int(RandInt64(int64(min), int64(max))))
}

/* XusiFunc ->
    @describe 创建一个指定范围内的随机数字字符串，如果长度不足最大值，则补足
    @param min int 最小值
    @param max int 最大值
<- End */
func RandNumberRepair(min int, max int) string {
	result := strconv.Itoa(int(RandInt64(int64(min), int64(max))))
	for i := len(result); i < len(strconv.Itoa(max)); i++ {
		result = "0" + result
	}
	return result
}
