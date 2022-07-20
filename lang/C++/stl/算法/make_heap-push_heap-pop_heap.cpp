#include <iostream>
#include <vector>
#include <algorithm>
#include <list>
#include <deque>
#include <memory>
#include <random>

using namespace std;
bool check(vector<bool>& max_value, vector<int>& tmp_mv)
{
	for (int i = 0; i < max_value.size(); ++i)
	{
		if (max_value[i] && tmp_mv[i] == 0)
		{
			return false;
		}
	}
	return true;
}
int main()
{
	
	default_random_engine e;
	vector<int> arr;
	const int N = 100;
	for (int i = 0; i < N; ++i)
		arr.emplace_back(e()%(N*10));
	make_heap(arr.begin(), arr.end());
	for (auto item : arr)
	{
		cout << item << " ";
	}
	cout << endl;

	arr.push_back(1200);
	push_heap(arr.begin(), arr.end());
	for (auto item : arr)
	{
		cout << item << " ";
	}
	cout << endl;

	pop_heap(arr.begin(), arr.end());
	for (auto item : arr)
	{
		cout << item << " ";
	}
	cout << endl;

	return 0;
}