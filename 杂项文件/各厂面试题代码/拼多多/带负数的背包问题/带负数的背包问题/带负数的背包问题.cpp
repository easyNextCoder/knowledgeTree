#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>

using namespace std;


class Solution2 {
public:
	/**
	 * 计算t的最小长度
	 * @param str string字符串 输入的字符串
	 * @return int整型
	 */
	int check(string& s1, int i, string& s2, int j, int len)
	{
		int tmpLen = 0;
		len = min((int)s2.size() - i, len);
		while (tmpLen < len)
		{
			if (s1[i + tmpLen] == s2[j + tmpLen])
				tmpLen++;
			else return tmpLen;
		}
		return tmpLen;
	}
	int getMinLen(string str) {
		// write code here
		int slen = str.size();
		int min_value = slen;
		for (int len = 1; len < slen; ++len)
		{

			string splitStr = str.substr(0, len);
			int start = len;
			int finish = slen;
			int equal_len = 0;
			while ( ( equal_len = check(splitStr, 0, str, start, len) ) == len)
			{
				//cout << equal_len << " ";
				start += len;
			}
			
			if (start + equal_len == slen )
			{
				min_value = min(min_value, len - equal_len);
			}
			if (start == slen)
				min_value = 0;
		}
		
		return min_value;
	}
};

int main()
{
	Solution2 so;
	string s = "abcabcac";
	cout <<"the result is:"<< so.getMinLen(s) << endl;;
}


class Solution {
public:
	/**
	 * 获取最小得分
	 * @param gz int整型一维数组 瓜子堆的组成
	 * @param gzLen int gz数组长度
	 * @return int整型
	 */
	int getMinScore(int* gz, int gzLen) {
		// write code here

		vector<vector<int>> f(gzLen + 1, vector<int>(gzLen + 1, 1e8));
		vector<int> sum(gzLen + 1, 0);
		for (int i = 1; i <= gzLen; ++i)
		{
			f[i][i] = 0;
		}

		for (int i = 1; i <= gzLen; ++i)
		{
			sum[i] = sum[i - 1] + gz[i - 1];
		}

		for (int len = 2; len <= gzLen; ++len)
		{
			for (int i = 1; i + len - 1 <= gzLen; ++i)
			{
				int j = i + len - 1;
				for (int k = i ; k < j; ++k)
				{
					f[i][j] = min(f[i][j], f[i][k] + f[k+1][j] + sum[j] - sum[i - 1]);
				}
			}
		}
		//int sum = accumulate(gz, gz + gzLen, 0);
		return f[1][gzLen] ;
	}
};
/*
int main()
{
	Solution so;
	const int N = 4;
	int gz[N] = {1,3,5,2};
	cout << so.getMinScore(gz, N) << endl;
	return 0;
}

*/

/*
void dfs(int i, int j, int color, vector<vector<int>>& f) {
	int len = f.size();
	if (!(i >= 0 && i < len && j >= 0 && j < len))return;
	if (f[i][j] == 0)return;
	if (f[i][j] != color)
	{
		f[i][j] = color;
		dfs(i + 1, j, color, f);
		dfs(i - 1, j, color, f);
		dfs(i, j + 1, color, f);
		dfs(i, j - 1, color, f);

	}
	
}

int main()
{
	int n, old_n;
	cin >> old_n;
	n = old_n;
	if (old_n % 2 == 0)n = old_n+1;
	vector<vector<int>> f(n, vector<int>(n, 9));
	for (int i = 0; i < n; ++i)
	{
		f[i][i] = 0;
			f[n / 2][i] = 0;
	}

	for (int i = 0; i < n; ++i)
	{
		f[n - i - 1][i] = 0;
			f[i][n / 2] = 0;
	}


	int count = 1;
	for (int i = n - 1; i >= 0; --i) {
		if (f[0][i] != 0)
		{
			dfs(0, i, count++, f);
			break;
		}
	}
	
	for (int i = 0; i < n; ++i) {
		if (f[0][i] != 0) {
			dfs(0, i, count++, f);
			break;
		}
	};
	for (int i = 0; i < n; ++i) {
		if (f[i][0] != 0) {
			dfs(i, 0, count++, f);
			break;
		}
	};
	for (int i = n - 1; i >= 0; --i) {
		if (f[i][0] != 0) {
			dfs(i, 0, count++, f);
			break;
		}
	};

	for (int i = 0; i < n; ++i) {
		if (f[n - 1][i] != 0) {
			dfs(n - 1, i, count++, f);
			break;
		}
	};
	for (int i = n - 1; i >= 0; --i) {
		if (f[n - 1][i] != 0) {
			dfs(n - 1, i, count++, f);
			break;
		}
	};
	for (int i = n - 1; i >= 0; --i) {
		if (f[i][n - 1] != 0) {
			dfs(i, n - 1, count++, f);
			break;
		}
	};
	for (int i = 0; i < n; ++i) {
		if (f[i][n - 1] != 0) {
			dfs(i, n - 1, count++, f);
			break;
		}
	};
	
	for (int i = 0; i < f.size(); ++i)
	{
		if (old_n % 2 == 0 && i == (int)f.size() / 2)
			continue;
		for (int j = 0; j < f.size(); ++j)
		{
			if (old_n % 2 == 0 && j == (int)f.size()/2)
			{
				continue;
			}
			cout << f[i][j] << " ";
		}
	}
	
	return 0;
}
*/