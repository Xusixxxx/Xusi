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
    @describe 处理数学计算内容的工具包
<- End */
package xmath

/* XusiFunc ->
    @describe 快速排序数组内的所有数，排序结果将返回到参数1中
  	@param nums []int 需要排序的数组
<- End */
func QuickSortAll(nums []int) {
	QuickSort(nums, 0, len(nums)-1)
}

/* XusiFunc ->
    @describe 快速排序，排序结果将返回到参数1中
    @param nums []int 需要排序的数组
    @param left int 开始的位置
    @param right int 结束的位置
<- End */
func QuickSort(nums []int, left int, right int) {
	// 取一个数作为参照，通常为数组中的第一个数
	// 这里取的则是(开始下标 + 结束下标) / 2
	// 如果长度为6，开始下标为5，结束下标为5，则取的也是5
	// 比如[5,1,52,6,27,12]，开始0，结束5
	// 那么参照数就为52，下标2
	val := nums[(left+right)/2]
	i, j := left, right
	// 如果满足了结束下标的数大于参照数
	// 那么结束下标减1
	for nums[j] > val {
		j--
	}
	// 如果满足了开始下标的数小于参照书
	// 那么开始下标加1
	for nums[i] < val {
		i++
	}
	// 替换两个数位置
	nums[i], nums[j] = nums[j], nums[i]
	i++
	j--
	// 进入下一次检查
	if i < right {
		QuickSort(nums, i, right)
	}
	if j > left {
		QuickSort(nums, left, j)
	}
}

/* XusiFunc ->
    @describe 冒泡排序，结果返回该参数中，同时返回一个数组
    @param nums []int 需要排序的数组
<- End */
func BubbleSort(nums []int) []int {
	// 获取数组长度
	length := len(nums)
	// 遍历数组所有成员
	for i := 1; i < length; i++ {
		// 使用每个成员与所有成员进行对比
		// 如果被对比的成员小于发起对比的成员，那么将他们交换位置
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return nums
}
