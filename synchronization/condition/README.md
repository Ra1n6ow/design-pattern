# 条件变量模式

条件变量模式（Condition Variable Pattern）用于线程或协程之间的同步和通信。在该模式中，一个或多个线程或协程等待某个特定条件的变化，然后再继续执行其他操作。Condition Variable 模式通常与 Mutex 或 RWMutex 配合使用，用于实现线程安全的等待和通知机制。

在 Condition Variable 模式中，通常包含两个主要操作：Wait 和 Signal。Wait 操作用于使一个线程或协程等待某个条件的满足，当条件不满足时会阻塞当前线程或协程。Signal 操作用于通知等待的线程或协程条件已经发生变化，使其可以继续执行。
