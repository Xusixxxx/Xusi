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

package logger

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Info(a ...interface{}) {
	format(LEVEL_INFO, a)
}

func Warn(a ...interface{}) {
	format(LEVEL_WARN, a)
}

func Debug(a ...interface{}) {
	format(LEVEL_DEBUG, a)
}

func Error(a ...interface{}) {
	format(LEVEL_ERROR, a)
	var buf [2 << 10]byte
	// 去除第二三行
	result := ""
	for i, line := range strings.Split(string(buf[:runtime.Stack(buf[:], true)]), "\n") {
		if i != 1 && i != 2 {
			result += line + "\n"
		}
	}
	fmt.Print(Red, result, Reset)
}

func Fatal(a ...interface{}) {
	format(LEVEL_FATAL, a)
	var buf [2 << 10]byte
	// 去除第二三行
	result := ""
	for i, line := range strings.Split(string(buf[:runtime.Stack(buf[:], true)]), "\n") {
		if i != 1 && i != 2 {
			result += line + "\n"
		}
	}
	fmt.Print(Red, result, Reset)
}

func Trace(a ...interface{}) {
	format(LEVEL_TRACE, a)
}

// 格式化
func format(lv string, a ...interface{}) {
	var (
		date    = time.Now()
		level   = "[" + lv + "]"
		pid     = strconv.Itoa(os.Getpid())
		lvColor string
	)

	// 适配颜色
	switch lv {
	case LEVEL_TRACE:
		lvColor = Cyan
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t")
	case LEVEL_FATAL:
		lvColor = Red
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t")
	case LEVEL_ERROR:
		lvColor = Magenta
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t")
	case LEVEL_DEBUG:
		lvColor = Yellow
		pc, _, line, _ := runtime.Caller(2)
		f := runtime.FuncForPC(pc)
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t",
			Yellow, f.Name()+" :"+strconv.Itoa(line), Reset, "\t")
	case LEVEL_WARN:
		lvColor = YellowBg
		pc, _, line, _ := runtime.Caller(2)
		f := runtime.FuncForPC(pc)
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t",
			Yellow, f.Name()+" :"+strconv.Itoa(line), Reset, "\t")
	case LEVEL_INFO:
		lvColor = Blue
		fmt.Print(
			Green, date.Format("15:04:05:213"), Reset, " ",
			pid, Reset, " ",
			lvColor, level, Reset, "\t")
	}

	// 输出参数
	result := fmt.Sprint(a...)
	fmt.Println(string([]byte(result)[1 : len(result)-1]))
}
