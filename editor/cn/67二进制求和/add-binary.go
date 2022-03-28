/**
给你两个二进制字符串，返回它们的和（用二进制表示）。

 输入为 非空 字符串且只包含数字 1 和 0。



 示例 1:

 输入: a = "11", b = "1"
输出: "100"

 示例 2:

 输入: a = "1010", b = "1011"
输出: "10101"



 提示：


 每个字符串仅由字符 '0' 或 '1' 组成。
 1 <= a.length, b.length <= 10^4
 字符串如果不是 "0" ，就都不含前导零。

 Related Topics 位运算 数学 字符串 模拟 👍 770 👎 0

*/

package main

import (
	"fmt"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {

	res := ""
	carry := 0
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		t := int(a[i]-'0') + int(b[j]-'0') + carry
		res = strconv.Itoa(t%2) + res
		carry = t / 2
	}

	for ; i >= 0; i = i - 1 {
		t := int(a[i]-'0') + carry
		res = strconv.Itoa(t%2) + res
		carry = t / 2
	}

	for ; j >= 0; j = j - 1 {
		t := int(b[j]-'0') + carry
		res = strconv.Itoa(t%2) + res
		carry = t / 2
	}

	if carry != 0 {
		res = "1" + res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := "11"
	b := "1"
	addBinary(a, b)
	fmt.Println("res", 10101)
}
