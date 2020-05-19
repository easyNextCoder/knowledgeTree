#include <iostream>

using namespace std;
#define N 4

void rotate(int matrix[][N])
{
	for(int i = 0; i<N/2; i++)
	{
		for(int j = i; j<N-i-1; j++)
		{
			/*//这个是逆时针旋转 
			int tmp = matrix[i][j];
			matrix[i][j] = matrix[j][N-i-1];
			matrix[j][N-i-1] = matrix[N-i-1][N-j-1];
			matrix[N-i-1][N-j-1] = matrix[N-j-1][i];
			matrix[N-j-1][i] = tmp;
			*/
			int tmp = matrix[j][N-i-1];
			matrix[j][N-i-1] = matrix[i][j];
			matrix[i][j] = matrix[N-j-1][i];
			matrix[N-j-1][i] = matrix[N-i-1][N-j-1];
			matrix[N-i-1][N-j-1] = tmp;
		}
	}
}
 
int main()
{ 
	int matrix[][N] = 
	{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16}
	};
	rotate(matrix);
	rotate(matrix);
	for(int i = 0; i<N; i++)
	{
		for(int j = 0; j<N; j++)
		{
			cout<<matrix[i][j]<<" ";
		}
		cout<<endl;
	}
	
	return 0;
}
