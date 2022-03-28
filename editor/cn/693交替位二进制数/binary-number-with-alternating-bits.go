/**
给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。



 示例 1：


输入：n = 5
输出：true
解释：5 的二进制表示是：101


 示例 2：


输入：n = 7
输出：false
解释：7 的二进制表示是：111.

 示例 3：


输入：n = 11
输出：false
解释：11 的二进制表示是：1011.



 提示：


 1 <= n <= 2³¹ - 1

 Related Topics 位运算 👍 138 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
//对输入 nn 的二进制表示右移一位后，得到的数字再与 nn 按位异或得到 aa。当且仅当输入 nn 为交替位二进制数时，aa 的二进制表示全为 11（不包括前导 00）。这里进行简单证明：当 aa 的某一位为 11 时，当且仅当 nn 的对应位和其前一位相异。当 aa 的每一位为 11 时，当且仅当 nn 的所有相邻位相异，即 nn 为交替位二进制数。
//
//将 aa 与 a + 1a+1 按位与，当且仅当 aa 的二进制表示全为 11 时，结果为 00。这里进行简单证明：当且仅当 aa 的二进制表示全为 11 时，a + 1a+1 可以进位，并将原最高位置为 00，按位与的结果为 00。否则，不会产生进位，两个最高位都为 11，相与结果不为 00。
//
//结合上述两步，可以判断输入是否为交替位二进制数。
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/binary-number-with-alternating-bits/solution/jiao-ti-wei-er-jin-zhi-shu-by-leetcode-s-bmxd/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func hasAlternatingBits(n int) bool {
	//for pre := 2; n > 0; n /= 2 {
	//	cur := n % 2
	//	if cur == pre {
	//		return false
	//	}
	//	pre = cur
	//}
	//return true
	n = n ^ (n >> 1)
	return n&(n+1) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
