#include <iostream>
#include <vector>

using namespace std;

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