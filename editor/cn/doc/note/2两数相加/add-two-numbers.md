这个题构造一个空的node节点head，然后循环遍历l1和l2，当l1活着l2不为nil的时候，活着进位carry>0;则继续

关键理解：

```go
carry = l1.val + l2.val
node.value = carry % 10
carry = carry / 10
```




