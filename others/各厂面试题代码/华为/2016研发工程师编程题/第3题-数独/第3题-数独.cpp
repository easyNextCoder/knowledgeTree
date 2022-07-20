
#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <set>

int mg[9][9];
using namespace std;
int shouldFiguredNum = 0;
int handle_mg(int mg[][9], int startRow, int startCol, int figuredNum)
{
	cout << figuredNum << ' ' << shouldFiguredNum << endl;;
	if (figuredNum == shouldFiguredNum)
	{
		return 1;
	}
	for (int i = startRow; i < 9; i++)
	{
		for (int j = startCol; j < 9; j++)
		{
			if (mg[i][j] == 0)
			{//找到可以填入的点
				int ts[9] = {0};
				for (int k = 0; k < 9; k++)
				{
					if(mg[i][k])
						ts[mg[i][k]-1] = 1;
					if(mg[k][j])
						ts[mg[k][j]-1] = 1;
				}
				
				for (int q = 0; q < 9; q++)
				{
					if (!ts[q])
					{
						int rval = 0;
						mg[i][j] = q + 1;
						rval = handle_mg(mg, i, j+1, figuredNum + 1);
						if (rval)return 1;
						mg[i][j] = 0;
					}
				}
				//找到一个点然后开枝散叶，然后返回
				return 0;
			}
		}
		startCol = 0;
	}
	return 0;
}
int main()
{
	string line;
	int count = 0;
	while (getline(cin, line))
	{

		int tmp;
		int countCol = 0;
		stringstream sstr(line);
		while (sstr >> tmp)
		{
			mg[count][countCol++] = tmp;
		}
		count++;
		if (count == 9)
		{

			count = 0;
			//handle
			for (int i = 0; i < 9; i++)
			{
				for (int j = 0; j < 9; j++)
				{
					if (mg[i][j] == 0)
					{//找到可以填入的点
						shouldFiguredNum++;
					}
				}
			}
			cout << "here." <<shouldFiguredNum<< endl;
			cout<<handle_mg(mg, 0, 0, 0);
			for (int i = 0; i < 9; i++)
			{
				for (int j = 0; j < 9; j++)
				{
					cout << mg[i][j] << " ";
				}
				cout << endl;
			}
		}
	}

	return 0;
}






