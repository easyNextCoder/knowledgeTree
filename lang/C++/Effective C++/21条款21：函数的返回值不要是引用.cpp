#include <iostream>

using namespace std;

class base{
public:
	int geta(){
		return a;
	}
	~base(){
		cout<<"������base ������������"<<endl;
	}
private:
	int a = 0;
};

base& get_base(){
	return *(new base);
}

int main(){
	base& ref = get_base();
	cout<<ref.geta()<<endl;
	//���heap�ϵ��ڴ�й© 
	base& main_ref = *(new base);
	cout<<main_ref.geta()<<endl;
	//������Ȼ���heap�ϵ��ڴ�й© 
	base main_obj;
	cout<<main_obj.geta()<<endl;
	//ջ�ϵĿռ�����û��й©	 
	return 0;
} 
