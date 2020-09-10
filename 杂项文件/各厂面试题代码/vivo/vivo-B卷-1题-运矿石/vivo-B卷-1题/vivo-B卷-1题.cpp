#define _CRT_SECURE_NO_WARNINGS
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <algorithm>
#include <string>

using namespace std;

/**

* Welcome to vivo
  */

#define MAX_NUM 101
  /*
  int f[101][2];
  int sum[101];
  int rmin = 1e8;
  int dfs(int n, int weight[], int start, int tmpSum)
  {
	  for(int i = start; i<n; i++)
	  {
		 int ntmpSum = tmpSum+weight[i];
		 int ntmpSum1 = tmpSum;
		 if(abs(rmin*2-sum[n]) > abs(ntmpSum*2-sum[n]))
		 {
			 rmin = ntmpSum;
		 }
		 dfs(n, weight, start+1, ntmpSum);
		 if(abs(rmin*2-sum[n]) > abs(ntmpSum1*2-sum[n]))
		 {
			 rmin = ntmpSum1;
		 }
		 dfs(n, weight, start+1, ntmpSum1);
	  }
	  return 0;
  }
  int solution(int n, int weight[]) {

	  // TODO Write your code here

	  for(int i = 0; i<n; i++)
	  {
		  sum[i+1]=sum[i] + weight[i];
		  //bug:初次联系容易把sum求错
	  }
	  //dfs(n, weight, 0, 0);
	  //bug：当数据量达到20个的时候，超时或者爆栈

	  //f[i][j]  表示第i块石头达到质量j

	  for(int i = 0; i<=n; ++i)
	  {
		  for(int j = weight[]; j<=sum[n])
	  }


	  return abs(rmin-(sum[n]-rmin));

  }
  */
int dp[10001][101] = { 0 };
int solution(int n, int weight[]) {

	// TODO Write your code here
	int s = 0;
	for (int i = 0; i < n; ++i)
	{
		s += weight[i];
	}


	int mi = 1e6;
	int ans = 0;
	dp[0][0] = 1;
	//dp[可以组成的石头重量][石头个数]
	for (int i = 0; i < n; ++i)
	{//遍历挑选重量为weight[i]的石头，加入到计算队列中

		for (int j = s; j >= weight[i]; --j)//由下面的if(dp)等式决定
		{//当前可以达到的石头的重量
			for (int k = 0; k <=n; ++k)
			{//遍历石头的个数
				
				if (dp[j-weight[i]][k])
				{
					dp[j][k + 1] = 1;//由上一个存在推出这个等式存在
				

						if (abs(j - s/2) < mi) {
							mi = abs(j - s/2);
							ans = j;
						}
				}
			}
		}
	}

	return s-2*ans;
}

int main()
{

	string str("");
	getline(cin, str);
	int a[MAX_NUM];
	int i = 0;
	char* p;
	int count = 0;


	const char* strs = str.c_str();
	p = strtok((char*)strs, " ");
	while (p)
	{
		a[i] = atoi(p);
		count++;
		p = strtok(NULL, " ");
		i++;
		if (i >= MAX_NUM)
			break;
	}

	int result = solution(count, a);
	cout << result << endl;
	return 0;

}


