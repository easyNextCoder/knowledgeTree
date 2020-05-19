#include <iostream>
#include <vector>

using namespace std;

void persDown(vector<int>& to_sort, int dad)
{
	int size = to_sort.size();
	int son = dad*2 + 1;
	//cout<<to_sort.size()<<endl;
	while(son <= size-1){
		
		//cout<<son<<endl;
		if(son + 1 < size)
		{
			if(to_sort[son] > to_sort[son+1])
				//swap(to_sort[son], to_sort[son+1]);//²»ÒªËæ±ãswap 
				son = son+1;
		}
		if(to_sort[son] < to_sort[dad]){
			swap(to_sort[son], to_sort[dad]);
		}
		dad = son;
		son = dad*2 + 1;
		//cout<<dad<<endl;
	}
	
	
}

int main()
{
	vector<int>to_sort = {9,8,7,6,5,4,3,2,1};
	int N = to_sort.size();
	for(int i = N/2-1; i>=0; i--)
	{
		persDown(to_sort, i);
	}
	for(auto item:to_sort)
	{
		cout<<item<<endl;
	}
	return 0;
}
