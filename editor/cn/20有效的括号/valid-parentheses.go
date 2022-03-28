/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

 有效字符串需满足：


 左括号必须用相同类型的右括号闭合。
 左括号必须以正确的顺序闭合。




 示例 1：


输入：s = "()"
输出：true


 示例 2：


输入：s = "()[]{}"
输出：true


 示例 3：


输入：s = "(]"
输出：false


 示例 4：


输入：s = "([)]"
输出：false


 示例 5：


输入：s = "{[]}"
输出：true



 提示：


 1 <= s.length <= 10⁴
 s 仅由括号 '()[]{}' 组成

 Related Topics 栈 字符串 👍 3082 👎 0

*/

package main

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	stack := list.New()
	pairs := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stack.PushBack(s[i])
		} else {
			e := stack.Back()
			if e != nil {
				if e.Value.(byte) != pairs[s[i]] {
					return false
				}
				stack.Remove(e)
			} else {
				// stack 当遇到右括号时，且stack为空，则返回false
				return false
			}
		}
	}
	return stack.Len() == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
