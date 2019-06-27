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

package util

import "xusi-projects/xusi-framework/core/util/xstring"

// HTTP请求数据解析
func HttpDataMapAnalysis(data map[string]string) string {
	result := ""
	for key, value := range data {
		result += key + "=" + value + "&"
	}
	return xstring.GetLast(result)
}
