> https://blog.csdn.net/hello_wyq/article/details/1108911?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param

> https://www.cnblogs.com/voipman/p/7520965.html

设置线程优先级的函数：

int pthread_setschedparam(pthread_t target_thread, int policy, const struct sched_param *param)

它主要用于设置线程的调用策略和优先级。

参数说明：

1.  target_thread是使用pthread_create所获得的线程ID。
 
2. 线程的调度有三种策略：SCHED_OTHER、SCHED_RR和SCHED_FIFO。
 
Policy用于指明使用哪种策略。下面我们简单的说明一下这三种调度策略。
SCHED_OTHER
它是默认的线程分时调度策略，所有的线程的优先级别都是0，线程的调度是通过分时来完成的。简单地说，如果系统使用这种调度策略，程序将无法设置线程的优先级。请注意，这种调度策略也是抢占式的，当高优先级的线程准备运行的时候，当前线程将被抢占并进入等待队列。这种调度策略仅仅决定线程在可运行线程队列中的具有相同优先级的线程的运行次序。
 
SCHED_FIFO
它是一种实时的先进先出调用策略，且只能在超级用户下运行。这种调用策略仅仅被使用于优先级大于0的线程。它意味着，使用SCHED_FIFO的可运行线程将一直抢占使用SCHED_OTHER的运行线程J。此外SCHED_FIFO是一个非分时的简单调度策略，当一个线程变成可运行状态，它将被追加到对应优先级队列的尾部((POSIX 1003.1)。当所有高优先级的线程终止或者阻塞时，它将被运行。对于相同优先级别的线程，按照简单的先进先运行的规则运行。我们考虑一种很坏的情况，如果有若干相同优先级的线程等待执行，然而最早执行的线程无终止或者阻塞动作，那么其他线程是无法执行的，除非当前线程调用如pthread_yield之类的函数，所以在使用SCHED_FIFO的时候要小心处理相同级别线程的动作。
 
SCHED_RR
鉴于SCHED_FIFO调度策略的一些缺点，SCHED_RR对SCHED_FIFO做出了一些增强功能。从实质上看，它还是SCHED_FIFO调用策略。它使用最大运行时间来限制当前进程的运行，当运行时间大于等于最大运行时间的时候，当前线程将被切换并放置于相同优先级队列的最后。这样做的好处是其他具有相同级别的线程能在“自私“线程下执行。
 
3．param是struct sched_param类型的指针，它仅仅包含一个成员变sched_priority，指明所要设置的静态线程优先级。