/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

 说明：本题中，我们将空字符串定义为有效的回文串。



 示例 1:


输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串


 示例 2:


输入: "race a car"
输出: false
解释："raceacar" 不是回文串




 提示：


 1 <= s.length <= 2 * 10⁵
 字符串 s 由 ASCII 字符组成

 Related Topics 双指针 字符串 👍 492 👎 0

*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; {
		if !unicode.IsLetter(int32(s[l])) && !unicode.IsDigit(int32(s[l])) {
			l++
			continue
		} else if !unicode.IsLetter(int32(s[r])) && !unicode.IsDigit(int32(s[r])) {
			r--
			continue
		} else if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
