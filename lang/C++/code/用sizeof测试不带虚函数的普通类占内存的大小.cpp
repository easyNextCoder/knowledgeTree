#include <iostream>

using namespace std;


class base{
	public:
		virtual get(){}; 
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
		char ac;//�漰���ڴ���룡 
		int bi;
};

void get(base a){
	cout<<"��������ʽ������ת����"<<endl;
};
int main(){
	base * b = new base();
	derived * db = new derived();
	get(*db);
	cout<<"�ڱ����ϣ�int������ռ�ö����ֽڣ�"<<sizeof(int)<<endl;
	cout<<"��ͨ�����麯������ռ���ڴ��С�ǣ�"<<sizeof(*b)<<endl;
	cout<<"��ͨ�����麯����������������ռ���ڴ��С�ǣ�"<<sizeof(*db)<<endl;;
	 
	delete b;//����������delete������ 
	delete db;//���ﲢû�м̳����ص�delete������ 
	return 0;
}
