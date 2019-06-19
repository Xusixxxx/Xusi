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

package util

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

/* XusiFunc ->
    @describe 对数据进行GZip压缩，返回bytes.Buffer和错误信息
    @param data []byte 需要压缩的数据
<- End */
func GZipCompress(data []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	// 写入压缩
	_, err := gzipWriter.Write(data)
	if err != nil {
		return buf, err
	}
	// 关闭压缩
	if err := gzipWriter.Close(); err != nil {
		return buf, err
	}
	return buf, nil
}

/* XusiFunc ->
    @describe 对已进行GZip压缩的数据进行解压缩，返回字节数组及错误信息
    @param dataByte []byte 需要进行解压缩的数据
<- End */
func UnGZipCompress(dataByte []byte) ([]byte, error) {
	data := *bytes.NewBuffer(dataByte)
	// 读取压缩过的数据
	gzipReader, err := gzip.NewReader(&data)
	if err != nil {
		return nil, err
	}
	// 读取数据
	result, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	// 关闭读取
	if err := gzipReader.Close(); err != nil {
		return nil, err
	}

	return result, nil
}
