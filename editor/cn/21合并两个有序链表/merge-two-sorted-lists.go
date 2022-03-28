/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



 示例 1：


输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]


 示例 2：


输入：l1 = [], l2 = []
输出：[]


 示例 3：


输入：l1 = [], l2 = [0]
输出：[0]




 提示：


 两个链表的节点数目范围是 [0, 50]
 -100 <= Node.val <= 100
 l1 和 l2 均按 非递减顺序 排列

 Related Topics 递归 链表 👍 2274 👎 0

*/

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val <= l2.Val {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	}
	return res
}

// 迭代
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	prehead := &ListNode{}
//	head := prehead
//	for l1 != nil && l2 != nil {
//		if l1.Val >= l2.Val {
//			head.Next = l2
//			l2 = l2.Next
//		}else{
//			head.Next = l1
//			l1 = l1.Next
//		}
//		head = head.Next
//	}
//
//	if l1 == nil {
//		head.Next = l2
//	}
//
//	if l2 == nil {
//		head.Next = l1
//	}
//
//	return prehead.Next
//}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

}
