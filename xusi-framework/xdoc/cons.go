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

package xdoc

// 解析常量
const (
	HEAD_PACKAGE     = "/* XusiPackage ->" // 包头
	FOOT_PACKAGE     = "<- End */"         // 包尾
	HEAD_STRUCT      = "/* XusiStrcut ->"  // struct头
	FOOT_STRUCT      = "// -< End"         // struct尾
	FOOT_CONST_GROUP = "// -< End"         // 常量组尾
	HEAD_FUNC        = "/* XusiFunc ->"    // func头
	FOOT_FUNC        = "// -< End"         // func尾

	TAG_CONS = "// &describe" // 常量标签
	TAG_ATTR = "// $describe" // 属性标签

	SIGN_DESCRIBE = "@describe" // 描述标记
	SIGN_PARAM    = "@param"    // 描述标记
)
