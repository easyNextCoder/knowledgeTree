#include <iostream>
#include <exception>
#include <vector>
using namespace std;

double division(int a, int b)
{
	if(b == 0)
	{
		throw "error: divide by zero.";
	}
	return a/b;
}

class MyException:public exception
{
public:
	const char* what() const throw()
	{
		return "C++ exceptioin test.";	
	}	
	vector<string> MyWhat() const throw(vector<string>)
	{
		return vector<string>({"line: 234", "wrong aixs"});
	}
}; 

int main()
{
	//���� try-catch���1 
	try
	{
		division(1, 0);	
	}
	catch(char const * p)//char const *p = const char *p
	{
		cout<<p<<endl;
		
		/*һ����������ָ��ײ��ָ���ָ�򶥲��ָ��*/
		cout<<"+--------------------------------------------+"<<endl;
		char s[20] = {"name"};
		char* const q = s;
		cout<<q<<endl;
		cout<<"+--------------------------------------------+"<<endl;
		
	}
	//����try-catch���2 
	try
	{
		try
		{
			throw MyException();
		}
		catch(MyException me)
		{
			cout<<me.what()<<endl;
			vector<string> out_info = me.MyWhat();
			for(auto item:out_info)
			{
				cout<<item<<endl;
			}
			throw;//�����������׳�����һ���� 
		}
		catch(exception e)
		{
			cout<<e.what()<<endl;//��������Ӧ�÷ŵ������� 
		}
	}
	catch(...)//�������е��쳣���� 
	{
		try{}
		catch(MyException me)
		{
			cout<<"we catched the rethrowed exception and handled."<<endl;
		}
		cout<<"we catched the rethrowed exception but didn't handle."<<endl;
	}
	
	return 0;
}
