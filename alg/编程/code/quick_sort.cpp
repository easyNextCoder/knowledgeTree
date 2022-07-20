#include <iostream>
#include <vector>

using namespace std;

int get_provit(vector<int>& vec, int first, int last)
{

	int mid = first + (last - first) / 2;
	if (vec[first] > vec[last])
		swap(vec[first], vec[last]);
	if (vec[mid] > vec[last])
		swap(vec[last], vec[mid]);
	if (vec[mid] < vec[first])
		swap(vec[mid], vec[first]);
	return vec[mid];
}
void quick_sort(vector<int>& vec, int first, int last)
{
	cout << first << "-" << last << endl;
	if (last - first <= 1)
	{
		if (vec[first] > vec[last])
			swap(vec[first], vec[last]);
		return;
	}
	int provit = get_provit(vec, first, last);
	swap(vec[first + (last - first) / 2], vec[last - 1]);
	int i = first;
	int j = last - 1;
	for (;;)
	{
		while (vec[++i] < provit) { ; }//竟然把vec都忘了
		while (vec[--j] > provit) { ; }
		if (i < j)
		{
			swap(vec[i], vec[j]);
		}
		else {
			break;
		}
	}
	swap(vec[i], vec[last - 1]);
	quick_sort(vec, first, i - 1);
	quick_sort(vec, i + 1, last);

	return;
}

int main()
{
	vector<int> vec = { 9,8,7 };
	quick_sort(vec, 0, vec.size() - 1);
	for (auto item : vec)
	{
		cout << item << " ";
	}
	cout << endl;
}