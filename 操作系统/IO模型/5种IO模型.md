> https://blog.csdn.net/baiye_xing/article/details/74331041?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.channel_param

IO的本质是socket的读取，数据先拷贝到内核的缓冲区种，然后拷贝到应用程序的地址空间（进程种）

1. BIO（blocking IO）：同步阻塞IO

    当用户进程进行系统调用时， 内核就开始了IO的第一个阶段，准备数据到缓冲区中，当数据都准备完成后，则将数据从内核缓冲区中拷贝到用户进程内存中，之后用户才解除block状态重新运行

2. NIO（nonblocking IO）:同步非阻塞IO

    用户进程不会被阻塞，但是用户进程需要不停的去轮询内核，看数据是否准备好了

3. 多路复用IO（IO multiplexing）

    IO多路复用是阻塞在select,epoll的系统调用上，而没有阻塞在真正的系统调用(recvfrom)之上

4. 信号驱动IO（signal driven IO）

    当数据准备完成之后，会主动通知用户进程数据已经准备完成，即对用户进程做一个回调。通知分为两种：**水平触发**和**边缘触发**
    **水平触发：用户进程不响应就会一直发送通知**
    **边缘触发：只通知用户一次**

5. AIO异步IO（asynchronous IO）

    当用户进程发起系统调用后，立刻就可以开始去做其他的事情，然后知道IO执行的两个阶段都完成后（等待数据准备，等待数据复制都是由内核来完成操作），内核会给用户进程发送通知，告诉用户进程操作已经完成了。