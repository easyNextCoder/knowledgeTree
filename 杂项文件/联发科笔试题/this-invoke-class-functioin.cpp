#include <iostream>

using namespace std;

class A{
	public:
		int a;
		void print(){
			cout<<"hello"<<endl;
		}
		void invoke_print()
		{
			this->print();
		}
};

int main()
{
	A a;
	a.invoke_print();	
	return 0;
}
