#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <vector>
#include <algorithm>

using namespace std;

/**
 * Welcome to vivo !
 */

#define MAX_NUM 100

void dfs(int boxs[], int N, vector<int> & result, int i, int j, int tmp_res, int last_value)
{
	if (!(i <= j)) {
		result.push_back(tmp_res);
	}
	else {
		int v = 0;
		while (i <= j)
		{
			if (boxs[i] == last_value)
			{
				v++;
				i++;
			}
			else if (boxs[j] == last_value)
			{
				v++;
				j--;
			}
			else {

				dfs(boxs, N, result, i, j, tmp_res + v * v, boxs[i]);
				dfs(boxs, N, result, i, j, tmp_res+v*v, boxs[j]);
			}
		}
		result.push_back(tmp_res + v * v);
	}
}
int solution(int boxs[], int N)
{
	// TODO Write your code here
	vector<int> result;
	int tmp_res = 0;
	dfs(boxs, N, result, 0, N - 1, tmp_res, boxs[0]);
	dfs(boxs, N, result, 0, N-1, tmp_res, boxs[N-1]);
	sort(result.begin(), result.end());
	return result.back();
}

int main()
{
	string str("");
	getline(cin, str);
	int boxs[MAX_NUM];
	int i = 0;
	char* p;
	int count = 0;

	const char* strs = str.c_str();
	p = strtok((char*)strs, " ");
	while (p)
	{
		boxs[i] = atoi(p);
		count++;
		p = strtok(NULL, " ");
		i++;
		if (i >= MAX_NUM)
			break;
	}

	int num = solution(boxs, count);
	cout << num << endl;
	return 0;
}