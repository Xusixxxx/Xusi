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

package main

import (
	"fmt"
	"time"
	"xusi-projects/xusi-framework/core/channel/product/chloe"
)

type TestStruct struct {
	result string
}

func (m *TestStruct) Exec() {
	for i := 0; i < 9999999999; i++ {
		fmt.Println("任务1处理第", i, "次")
		time.Sleep(1000 * time.Millisecond)
	}
}

type TestStruct2 struct {
	result string
}

func (m *TestStruct2) Exec() {
	for i := 0; i < 9999999999; i++ {
		fmt.Println("任务2处理第", i, "次")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	v1 := &TestStruct{}
	v2 := &TestStruct2{}
	chloe.Load().AddTask(v1)
	chloe.Load().AddTask(v2)
	chloe.Load().Run()

	for i := 0; i < 9999999999; i++ {
		fmt.Println("主线程任务处理第", i, "次")
		time.Sleep(1000 * time.Millisecond)
	}

	time.Sleep(1000 * time.Minute)
}
