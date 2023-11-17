# Morris-go
Implementing Three Types of Morris Traverse with Go
morris遍历的特点  时间o(n),空间o(1)
实现前中后序遍历，基本上是在morris序的基础上调节打印的顺序

先序遍历：在第一次遇到节点时打印
中序遍历：对于能来到自己两次的，在第二次打印；只能来到自己一次的，直接打印
后序遍历：打印时机是在第二次回到自己时，不过打印的是其左孩子的右边界，并且逆序打印



作者：HandsomeDemon
链接：https://juejin.cn/spost/7302254338856435731
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
