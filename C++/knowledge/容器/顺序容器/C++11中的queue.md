## queue中支持的队列

* queue
	* 队列初始化
		* queue<int> myQueue;
* priority_queue
	* 优先级队列的初始化
		* priority_queue<Type, Container, Functional>
		* 示例：priority_queue<int, vector<int>, less<int> >
	
* 共同支持的操作

	top 访问队头元素
	empty 队列是否为空
	size 返回队列内元素个数
	push 插入元素到队尾 (并排序)
	emplace 原地构造一个元素并插入队列
	pop 弹出队头元素
	swap 交换内容