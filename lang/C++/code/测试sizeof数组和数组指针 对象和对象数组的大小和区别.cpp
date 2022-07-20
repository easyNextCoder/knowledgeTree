#include <iostream>

using namespace std;

class base{
	public:
		virtual get(){};
		//int basei;
};

int main(){
	
	int arr[10];
	int* arr2 = new int[10];
	base base_arr[10];
	base *base_arr2 = new base[10];
	
	cout<<sizeof(arr)<<endl;//一个整块是放在bss段 
	cout<<sizeof(arr2)<<endl;
	cout<<sizeof(*arr2)<<endl;
	cout<<sizeof(base_arr)<<endl;//如果base类中存在virtual函数，则这些vptr指针都是放在bss上面的 
	cout<<sizeof(base_arr2)<<endl;
	cout<<sizeof(*base_arr2)<<endl;
	
	return 0;
}
