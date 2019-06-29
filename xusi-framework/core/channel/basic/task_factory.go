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

// 任务处理工厂结构
type TaskFactory struct {
	id          int                 // 工作id
	WorkPool    chan chan Task      // 工作队列
	TaskChannel chan Task           // 从TaskChannel中获取任务进行处理
	Result      map[interface{}]int // 将工作结果放入result
	quit        chan bool           // 停止工作的信号
}

// 创建一个任务工厂
func CreateTaskFactory(workerPool chan chan Task, result map[interface{}]int, id int) TaskFactory {
	return TaskFactory{
		id:          id,
		WorkPool:    workerPool,
		TaskChannel: make(chan Task),
		Result:      result,
		quit:        make(chan bool),
	}
}

// 开始处理队列任务
func (taskFactory TaskFactory) Start() {
	go func() {
		for {
			// 将任务添加到工作队列中
			taskFactory.WorkPool <- taskFactory.TaskChannel

			select {
			// 从任务通道中获取到任务进行处理，任务在处理期间由于TaskChannel为同步通道，会进行阻塞
			case task := <-taskFactory.TaskChannel:
				// 收到任务处理通知，开始处理
				task.TaskImpl.Exec()
				// 执行结束返回等待消息
				if task.isWait {
					task.Wait <- 1
				}
			// 当收到任务终止的消息
			case <-taskFactory.quit:
				return
			}
		}
	}()
}

// 终止任务
func (taskFactory TaskFactory) Stop() {
	go func() {
		taskFactory.quit <- true
	}()
}
