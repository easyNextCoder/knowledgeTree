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
	
	cout<<sizeof(arr)<<endl;//һ�������Ƿ���bss�� 
	cout<<sizeof(arr2)<<endl;
	cout<<sizeof(*arr2)<<endl;
	cout<<sizeof(base_arr)<<endl;//���base���д���virtual����������Щvptrָ�붼�Ƿ���bss����� 
	cout<<sizeof(base_arr2)<<endl;
	cout<<sizeof(*base_arr2)<<endl;
	
	return 0;
}
