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

package xrand

import (
	"math/rand"
)

/* XusiFunc ->
    @describe 随机生成Host Name
<- End */
func HostName() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	hostname := []byte{}
	r := rand.New(rand.NewSource(RandInt64(0, 999999)))
	for i := 0; i < 12; i++ {
		hostname = append(hostname, bytes[r.Intn(len(bytes))])
	}
	return string(hostname)

}
