#include <iostream>
#include <vector>
#include <map>
#include <unordered_map>
#include <string>
#include <set>
#include <algorithm>
#include <queue>

using namespace std;

int main()
{
	int N;
	cin >> N;
	vector<int> f(N, 0);
	for (int i = 0; i < N; i++)
		cin >> f[i];
	priority_queue<int, vector<int>, less<int>> q;
	for (auto item : f)
		q.push(item);
	int sum = 0;
	while (q.size() >= 2)
	{
		int tmp = q.top();
		q.pop();
		int newtmp = q.top();
		while (tmp == newtmp && q.size() > 0)
		{
			q.pop();
		}
		sum += tmp - newtmp;
	}
	if (!q.empty())sum += q.top();
	cout << sum << endl;
	/*
	sort(f.begin(), f.end());
	int sum = 0;
	for(int i = 0; i<f.size()-1; i++)
		sum += f.back()-f[i];
	cout<<sum<<endl;
	*/

	return 0;
}

