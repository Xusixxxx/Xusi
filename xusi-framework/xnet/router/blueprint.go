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
    @describe 路由器蓝图，可根据该蓝图打造各式各样的路由器
<- End */
package router

import (
	"xusi-projects/xusi-framework/xnet/context"
	"xusi-projects/xusi-framework/xnet/router/basic"
)

type Blueprint interface {
	// 查找路由
	Find(*context.Context) []func(*context.Context)
	// 添加路由
	AddRoute(basic.Patterns, basic.Methods, basic.Functions)
	Add(basic.Patterns, basic.Methods, func(*context.Context))
	GetArr(basic.Patterns, basic.Functions)
	Get(string, func(*context.Context))
	PostArr(basic.Patterns, basic.Functions)
	Post(string, func(*context.Context))
	PutArr(basic.Patterns, basic.Functions)
	Put(string, func(*context.Context))
	DeleteArr(basic.Patterns, basic.Functions)
	Delete(string, func(*context.Context))
	HeadArr(basic.Patterns, basic.Functions)
	Head(string, func(*context.Context))
	TraceArr(basic.Patterns, basic.Functions)
	Trace(string, func(*context.Context))
	ConnectArr(basic.Patterns, basic.Functions)
	Connect(string, func(*context.Context))
	OptionsArr(basic.Patterns, basic.Functions)
	Options(string, func(*context.Context))
	AllArr(basic.Patterns, basic.Functions)
	All(string, func(*context.Context))
}

func AddRoute(blueprint Blueprint, pattern basic.Patterns, methods basic.Methods, functions basic.Functions) {
	blueprint.AddRoute(pattern, methods, functions)
}

func Add(blueprint Blueprint, patterns basic.Patterns, methods basic.Methods, function func(*context.Context)) {
	blueprint.Add(patterns, methods, function)
}
func GetArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.GetArr(patterns, functions)
}
func Get(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Get(pattern, function)
}
func PostArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.PostArr(patterns, functions)
}
func Post(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Post(pattern, function)
}
func PutArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.PutArr(patterns, functions)
}
func Put(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Put(pattern, function)
}
func DeleteArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.DeleteArr(patterns, functions)
}
func Delete(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Delete(pattern, function)
}
func HeadArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.HeadArr(patterns, functions)
}
func Head(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Head(pattern, function)
}
func TraceArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.TraceArr(patterns, functions)
}
func Trace(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Trace(pattern, function)
}
func ConnectArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.ConnectArr(patterns, functions)
}
func Connect(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Connect(pattern, function)
}
func OptionsArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.OptionsArr(patterns, functions)
}
func Options(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.Options(pattern, function)
}
func AllArr(blueprint Blueprint, patterns basic.Patterns, functions basic.Functions) {
	blueprint.AllArr(patterns, functions)
}
func All(blueprint Blueprint, pattern string, function func(*context.Context)) {
	blueprint.All(pattern, function)
}
