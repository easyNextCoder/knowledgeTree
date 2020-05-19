#include <iostream>
#include <vector>
#include <list>
#include <queue>//priority_queue
#include <deque>
#include <stack>

using namespace std;

int main()
{
	//vector
	cout<<"test vector."<<endl;
	vector<int>vec(10,1);
	vector<int>vec2(10,3);
	vec.push_back(2);
	vec.insert(vec.begin(), vec2.begin(), vec2.end());
	vec.erase(vec.begin() + 10);//返回的是下一个迭代器 
	vector<int>::iterator iter; 
	for(auto item:vec)
	{
		cout<<" "<<item;
	}
	
	//list
	cout<<"test list."<<endl;
	list<int>ls;
	ls.push_back(1);
	ls.push_front(2);
	cout<<"ls.top() is: "<<ls.front()<<endl;
	ls.pop_back();
	ls.pop_front();
	ls.erase(ls.begin());
	for(auto item:ls)
	{
		cout<<item<<endl;	
	}	
	
	//deque
	cout<<"test deque."<<endl;
	deque<int>dq;
	dq.push_back(1);
	dq.push_front(2);
	dq.push_front(3);
	dq.push_front(4);
	dq.pop_back();
	dq.pop_front();
	cout<<"dq.front(): "<<dq.front()<<endl;
	cout<<"dq.back(): "<<dq.back()<<endl;
	dq.erase(dq.begin());
	deque<int>::iterator dp_iter;
	for(auto item:dq)
	{
		cout<<item<<endl;
	}
	
	//queue
	cout<<"test queue."<<endl;
	queue<int>q;
	q.push(1);
	q.push(2);
	q.pop();
	cout<<"q.front(): "<<q.front()<<endl;
	cout<<"q.back(): "<<q.back()<<endl;
	
	//stack
	cout<<"test stack."<<endl;
	stack<int> stk;
	stk.push(1);
	stk.push(2);
	cout<<"stack pop() :"<<stk.top()<<endl;;
	
	
	//priority_queue
	cout<<"test priority_queue."<<endl;
	priority_queue<int>pq;
	pq.push(1);
	pq.push(2);
	pq.push(3);
	pq.pop();
	cout<<"priority_queue pq.top(): "<<pq.top()<<endl;	
	
	
	/*
		顺序容器操作总结：
		vector  支持push_back() pop_back() front() back()  earse() iterator 
		list deque 支持push_back() push_front() pop_back() pop_front() earse() iterator 
		以下三种容器只支持对于这种数据结构而言应用的操作，不支持迭代器 
		stack     支持push()  pop()  top() <对于栈而言只支持访问栈顶的元素>
		queue	  支持push()  pop() front() back()<对于队列而言有前有后> 
		priority_queue	支持push() pop() top()<对于优先级队列而言只支持插入弹出和访问堆顶元素>
		
	*/
	
	return 0;	
} 
 
