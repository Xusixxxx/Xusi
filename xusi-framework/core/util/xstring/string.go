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
    @describe 针对string操作的相关工具包
<- End */
package xstring

import (
	"bytes"
	"regexp"
	"strings"
)

/* XusiFunc ->
    @describe 获取去除最后一个字符后的字符串
    @param str string 需要处理的字符串
<- End */
func GetLast(str string) string {
	if len(str) > 1 {
		return string([]byte(str)[0 : len(str)-1])
	} else {
		return ""
	}
}

/* XusiFunc ->
    @describe 根据左边和右边的内容，获取指定字符串两者中间的内容，如果没找到则返回空字符串
    @param str string 需要查找的字符串
    @param start string 左边的标识符
    @param end string 右边的标识符
<- End */
func GetBetweenStr(str string, start string, end string) string {
	copy := str
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	if str == copy {
		return ""
	} else {
		return str
	}
}

/* XusiFunc ->
    @describe 将字符串中多个连续空格转换为一个空格，返回转换结果
    @param s string 需要转换的字符串
<- End */
func MoreSpaceToOnce(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)       //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

/* XusiFunc ->
    @describe 检查字符串是否为空字符串，返回bool
    @param str string 需要检查的字符串
<- End */
func IsEmptyString(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}

/* XusiFunc ->
    @describe 检查字符串是否为大写字母开头
    @param str string 需要检查的字符串
<- End */
func IsUpperPrefix(str string) bool {
	return ([]byte(str)[0] >= 65) && ([]byte(str)[0] <= 90)
}

/* XusiFunc ->
    @describe 格式化数字字符串为三位逗号分隔形式
    @param str string 字符串
<- End */
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}

/* XusiFunc ->
    @describe 拼接字符串
    @param args ...string 要被拼接的字符串序列
<- End */
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
