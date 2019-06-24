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

/* XusiPackage ->
    @describe 记录器，用于支撑日志功能
<- End */
package logger

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/* XusiFunc ->
    @describe 输出Info日志
    @param a ...interface{} 输出的内容
<- End */
func Info(a ...interface{}) {
	if Conf.Disable {
		return
	}
	format(LEVEL_INFO, a)
}

/* XusiFunc ->
    @describe 输出Warn日志
    @param a ...interface{} 输出的内容
<- End */
func Warn(a ...interface{}) {
	if Conf.Disable {
		return
	}
	format(LEVEL_WARN, a)
}

/* XusiFunc ->
    @describe 输出Debug日志
    @param a ...interface{} 输出的内容
<- End */
func Debug(a ...interface{}) {
	if Conf.Disable {
		return
	}
	format(LEVEL_DEBUG, a)
}

/* XusiFunc ->
    @describe 输出Error日志
    @param a ...interface{} 输出的内容
<- End */
func Error(a ...interface{}) {
	if Conf.Disable {
		return
	}
	format(LEVEL_ERROR, a)
}

/* XusiFunc ->
    @describe 输出Fatal日志
    @param a ...interface{} 输出的内容
<- End */
func Fatal(a ...interface{}) {
	if Conf.Disable {
		return
	}
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

/* XusiFunc ->
    @describe 输出Trace日志
    @param a ...interface{} 输出的内容
<- End */
func Trace(a ...interface{}) {
	if Conf.Disable {
		return
	}
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
		if Conf.Mode == MODE_DEV {
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t")
		} else {
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t")
		}
	case LEVEL_FATAL:
		lvColor = Red
		if Conf.Mode == MODE_DEV {
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t")
		} else {
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t")
		}

	case LEVEL_ERROR:
		lvColor = Magenta
		if Conf.Mode == MODE_DEV {
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t")
		} else {
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t")
		}

	case LEVEL_DEBUG:
		lvColor = Yellow
		if Conf.Mode == MODE_DEV {
			pc, _, line, _ := runtime.Caller(2)
			f := runtime.FuncForPC(pc)
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t",
				Yellow, f.Name()+" :"+strconv.Itoa(line), Reset, "\t")
		} else {
			pc, _, line, _ := runtime.Caller(2)
			f := runtime.FuncForPC(pc)
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t",
				f.Name()+" :"+strconv.Itoa(line), "\t")
		}

	case LEVEL_WARN:
		if Conf.Mode == MODE_DEV {
			lvColor = YellowBg
			pc, _, line, _ := runtime.Caller(2)
			f := runtime.FuncForPC(pc)
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t",
				Yellow, f.Name()+" :"+strconv.Itoa(line), Reset, "\t")
		} else {
			pc, _, line, _ := runtime.Caller(2)
			f := runtime.FuncForPC(pc)
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t",
				f.Name()+" :"+strconv.Itoa(line), "\t")
		}

	case LEVEL_INFO:
		lvColor = Blue
		if Conf.Mode == MODE_DEV {
			fmt.Print(
				Green, date.Format("15:04:05:213"), Reset, " ",
				pid, Reset, " ",
				lvColor, level, Reset, "\t")
		} else {
			fmt.Print(
				date.Format("15:04:05:213"), " ",
				pid, " ",
				level, "\t")
		}

	}

	// 输出参数
	result := fmt.Sprint(a...)
	if Conf.Mode != MODE_DEV {
		//移除logger字符
		result = strings.ReplaceAll(result, Blue, "")
		result = strings.ReplaceAll(result, BlueBg, "")
		result = strings.ReplaceAll(result, Yellow, "")
		result = strings.ReplaceAll(result, YellowBg, "")
		result = strings.ReplaceAll(result, Red, "")
		result = strings.ReplaceAll(result, RedBg, "")
		result = strings.ReplaceAll(result, Green, "")
		result = strings.ReplaceAll(result, GreenBg, "")
		result = strings.ReplaceAll(result, Magenta, "")
		result = strings.ReplaceAll(result, MagentaBg, "")
		result = strings.ReplaceAll(result, Reset, "")
		result = strings.ReplaceAll(result, Cyan, "")
		result = strings.ReplaceAll(result, CyanBg, "")
	}
	fmt.Println(string([]byte(result)[1 : len(result)-1]))
}
