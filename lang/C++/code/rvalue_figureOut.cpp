#include <iostream>
using namespace std;

class RightRef{
public:
	RightRef():b(new int()){
		cout<<"constructor"<<endl;
	}
	RightRef(RightRef& input){
		//b = in.b;
		cout<<"RightRef(&)"<<endl;
	}
	RightRef(RightRef&& input){
		//in.b = nullptr;
		cout<<"invoke RightRef(&&)"<<endl;
	}
private:
	int * b;
};
RightRef getRightRef(){
	RightRef a;
	return a;
}
int main(){
	Right a = getRightRef();
	//����������У�����Ӧ�õ����ƶ����캯�������������ֵ���ı������������Ż���ֱ�ӰѶ��еĶ������Ŀ�ı��� 
	return 0;
}
