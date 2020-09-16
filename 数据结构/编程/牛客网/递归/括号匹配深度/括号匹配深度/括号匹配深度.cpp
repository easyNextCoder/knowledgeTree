#include <iostream>
#include <vector>
#include <string>

using namespace std;

int count = 0;

bool isValid(int i, int j)
{
	return i >= 0 && i < 6 && j >= 0 && j < 6;
}

int get(int i, int j, vector<string>& con)
{
	if (i == 6 && j == 0)
	{
		cout << "her1" << endl;
		cout << count << endl;
		count++;
		return 0;
	}
	else if (j > 5) {
		cout << "here2" << endl;
		return get(i + 1, 0, con);
	}
	else {
		cout << "here3" << endl;
		if (con[i][j] == '*')
		{
			return get(i, j + 1, con);
		}
		else {
			vector<bool> visited(6, false);
			if (isValid(i - 1, j) && con[i - 1][j] != '*') {
				cout << i << j << endl;
				visited[con[i - 1][j] - '0'] = true;
			}
			if (isValid(i + 1, j) && con[i + 1][j] != '*') {
				visited[con[i + 1][j] - '0'] = true;
			}
			cout << "here4" << endl;
			if (isValid(i, j - 1) && con[i][j - 1] != '*') {
				visited[con[i][j - 1] - '0'] = true;
			}
			if (isValid(i, j + 1) && con[i][j + 1] != '*') {
				visited[con[i][j + 1] - '0'] = true;
			}

			for (int k = 0; k < 6; ++k)
			{
				if (!visited[k])
				{
					con[i][j] = k + '0';
					get(i, j + 1, con);
					con[i][j] = '#';
				}
			}
		}
	}


}

int main() {
	string tmp;
	vector<string> con;
	for (int i = 0; i < 6; i++)
	{
		cin >> tmp;
		con.push_back(tmp);
	}
	get(0, 0, con);
	cout << count << endl;

}

/*
##****
##****
******
******
******
******

*/

/*
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

string input;
int maxLeft = 0;

int tmpMax = 1;
vector<char> ccon;
int depth = 0;

//(())()(((())))
int u = 0;
int dfs(int n)
{

	int left = 0;
	//一种是用栈
	//一种用递归
	//一种用动态规划

	while (u < n)//为了去除右括号而向前移动
	{
		
			
		

		if (input[u] == '(')
		{
			if (depth != 0)
			{
				if (left != 0)
				{
					left = max(left, depth);
					depth = 0;
				}
				else {
					left = depth;
					depth = 0;
				}
			}
			u++;
			depth = 1 + dfs(n);
			u++;
		}
		else if (input[u] == ')')
		{
			if (left != 0)
				return max(left, depth);
			else
				return depth;
		}

	}
	return max(left, depth);
}

int main()
{

	cin >> input;
	cout << dfs(input.size()) << endl;;
	return 0;
}
*/