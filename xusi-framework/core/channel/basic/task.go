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

package basic

// 需要处理的任务数据类型
type Task struct {
	TaskImpl          // 任务结构
	Wait     chan int // 等待chan
	isWait   bool     // 是否需要等待
}

// 任务需要满足接口
type TaskImpl interface {
	Exec() // 执行
}

// 创建任务(分为是否需要等待执行完毕的任务)
func CreateTask(taskImpl TaskImpl, wait ...chan int) Task {
	if len(wait) >= 1 {
		return Task{TaskImpl: taskImpl, Wait: wait[0], isWait: true}
	} else {
		return Task{TaskImpl: taskImpl, isWait: false}
	}
}
