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

package xnet

import (
	"xusi-projects/xusi-framework/xnet/context"
	"xusi-projects/xusi-framework/xnet/router"
	"xusi-projects/xusi-framework/xnet/router/basic"
)

func (xnet XNet) AddRoute(pattern basic.Patterns, methods basic.Methods, functions basic.Functions) {
	router.AddRoute(xnet.Router, pattern, methods, functions)
}

func (xnet XNet) Add(patterns basic.Patterns, methods basic.Methods, function func(*context.Context)) {
	router.Add(xnet.Router, patterns, methods, function)
}
func (xnet XNet) GetArr(patterns basic.Patterns, functions basic.Functions) {
	router.GetArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Get(pattern string, function func(*context.Context)) {
	router.Get(xnet.Router, pattern, function)
}
func (xnet XNet) PostArr(patterns basic.Patterns, functions basic.Functions) {
	router.PostArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Post(pattern string, function func(*context.Context)) {
	router.Post(xnet.Router, pattern, function)
}
func (xnet XNet) PutArr(patterns basic.Patterns, functions basic.Functions) {
	router.PutArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Put(pattern string, function func(*context.Context)) {
	router.Put(xnet.Router, pattern, function)
}
func (xnet XNet) DeleteArr(patterns basic.Patterns, functions basic.Functions) {
	router.DeleteArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Delete(pattern string, function func(*context.Context)) {
	router.Delete(xnet.Router, pattern, function)
}
func (xnet XNet) HeadArr(patterns basic.Patterns, functions basic.Functions) {
	router.HeadArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Head(pattern string, function func(*context.Context)) {
	router.Head(xnet.Router, pattern, function)
}
func (xnet XNet) TraceArr(patterns basic.Patterns, functions basic.Functions) {
	router.TraceArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Trace(pattern string, function func(*context.Context)) {
	router.Trace(xnet.Router, pattern, function)
}
func (xnet XNet) ConnectArr(patterns basic.Patterns, functions basic.Functions) {
	router.ConnectArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Connect(pattern string, function func(*context.Context)) {
	router.Connect(xnet.Router, pattern, function)
}
func (xnet XNet) OptionsArr(patterns basic.Patterns, functions basic.Functions) {
	router.OptionsArr(xnet.Router, patterns, functions)
}
func (xnet XNet) Options(pattern string, function func(*context.Context)) {
	router.Options(xnet.Router, pattern, function)
}
func (xnet XNet) AllArr(patterns basic.Patterns, functions basic.Functions) {
	router.AllArr(xnet.Router, patterns, functions)
}
func (xnet XNet) All(pattern string, function func(*context.Context)) {
	router.All(xnet.Router, pattern, function)
}
