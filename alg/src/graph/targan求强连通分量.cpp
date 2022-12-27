//https://www.bilibili.com/video/BV19J411J7AZ?p=5视频讲解

#include <iostream>
#include <vector>
#include <algorithm>
#include <list>
#include <deque>
#include <memory>
#include <random>
#include <stack>

using namespace std;


using namespace std;

#define M (INT_MAX)
#define PRINT_ARRAY(a,n)    do{for(int i = 0; i < n; i++) cout<<a[i]<<"|"; cout<<endl;}while(0)

/**********************************************
	1 → 0 → 3
	↑ ↙     ↓
	2         4
	3 → 4 ← 6 → 2
	↑↓  ↓ ↗ ↓ ↙↑
	7 → 5 → 0 → 1
**********************************************/
// #define V (5)
// int g[V][V] = 
// {
//     {0,0,1,1,0},
//     {1,0,0,0,0},
//     {0,1,0,0,0},
//     {0,0,0,0,1},
//     {0,0,0,0,0}
// };

#define V (8)
int g[V][V] =
{ // 0 1 2 3 4 5 6 7  
	{0,1,0,0,0,0,0,0},
	{0,0,1,0,0,0,0,0},
	{1,0,0,0,0,0,0,0},
	{0,0,0,0,1,0,0,1},
	{0,0,0,0,0,1,0,0},
	{1,0,0,0,0,0,1,0},
	{1,0,1,0,1,0,0,0},
	{0,0,0,1,0,1,0,0}
};

/**********************************************
	强连通分量 strongly connected component
**********************************************/

void tarjan_dfs(int x, int dfn[], int low[], stack<int>& s, bool in_stack[])
{
	static int time = 1;
	dfn[x] = low[x] = time++;
	s.push(x);
	in_stack[x] = true;

	for (int y = 0; y < V; y++)
	{
		if (g[x][y])
		{
			if(0 == dfn[y])
			{
				tarjan_dfs(y, dfn, low, s, in_stack);
				low[x] = min(low[x], low[y]);
			}
			else if (in_stack[y])
				low[x] = min(low[x], dfn[y]);
		}
	}

	if (dfn[x] == low[x])
	{
		int tmp;
		do
		{
			tmp = s.top(); s.pop();
			in_stack[tmp] = false;
			cout << tmp << "-";
		} while (tmp != x);
		cout << endl;
	}
}

void scc_tarjan()
{
	int dfn[V] = { 0 }, low[V] = { 0 };
	bool in_stack[V] = { false };
	stack<int> s;
	for (int i = 0; i < V; i++)
		if (!dfn[i])
			tarjan_dfs(i, dfn, low, s, in_stack);
	/*
	targan_dfs(i, dfn, low, s, in_stack)
	{
		s.push_back(i);
		dfn[i] = low[i] = time++;
		in_stack[i] = true;
		for (int j = 0; j < V; ++j)
		{
			if (G[i][j])
			{
				if (dfn[j])
				{
					targan(j, dfn, low, s, in_stack);
					low[i] = min(low[i], low[j]);
				}
				else if (in_stack[y]) {
					low[] = min(low[x], dfn[y]);
				}
				
			}
			
		}
		if (low[x] == low[y])
		{

		}
	}
	*/
}

void targan(int low[], int dfn[], int in_stack[], int index, stack<int> mstack)
{
	static int time = 1;
	low[index] = dfn[index] = time++;
	mstack.push(index);
	for (int i = 0; i < V; ++i)
	{
		if (g[index][i])
		{
			if (!dfn[i])
			{
				targan(low, dfn, in_stack, i, mstack);
				low[index] = min(low[index], low[index]);
			}
			else if(in_stack[i]){
				low[index] = min(low[index], dfn[i]);
			}
		}
	}

	if (low[index] == dfn[index])
	{
		int tmp;
		do {
			tmp = mstack.top(); mstack.pop();
			in_stack[tmp] = false;
			cout << tmp<<" ";
		} while (tmp != index);
	}
	cout << endl;
}

void xyk_targan()
{
	
	int low[V] = { 0 };
	int dfn[V] = { 0 };
	int in_stack[V] = { 0 };
	stack<int> mstack;
	for (int i = 0; i < V; ++i)
	{
		if (!dfn[i])
			targan(low, dfn, in_stack, i, mstack);
	}
	
}

int main()
{
	scc_tarjan();
	xyk_targan();
	return 0;
}