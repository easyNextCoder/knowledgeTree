#include <iostream>

using namespace std;


class base{
	public:
		void operator delete(void*s){
			cout<<"������delete�������"<<endl;
			
			return ;
		}
		
};


class derived{
	public:
		operator base(){
			return obj;
		}
	private:
		base obj;
};

void get(base a){
	cout<<"��������ʽ������ת����"<<endl;
};
int main(){
	base * b = new base();
	derived dobj;
	get(dobj);
	
	delete b;
	return 0;
}
