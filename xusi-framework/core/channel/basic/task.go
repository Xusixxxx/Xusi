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
	TaskStruct TaskImpl // 任务结构
}

// 任务需要满足接口
type TaskImpl interface {
}

// 创建任务
func CreateTask(taskStruct TaskImpl) Task {
	return Task{TaskStruct: taskStruct}
}
