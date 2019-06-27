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

package main

import (
	"fmt"
	"strings"
)

var result = []string{"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,良,nil,nil,nil,nil",
	"nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,良,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,不合格,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,合格,nil,优,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,不合格,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,合格,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,优,nil,nil,优,优,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil",
	"nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,良,nil,nil,nil,nil",
}

var couse = map[int]string{}

func main() {
	couse[1] = `《"美”之欣赏》`
	couse[2] = `《"美”之欣赏》`
	couse[3] = `《探秘心理世界》（预备）`
	couse[4] = `《探秘心理世界》（预备）`
	couse[5] = `《编织精彩》`
	couse[6] = `《编织精彩》`
	couse[7] = `《中学生法语口语入门》`
	couse[8] = `《中学生法语口语入门》`
	couse[9] = `《体育游戏》（预备）`
	couse[10] = `《体育游戏》（预备）`
	couse[11] = `《"小白鸽”合唱团》（需面试）（预备）`
	couse[12] = `《"小白鸽”合唱团》（需面试）（预备）`
	couse[13] = `《生活百科大解密》`
	couse[14] = `《生活百科大解密》`
	couse[15] = `《巧手手工》（预备）`
	couse[16] = `《巧手手工》（预备）`
	couse[17] = `《Scratch动画与游戏制作》`
	couse[18] = `《Scratch动画与游戏制作》`
	couse[19] = `《中国鱼类资源的介绍与保护》（预备）`
	couse[20] = `《中国鱼类资源的介绍与保护》（预备）`
	couse[21] = `《民族舞蹈》（需面试）（预备）`
	couse[22] = `《民族舞蹈》（需面试）（预备）`
	couse[23] = `《民乐合奏》（需面试）（预备）`
	couse[24] = `《民乐合奏》（需面试）（预备）`
	couse[25] = `《数字油画》（需面试）（预备）`
	couse[26] = `《数字油画》（需面试）（预备）`
	couse[27] = `中国象棋（预备）`
	couse[28] = `中国象棋（预备）`
	couse[29] = `天空之城——宫崎骏的动画世界`
	couse[30] = `天空之城——宫崎骏的动画世界`
	couse[31] = `国际象棋（预备）`
	couse[32] = `国际象棋（预备）`
	couse[33] = `启明星科技社（需面试）（预备）`
	couse[34] = `启明星科技社（需面试）（预备）`
	couse[35] = `《英语电影听力》（预备）`
	couse[36] = `《英语电影听力》（预备）`
	couse[37] = `武术（预备）`
	couse[38] = `武术（预备）`
	couse[39] = `英语戏剧（需面试）（预备）`
	couse[40] = `英语戏剧（需面试）（预备）`
	couse[41] = `爵士乐队（需面试）（预备）`
	couse[42] = `爵士乐队（需面试）（预备）`

	// 得到一行 nil,nil,nil,nil
	for _, line := range result {
		isUse := false
		count := 0
		// 分解为数组
		for i, value := range strings.Split(line, ",") {
			// 如果不等于nil，则输出，并且使用过
			if value != "nil" {
				isUse = true
				count = count + 1
				fmt.Print(couse[i] + "|" + value + "|")
			}
		}
		if count > 0 {
			fmt.Println("")
		}
		if isUse == false {
			fmt.Println("")
		}
	}
}
