/**
给你一个字符串 s，找到 s 中最长的回文子串。



 示例 1：


输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。


 示例 2：


输入：s = "cbbd"
输出："bb"




 提示：


 1 <= s.length <= 1000
 s 仅由数字和英文字母组成

 Related Topics 字符串 动态规划 👍 4867 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {

		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}

		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]

}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left -= 1
			right += 1
		} else {
			break
		}
	}
	return left + 1, right - 1
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
