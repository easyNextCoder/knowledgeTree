#include <iostream>
#include <string>

using namespace std;

int times = 0;
int minTimes = 1e8;
int target = 0;
string tmp;
string out;
bool dfs(int n)
{
	if (n == target)
	{
		out = tmp;
		minTimes = times;
		return true;

	}
	else if (n > target) {
		return false;
	}
	else {

		times++;
		tmp.push_back('2');
		bool rval1 = dfs(2 * n + 1);
		if (rval1 == true)return true;
		tmp.pop_back();
		times--;

		times++;
		tmp.push_back('3');
		bool rval = dfs(2 * n + 2);
		if (rval == true)return true;
		tmp.pop_back();
		times--;



	}
}

int main()
{
	int n;
	cin >> n;
	target = n;
	dfs(0);
	cout << out << endl;
	return 0;
}