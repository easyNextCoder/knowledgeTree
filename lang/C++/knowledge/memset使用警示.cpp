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
		//自己以前使用memset就出现过问题自己使用的是memset(p[i], 0, LEN);
		//上面是一个经常犯错误的点 
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
