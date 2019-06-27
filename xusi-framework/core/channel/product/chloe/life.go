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
	"runtime"
	"xusi-projects/xusi-framework/core/channel/basic"
)

var channel *Chloe

func init() {
	maxTaskFactoryNumber := (runtime.NumCPU() * 2) - 1
	channel = &Chloe{
		Channel: &basic.Channel{},
		Config: &Config{
			Config:               &basic.Config{},
			MaxTaskFactoryNumber: maxTaskFactoryNumber,
		},
		TaskPool:  make(chan chan basic.Task, maxTaskFactoryNumber),
		Result:    map[interface{}]int{},
		taskQueue: make(chan basic.Task, 2*maxTaskFactoryNumber),
	}
}

func Load() *Chloe {
	return channel
}
