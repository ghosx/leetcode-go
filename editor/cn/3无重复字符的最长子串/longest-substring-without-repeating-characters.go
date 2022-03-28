/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



 示例 1:


输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


 示例 2:


输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。


 示例 3:


输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。




 提示：


 0 <= s.length <= 5 * 10⁴
 s 由英文字母、数字、符号和空格组成

 Related Topics 哈希表 字符串 滑动窗口 👍 7113 👎 0

*/

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i, i2 := range s {
		fmt.Println(i, string(i2))
		index := strings.Index(s[start:i], string(i2))
		//fmt.Println("index", index)
		if index != -1 {
			start += index + 1
			end += index + 1
		} else {
			if i+1 > end {
				end = i + 1
			}
		}
	}
	//fmt.Println(start, end)
	return end - start
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := lengthOfLongestSubstring("asdasdasd哈")
	fmt.Println(a)
}
