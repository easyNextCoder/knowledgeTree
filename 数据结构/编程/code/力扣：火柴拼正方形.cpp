
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:

	bool helper(vector<int>& nums, int rec[], int aimLength, int tmpLength, int tmpIndex, int borderHaveMadeNum)
	{
		/*
			if (!(tmpIndex < nums.size()))
				return false;
			需要把这一段屏蔽：原因是，tmpLength的增长是慢tmpIndex一步的
		*/
		


		if (tmpLength < aimLength && tmpIndex < nums.size())
		{
			if (!rec[tmpIndex])
			{
				//加上当前index的值
				rec[tmpIndex] = 1;
				bool rval1 = helper(nums, rec, aimLength, tmpLength + nums[tmpIndex], tmpIndex + 1, borderHaveMadeNum);
				if (rval1)return rval1;
				rec[tmpIndex] = 0;
				//不加上当前index的值，继续往后找
				bool rval2 = helper(nums, rec, aimLength, tmpLength, tmpIndex + 1, borderHaveMadeNum);
				if (rval2)return rval2;
			}
			else {
				bool rval2 = helper(nums, rec, aimLength, tmpLength, tmpIndex + 1, borderHaveMadeNum);
				if (rval2)return rval2;
			}

		}
		else if (tmpLength == aimLength) {
			if (borderHaveMadeNum == 3)
				return true;
			else {
				bool rval2 = helper(nums, rec, aimLength, 0, 0, borderHaveMadeNum + 1);
				if (rval2)return rval2;
			}
		}
		else {
			return false;
		}
	}

	bool makesquare(vector<int>& nums) {
		int total = 0;

		for (auto item : nums)
			total += item;
		//cout << total << endl;
		if (total % 4)return false;
		int borderLength = total / 4;
		//cout << "borderLength is: "<<borderLength << endl;

		int* recordUsed = new int[nums.size()];
		
		for (int i = 0; i < nums.size(); i++)
		{
			recordUsed[i] = 0;
		}
		
		return helper(nums, recordUsed, borderLength, 0, 0, 0);

	}
};

int main()
{
	Solution solution;
	vector<int> vec = { 3, 3, 3, 3, 1,1,1,0,1,2,1,1,2,2 ,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1};
	cout << solution.makesquare(vec);
	return 0;
}
