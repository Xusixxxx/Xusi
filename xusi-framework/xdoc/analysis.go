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
    @describe 维持xdoc运行的核心包，包含了xdoc的启动、渲染、常量、解析等
<- End */
package xdoc

import (
	"strings"
	"xusi-projects/xusi-framework/core/asset"
	"xusi-projects/xusi-framework/core/util"
	"xusi-projects/xusi-framework/xdoc/model"
)

// 开始解析
func startAnalysis(content string, assetFile asset.Assets) {
	// 取得逐行数据
	lines := strings.Split(content, "\n")
	// 建立文档模型
	packageModel := model.PackageModel{
		IsRoot:   false,
		Name:     "",
		Describe: "",
		Const:    map[string]model.ConstModel{},
		Struct:   map[string]model.StructModel{},
		Func:     map[string]model.FuncModel{},
		DirPath:  assetFile.DirPath,
	}

	// 开始正式解析
	packageName := analysisPackage(content, lines, &packageModel)

	// 内容解析
	analysisConst(content, lines, packageName)
	analysisStruct(content, lines, packageName)
	analysisFunc(content, lines, packageName)
}

// 解析函数
func analysisFunc(content string, lines []string, packageName string) {
	funcModel := make(map[string]model.FuncModel)
	// 函数描述开始、结束行号
	var startNumber, endNumber int
	for lineNumber, line := range lines {
		// 得到开始和结束行号
		if strings.Contains(line, HEAD_FUNC) {
			startNumber = lineNumber
		}
		if (startNumber != 0) && strings.HasPrefix(line, "func ") {
			endNumber = lineNumber + 1
		}
		// 解析方法体
		if startNumber != 0 && endNumber != 0 {
			models := model.FuncModel{
				Params: map[string]model.AttrModel{},
			}
			// 建立参数描述存储池
			// 参数描述会在参数前解析，所以先存储
			paramDescribePool := map[string][]string{}
			for i := startNumber; i < endNumber; i++ {
				formatLine := util.MoreSpaceToOnce(lines[i])
				// 解析名称
				if strings.HasPrefix(formatLine, "func") {
					funcBody := strings.TrimSpace(strings.Replace(formatLine, "func", "", 1))
					// 检测是否为括号开头，如果是，那么函数名称在第一个括号后面
					// 如果不是，函数名称在第一个括号前面
					if strings.HasPrefix(funcBody, "(") {
						// 函数名在括号后
						// 切片取函数名
						// 0 ： (xusi The
						// 1 ：  Init(mode string, isDebug bool) Application {
						slice := strings.SplitN(funcBody, ")", 2)
						models.Name = strings.TrimSpace(util.MoreSpaceToOnce(strings.SplitN(strings.TrimSpace(slice[1]), "(", 2)[0]))
					} else {
						// 函数名在括号前
						// 切片
						// 0 ：Init
						// 1 ：mode string, isDebug bool) Application {
						models.Name = util.MoreSpaceToOnce(strings.TrimSpace(strings.SplitN(funcBody, "(", 2)[0]))
					}

					// 根据名字切片得到参数
					paramSlice := strings.Split(strings.SplitN(strings.SplitN(strings.Split(formatLine, models.Name)[1], "(", 2)[1], ")", 2)[0], ",")
					for i, param := range paramSlice {
						if param == "" {
							continue
						}
						nameAndTypeSlice := strings.SplitN(strings.TrimSpace(util.MoreSpaceToOnce(param)), " ", 2)
						if strings.Contains(nameAndTypeSlice[1], "(") {
							nameAndTypeSlice[1] += ")"
						}
						if _, ok := models.Params[nameAndTypeSlice[0]]; ok {
							p := models.Params[nameAndTypeSlice[0]]
							p.Name = nameAndTypeSlice[0]
							p.Type = nameAndTypeSlice[1]
							p.Index = i
							models.Params[p.Name] = p
						} else {
							models.Params[nameAndTypeSlice[0]] = model.AttrModel{
								Name:  nameAndTypeSlice[0],
								Type:  nameAndTypeSlice[1],
								Index: i,
							}
						}
					}
				}
				// 解析描述
				if strings.HasPrefix(strings.TrimSpace(formatLine), SIGN_DESCRIBE) {
					models.Describe = util.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(formatLine, SIGN_DESCRIBE, "")))
				}
				// 解析参数描述
				if strings.HasPrefix(strings.TrimSpace(formatLine), SIGN_PARAM) {
					typeAndDescribeSlice := strings.SplitN(util.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(formatLine, SIGN_PARAM, ""))), " ", 2)
					// 如果切片量大于2，则标识参数类型有携带空格，如 func(xusi Xusi)的函数参数
					sliceTempN := strings.Split(typeAndDescribeSlice[1], " ")
					// 拼接第1个切片到倒数第2个
					tempStr := ""
					for _, value := range sliceTempN[0 : len(sliceTempN)-1] {
						tempStr += string(value) + " "
					}
					sliceTemp := []string{tempStr, sliceTempN[len(sliceTempN)-1]}
					switch len(sliceTemp) {
					case 1:
						typeAndDescribeSlice = []string{typeAndDescribeSlice[0], sliceTemp[0]}
					case 2:
						typeAndDescribeSlice = []string{typeAndDescribeSlice[0], sliceTemp[0], sliceTemp[1]}
					}
					switch len(typeAndDescribeSlice) {
					case 1:
						paramDescribePool[typeAndDescribeSlice[0]] = []string{"", ""}
					case 2:
						paramDescribePool[typeAndDescribeSlice[0]] = []string{typeAndDescribeSlice[1], ""}
					case 3:
						paramDescribePool[typeAndDescribeSlice[0]] = []string{typeAndDescribeSlice[1], typeAndDescribeSlice[2]}
					}
				}
			}
			startNumber = 0
			endNumber = 0

			// 取出存储池的内容赋值
			for key, value := range paramDescribePool {
				if _, ok := models.Params[key]; ok {
					p := models.Params[key]
					p.Name = key
					p.Type = value[0]
					p.Describe = value[1]
					models.Params[p.Name] = p
				} else {
					models.Params[key] = model.AttrModel{
						Name:     key,
						Type:     value[0],
						Describe: value[1],
					}
				}
			}

			funcModel[models.Name] = models
		}
	}
	// 合并map
	for _, value := range funcModel {
		if !util.IsEmptyString(value.Name) && util.IsUpperPrefix(value.Name) {
			Docs[packageName].Func[value.Name] = value
		}
	}
}

// 解析结构体
func analysisStruct(content string, lines []string, packageName string) {
	structModels := make(map[string]model.StructModel)
	// 结构体描述开始、结束行号
	var startNumber, endNumber int
	for lineNumber, line := range lines {
		// 得到结构体的开始结束行号，开始解析
		if strings.Contains(line, HEAD_STRUCT) {
			startNumber = lineNumber
		}
		if (startNumber != 0) && strings.HasSuffix(line, FOOT_STRUCT) {
			endNumber = lineNumber
		}
		// 解析结构体
		if startNumber != 0 && endNumber != 0 {
			modeler := model.StructModel{
				Attr: map[string]model.AttrModel{},
			}
			for i := startNumber; i <= endNumber; i++ {
				childrenLine := lines[i]
				// 解析名称
				// 如果不为公开型结构体，直接忽略跳过
				if strings.HasPrefix(childrenLine, "type") {
					modeler.Name = strings.TrimSpace(util.GetBetweenStr(childrenLine, "type", "struct"))
				}
				// 解析描述
				if strings.HasPrefix(strings.TrimSpace(childrenLine), SIGN_DESCRIBE) {
					modeler.Describe = util.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(childrenLine, SIGN_DESCRIBE, "")))
				}
				// 解析属性
				if strings.Contains(childrenLine, TAG_ATTR) {
					attrModel := model.AttrModel{}
					// 切片1为名称和类型
					// 切片2为附加属性和描述
					// Host  string | json:"host"` // $describe 主机
					// 如果只有1个切片，标识没有附加属性
					attrAndDescribeSlice := strings.SplitN(util.MoreSpaceToOnce(childrenLine), "`", 2)
					// 如果不为公开属性，则直接忽略跳过
					// 如果没有编写属性名称，也直接忽略跳过
					// 切片1为名称
					// 切片2为类型
					// 如果只有1个切片，则表示没有名称
					nameAndTypeSlice := strings.Split(strings.TrimSpace(attrAndDescribeSlice[0]), " ")
					if (len(nameAndTypeSlice) == 1) || !util.IsUpperPrefix(childrenLine) {
						continue
					} else {
						// 赋值
						switch len(attrAndDescribeSlice) {
						case 1:
							attrModel.Describe = strings.TrimSpace(strings.ReplaceAll(attrAndDescribeSlice[0], TAG_ATTR, ""))
						case 2:
							AdditionAndDescribe := strings.Split(attrAndDescribeSlice[1], "`")
							attrModel.Describe = strings.TrimSpace(strings.ReplaceAll(AdditionAndDescribe[1], TAG_ATTR, ""))
							attrModel.Addition = "`" + util.MoreSpaceToOnce(strings.TrimSpace(AdditionAndDescribe[0])) + "`"
						}
						attrModel.Name = nameAndTypeSlice[0]
						attrModel.Type = nameAndTypeSlice[1]

						modeler.Attr[attrModel.Name] = attrModel
					}
				}
			}
			structModels[modeler.Name] = modeler
			startNumber = 0
			endNumber = 0
		}
	}
	// 合并map
	for _, value := range structModels {
		if !util.IsEmptyString(value.Name) && util.IsUpperPrefix(value.Name) {
			Docs[packageName].Struct[value.Name] = value
		}
	}
}

// 解析常量
func analysisConst(content string, lines []string, packageName string) {
	constModels := make(map[string]model.ConstModel)
	var startNumber, endNumber int
	for lineNumber, line := range lines {
		// 解析常量组
		// 如果为常量组 const ( ... )
		if strings.Contains(strings.TrimSpace(util.MoreSpaceToOnce(line)), "const (") {
			startNumber = lineNumber + 1 // 加1为了忽略掉 const ( 行
			continue
		}
		if (startNumber != 0) && strings.HasSuffix(util.MoreSpaceToOnce(strings.TrimSpace(line)), FOOT_CONST_GROUP) {
			endNumber = lineNumber
			continue
		}
		// 解析常量组
		if (startNumber != 0) && (endNumber != 0) {
			for i := startNumber; i < endNumber; i++ {
				if !strings.Contains(lines[i], TAG_CONS) {
					continue
				}
				model, isOpen := constFormat(lines[i])
				if isOpen {
					constModels[model.Name] = model
				} else {
					continue
				}
			}
			startNumber = 0
			endNumber = 0
		}

		// 解析普通常量
		if strings.Contains(line, TAG_CONS) {
			// 如果为const + 空格开头，清除该字符串
			if strings.HasPrefix(line, "const ") {
				model, isOpen := constFormat(strings.Replace(line, "const ", "", -1))
				if isOpen {
					constModels[model.Name] = model
				} else {
					continue
				}
			}
		}
	}
	// 合并map
	for _, value := range constModels {
		Docs[packageName].Const[value.Name] = value
	}
}

// 解析包名
func analysisPackage(content string, lines []string, packageModel *model.PackageModel) string {
	var packageName, packageDescribe string

	// 取到包名
	for _, line := range lines {
		if strings.HasPrefix(line, "package ") {
			packageName = strings.ReplaceAll(line, "package ", "")
			packageName = strings.TrimSpace(packageName)
			break
		}
	}

	// 取到包描述
	// 遍历每一行，如果这一行只为 HEAD_PACKAGE || FOOT_PACKAGE
	// 那么记录开始和结尾行号再遍历
	var startNumber, endNumber int
	for index, line := range lines {
		// 记录开始行号
		if strings.TrimSpace(line) == HEAD_PACKAGE {
			startNumber = index
		}
		if startNumber != 0 && strings.TrimSpace(line) == FOOT_PACKAGE {
			endNumber = index
		}
		// 遍历内容区
		/* XusiPackage ->
		    @describe Xusi
		<- End */
		if startNumber != 0 && endNumber != 0 {
			for i := startNumber; i < endNumber; i++ {
				if strings.HasPrefix(lines[i], SIGN_DESCRIBE) {
					packageDescribe = util.MoreSpaceToOnce(strings.TrimSpace(strings.ReplaceAll(lines[i], SIGN_DESCRIBE, "")))
					break
				}
			}
			startNumber = 0
			endNumber = 0
		}
	}

	// 赋值
	packageModel.Name = packageName
	packageModel.Describe = packageDescribe

	// 如果存在
	if _, ok := Docs[packageModel.GetPackagePath()]; ok {
		pModel := Docs[packageModel.GetPackagePath()]
		if !util.IsEmptyString(packageDescribe) {
			pModel.Describe = packageDescribe
		}
		for key, value := range packageModel.Const {
			pModel.Const[key] = value
		}
		for key, value := range packageModel.Struct {
			pModel.Struct[key] = value
		}
		for key, value := range packageModel.Func {
			pModel.Func[key] = value
		}
		Docs[packageModel.GetPackagePath()] = pModel
	} else {
		Docs[packageModel.GetPackagePath()] = *packageModel
	}

	return packageModel.GetPackagePath()
}

func constFormat(formatLine string) (model.ConstModel, bool) {
	model := model.ConstModel{}
	// 对一个出现的“=”符号进行切割，需要最终结果为两个切片
	// 并且需要对切片清理前后多余空格
	// slice1 = 常量名字 or 常量名字 + 类型      	CONS_NAME / CONST_NAME string
	// slice2 = 常量值 + 描述					"value" // &describe 生产环境
	slice := strings.SplitN(formatLine, "=", 2)
	slice[0] = util.MoreSpaceToOnce(strings.TrimSpace(slice[0]))
	slice[1] = strings.TrimSpace(slice[1])
	// 解析切片1
	// 如果只有名称，那么切片数量为1，并且值为本身
	// 如果包含了类型，那么切片数量为2，分别为名称和类型
	// 也就数量代表了”1“为只有名称，”2“为包含了名称和类型
	nameAndTypeSlice := strings.Split(slice[0], " ")
	// 如果不为公开内容，则忽略，进入下一个
	if !util.IsUpperPrefix(nameAndTypeSlice[0]) {
		return model, false
	}
	switch len(nameAndTypeSlice) {
	case 1:
		model.Name = nameAndTypeSlice[0]
	case 2:
		model.Name = nameAndTypeSlice[0]
		model.Type = nameAndTypeSlice[1]
	}
	// 解析切片2
	// 对切片2进行再切片，子切片中，第一个为”Value“，第二个则为”Describe”和“Type“
	// 对两个子切片进行格式化
	// 让值遵循原有意愿，但是需要对描述进行无意义字符串的格式化
	valueAndDescribeSlice := strings.SplitN(slice[1], TAG_CONS, 2)
	valueAndDescribeSlice[0] = strings.TrimSpace(valueAndDescribeSlice[0])
	valueAndDescribeSlice[1] = util.MoreSpaceToOnce(strings.TrimSpace(valueAndDescribeSlice[1]))
	// 赋值
	model.Value = valueAndDescribeSlice[0]
	model.Describe = valueAndDescribeSlice[1]
	if util.IsEmptyString(model.Type) {
		model.Type = "unknown"
	}
	return model, true
}
