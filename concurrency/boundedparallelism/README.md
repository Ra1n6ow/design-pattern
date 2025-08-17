# 有界并行模式

有界并行性模式（Bounded Parallelism Pattern）用于控制并发执行任务的数量，避免系统资源被过度占用。该模式限制在同时执行的任务数量，可以有效控制并发度，避免系统过载。在 Bounded Parallelism 模式中，任务被分成多个批次，每个批次内可以同时执行的任务数量是有限的，超过限制的任务会等待前面任务完成后再执行。

提示：有界并行模式类似于并行模式，但允许在分配上设置限制。

在 Go 语言中，可以利用 Goroutines 和 Channels 结合 semaphore 控制并发执行任务的数量，实现 Bounded Parallelism 模式。semaphore 是一个计数信号量，用于控制对共享资源的访问。
