#include <iostream>
#include <vector>

using namespace std;

typedef int biggestOne;

void percDown(vector<int>& heap, int father, int len)
{
	int dad = father;
	int son1 = dad*2+1;
	while(son1<len)
	{
		if(son1+1 < len && heap[son1+1] > heap[son1])
		{
			son1++;
		}
		if(heap[dad] > heap[son1])
		{
			return;
		}else{
			swap(heap[son1], heap[dad]);
			dad = son1;
			son1 = dad*2+1;
		}
	}
	
}

void make_heap(vector<int>& heap)
{
	for(int i = heap.size()/2; i>=0; i--)
	{//从最后一个父亲节点开始调整,建立一个最大堆 
		percDown(heap, i, heap.size());		
	}
	
	for(auto item:heap)
	{
		cout<<item<<" ";
	}
	cout<<endl;
	
	int first = 0;
	int len = heap.size();
	while(first < len-1)
	{
		cout<<heap[first]<<" "<<heap[len-1]<<endl;
		swap(heap[first], heap[len-1]);
		percDown(heap, first, --len);
	}
	
	
}

int main()
{
	vector<int> heap = {1,6,5,3,5,6,3};
	make_heap(heap); 
	for(auto item:heap)
	{
		cout<<item<<" ";
	}
	cout<<endl;
	return 0;
}
