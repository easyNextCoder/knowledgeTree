// 全局倒置与局部倒置.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int searchFirstLess(int target, vector<int>& dp, int first, int last)
    {
        int left = first;
        int right = last;
        int mid = first + (last - first+1) / 2;
        //而分搜索只适合找不能重复的元素吗？
        cout << target << endl;
        while (left <= right)
        {
            cout << left << ":" << right << endl;
            mid = left + (right - left+1) / 2;//在这个里面+1代表，当还有两个元素的时候，mid始终是第二个元素，不加1则始终是第一个元素
            if (dp[mid] > target)
            {
                left = mid;
            }
            else if (dp[mid] < target)
            {
                right = mid;
            }
            else {
                ;
            }
        }
        return left + 1;
    }
    bool isIdealPermutation(vector<int>& A) {
        int globalR = 0;
        int localR = 0;
        vector<int>dp(A.size());
        dp[A.size() - 1] = A[A.size() - 1];//dp保存着当前及当前位置之后的最大值
        int tmp_max = dp.back();
        for (int i = A.size() - 1; i >= 0; i--)
        {
            if (tmp_max > A[i])
            {
                dp[i] = tmp_max;
            }
            else {
                tmp_max = A[i];
                dp[i] = tmp_max;
            }
            cout << dp[i] << "-";
        }
        cout << endl;

        for (int i = 0; i < A.size() - 1; i++)
        {
            int index = searchFirstLess(A[i], dp, i + 1, A.size()-1);
            globalR += A.size() - index + 1;
        }

        for (int j = 0; j < A.size() - 1; j++)
        {
            if (A[j] > A[j + 1])
                localR++;
        }

        return globalR == localR;

    }
};


int main()
{
    Solution solution;
    vector<int> vec = { 1,0,2 };
    solution.isIdealPermutation(vec);
    std::cout << "Hello World!\n";
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
