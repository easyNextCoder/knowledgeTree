#include <iostream>
#include <sstream>
#include <time.h>
using namespace std;

int main()
{
	int stime_start = clock();
	int stime_end = 0;
	for(int i = 0; i<1000000; i++)
	{
		string a = "123";
		atoi(a.c_str());
	}
	stime_end = clock();
	cout<<stime_end-stime_start<<endl; 
	
	int out = 0;
	int itime_start = clock();
	int itime_end = clock();
	string b= "123";
	for(int i = 0; i<1000000; i++)
	{
		stringstream ss(b);
		ss>>out;
	}
	itime_end = clock();
	cout<<itime_end-itime_start<<endl;
	return 0;
	//���Խ�����������stringstream����Ĺ���ʱ�䣬atoiԱʤ��stringstream. 
}
