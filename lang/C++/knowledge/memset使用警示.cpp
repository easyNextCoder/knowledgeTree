#include <iostream>
#include <string.h>

using namespace std;
#define LEN 10
int main()
{
	int **p = new int*[LEN];
	for(int i = 0; i<LEN; i++)
	{
		p[i] = new int[LEN];
		memset(p[i], 0, LEN*sizeof(int));
		//�Լ���ǰʹ��memset�ͳ��ֹ������Լ�ʹ�õ���memset(p[i], 0, LEN);
		//������һ������������ĵ� 
	}
	
	
	for(int i = 0; i<LEN; i++)
	{
		for(int j = 0; j<LEN; j++)
		{
			cout<<p[i][j]<<" ";
		}
		cout<<endl;
	}
		
	return 0;
}
