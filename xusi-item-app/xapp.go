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

// Xusi 拓展功能项目
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"xusi-projects/xusi-framework/core/util"
	build2 "xusi-projects/xusi-item-app/build"
)

const VERSION = "1.0.0"

// 入参变量
var (
	help    bool   // 帮助
	version bool   // 版本
	create  string // 创建App
	build   bool   // 构建App
)

func init() {
	flag.BoolVar(&help, "help", false, "show xusi help")
	flag.BoolVar(&version, "version", false, "show xusi version")
	flag.StringVar(&create, "create", "", "create a xusi app")
	flag.BoolVar(&build, "build", false, "bulid xusi app")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {

	flag.Parse()

	if help {
		flag.Usage()
	} else if version {
		fmt.Print("xusi-app version:" + VERSION)
	} else if build {
		fmt.Println("start build xusi app...")
		runPath, err := util.GetRunPath()
		if err != nil {
			panic(err)
		}
		fmt.Println("build xusi app : " + runPath + "\r\n")

		// 开始时间
		startTime := time.Now().UnixNano() / 1e6
		fmt.Println("ready start xusi app build, now date:", startTime)
		// 开始构建
		err = build2.Build(runPath)
		if err != nil {
			panic(err)
		} else {
			// 构建结束
			fmt.Println("build success...\r\n", "consume :", (time.Now().UnixNano()/1e6)-startTime, "ms...")
		}
	} else if create != "" {

	} else {
		return
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `
xusi-app version: `+VERSION+`
Usage: xusi [-help] [-version] [-create appName] [-build] ...

Options:
`)
	flag.PrintDefaults()
}
