/**
给定整数 n 和 k，返回 [1, n] 中字典序第 k 小的数字。



 示例 1:


输入: n = 13, k = 2
输出: 10
解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。


 示例 2:


输入: n = 1, k = 1
输出: 1




 提示:


 1 <= k <= n <= 10⁹

 Related Topics 字典树 👍 343 👎 0

*/

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/*
考虑前缀树。
字典序会给我们这样一个顺序，
1作为根节点（前缀），子节点为10 - 19（以1为前缀）;
10作为根节点，子节点为100 - 109（以10为前缀）;
以此类推

我们需要找k属于哪个根节点下的哪个子节点。
假如1为根节点的全部的节点数都不够，那么bfs到2为根节点继续找
*/
func findKthNumber(n int, k int) int {
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > n {
			return 0
		}
		if r > n {
			r = n
		}
		t := r - l + 1
		return t + dfs(l*10, r*10+9)
	}

	cur := 1
	k--
	for k > 0 {
		// dfs找到cur开头的多级子树下面一共有多少个数字在1-n之间
		cnts := dfs(cur, cur)
		// 如果cnts小于k，说明以cur开头的前缀树下面的节点树小于k哥，凑不够第k小个树，所以接着
		// 找cur+1的前缀树下面有多少节点，由于cur前缀树下已经找到了cnts个，所以在cur+1的树下面只需要找第k-cnts小的节点就ok了
		if cnts <= k {
			k -= cnts
			cur++
			// 如果cnts大于k，说明当前cur开头的前缀树下的节点数量多余k，可以找出来第k小的树，所以k--（减去当前cur这个节点），开始找以cur*10
			// 为前缀的树下面有多少个节点
		} else {
			k--
			cur *= 10
		}
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findKthNumber(13, 11))
}
