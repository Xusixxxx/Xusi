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

package chloe

import (
	"xusi-projects/xusi-framework/core/channel/basic"
)

type Chloe struct {
	*basic.Channel
	*Config
	TaskPool        chan chan basic.Task // 任务池
	Result          map[interface{}]int  // 存放任务处理结果的结果集
	taskQueue       chan basic.Task      // 待处理的任务通道
	TaskFactoryList []basic.TaskFactory  // 任务工厂列表（主要用于停止任务）
}

// 运行Chloe
func (chloe Chloe) Run() {
	//启动所有的Worker
	for i := 0; i < channel.Config.MaxTaskFactoryNumber; i++ {
		taskFactory := basic.CreateTaskFactory(chloe.TaskPool, chloe.Result, i)
		chloe.TaskFactoryList = append(chloe.TaskFactoryList, taskFactory)
		taskFactory.Start()
	}
	go chloe.dispatch()
}

// 派遣分发任务给任务工厂进行处理
func (chloe Chloe) dispatch() {
	for {
		select {
		case task := <-chloe.taskQueue:
			go func(task basic.Task) {
				//从任务工厂池中取出一个任务通道
				taskChannel := <-chloe.TaskPool
				//向这个通道中发送任务，任务工厂中的接收配对操作会被唤醒
				taskChannel <- task
			}(task)
		}
	}
}

// 添加任务到任务通道
func (chloe Chloe) AddTask(task basic.TaskImpl, wait ...chan int) {
	// 向任务通道发送任务
	chloe.taskQueue <- basic.CreateTask(task, wait...)
}

//停止所有任务
func (chloe Chloe) Stop() {
	for _, taskFactory := range chloe.TaskFactoryList {
		taskFactory.Stop()
	}
}
