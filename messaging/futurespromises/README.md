# 未来与承诺模式

未来与承诺模式（Futures & Promises Pattern）用于异步处理任务和数据。在这种模式下，Future 表示一个未完成的操作或值，而 Promise 则表示接收该操作结果的凭证。

未来与承诺模式基本原理如下：

- Future：表示一个未完成的操作或值，充当了异步计算结果的代理；
- Promise：用于接收 Future 的计算结果，也可以设置该 Future 的结果。

Futures & Promises 模式通过将任务提交到 Goroutines 中异步执行，并使用通道进行异步通信，实现了协程之间的协作和数据交换。通过 Future 可以获取异步操作的结果，而 Promise 则用于设置该结果，实现了一种简单的异步编程模式。

在 Go 语言中，可以通过 Goroutines、通道和匿名函数等特性实现 Futures & Promises 模式。
