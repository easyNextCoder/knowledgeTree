#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;
bool isr(string& a, int left, int right)
{
	if (left >= right)return true;
	int len = right - left + 1;
	for (int i = 0; i < len / 2; ++i)
	{
		if (a[left + i] != a[right + i])
			return false;
	}
	return true;
}
bool is(string& a, int left1, int right1, string& b, int left2, int right2)
{
	int len = min(right1 - left1 + 1, right2 - left2 + 1);
	int i = 0;
	while (i < len)
	{
		if (a[left1 + i] != b[left2 + i])
			return false;
		i++;
	}
	return true;
}
const int M = 15;

vector<vector<char>> f = { {0, 1}, {0, -1}, {1, 0}, {-1, 0},{1,1},{-1,-1},{-1,1},{1,-1} };
vector<vector<int>> memo(M, vector<int>(M, -1));
vector<vector<int>> visited(M, vector<int>(M, 0));

int get(int x, int y, int m, int n, vector<vector<char>>& con) {
	//cout << x << ":" << y << " ";
	if (x == m && n == y)
	{
		cout << "we finally get it." << endl;
		return 1;
	}
	else if (!(x >= 0 && x < con.size() && y >= 0 && y < con.size()))
	{
		return 1e9;
	}
	else if (con[x][y] == '#' || con[x][y] == '@') {
		return 1e9;
	}
	else if (visited[x][y])
	{
		return 1e9;
	}
	//cout << x << ":" << y << memo[x][y] << endl;
	if (memo[x][y] != -1)
	{
		return memo[x][y];
	}

	int min_out = 1e8;
	visited[x][y] = 1;
	for (auto& item : f) {
		min_out = min(min_out, 1 + get(x + item[0], y + item[1], m, n, con));
		
	}
	memo[x][y] = min_out;

	visited[x][y] = 0;
	if (min_out < 1e8)
		cout << x << ":" << y << ":" << min_out << endl;
	return min_out;
}
int main()
{
	int N;
	cin >> N;
	int x, y, m, n;
	cin >> x >> y >> m >> n;
	vector<vector<char>> con(N, vector<char>(N, 0));
	cout << N << endl;
	for (int i = 0; i < N; ++i)
	{
		for (int j = 0; j < N; ++j)
		{
			char a;

			cin >> a;
			con[i][j] = a;
		}
	}


	int ret = get(m, n, x, y, con);
	cout << endl;
	cout << "here is the result:" << endl;
	if (ret > 1e6)
		cout << -1 << endl;
	else
		cout << ret << endl;
	return 0;
}

/*

15
0 7 7 7
*5#++B+B+++++$3
55#+++++++###$$
###$++++++#+*#+
++$@$+++$$$3+#+
+++$$+++$+4###+
A++++###$@+$++A
+++++#++$#$$+++
A++++#+5+#+++++
+++$$#$++#++++A
+++$+@$###+++++
+###4+$+++$$+++
+#+3$$$+++$##++
+#*+#++++++#$$+
$####+++++++$##
3$+++B++B++++#5


13
*/