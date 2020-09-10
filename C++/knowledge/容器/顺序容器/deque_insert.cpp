#include <iostream>
#include <deque>

using namespace std;

int main()
{
	deque<int> a(2, 1);
	a.insert(a.begin(), 2);
	for(auto item:a)
		cout<<item<<endl;
	return 0;
} 
