// 力扣：交点.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
	vector<double> intersection(vector<int>& start1, vector<int>& end1, vector<int>& start2, vector<int>& end2) {
		double x1 = start1[0];
		double y1 = start1[1];
		double x2 = end1[0];
		double y2 = end1[1];
		double k1 = 0;
		double b1 = 0;

		double m1 = start2[0];
		double n1 = start2[1];
		double m2 = end2[0];
		double n2 = end2[1];
		double k2 = 0;
		double b2 = 0;

		typedef pair<double, double> Point;
		if (x1 == x2 && m1 == m2)
		{
			if (x1 == m1) {
				
				vector<double>sorty = { y1,y2,n1,n2 };
				sort(sorty.begin(), sorty.end());
				//判断线段是否是重合的
				if (abs(y2-y1) + abs(n2-n1) >= sorty[3] - sorty[0])//bug=号产生首尾相接的线段
					return { x1,sorty[1] };
				else
					return {};
			}
			else {
				return {};
			}
		}
		else if (x1 == x2 || m1 == m2)
		{
			if (x1 == x2)
			{
				k2 = (n2 - n1) / (m2 - m1);;
				b2 = n1 - k2 * m1;
				double outx = x1;
				double outy = k2 * x1 + b2;
				if (outy >= min(y1, y2) && outy <= max(y1, y2) && outy >= min(n1, n2) && outy <= max(n1, n2)
					&& outx <= max(m1, m2) && outx >= min(m1, m2))
				{
					return { outx, outy };
				}
				else {
					return {};
				}
			}
			else {
				k1 = (y2 - y1) / (x2 - x1);
				b1 = x1 - k1 * y1;
				double outx = m1;
				double outy = k1 * m1 + b1;
				if (outy >= min(y1, y2) && outy <= max(y1, y2) && outy >= min(n1, n2) && outy <= max(n1, n2)
					&& outx <= max(x1, x2) && outx >= min(x1, x2))
				{
					return { outx, outy };
				}
				else {
					return {};
				}
			}
		}
		else {
			k1 = (y2 - y1) / (x2 - x1);
			b1 = y1 - k1 * x1;//bug:x与y写反了
			k2 = (n2 - n1) / (m2 - m1);
			b2 = n1 - k2 * m1;
		}

		double kdiff = (k2 - k1);

		if (abs(kdiff) < pow(10, -6))//bug1没设想到kdiff是负值
		{
			double bdiff = (b2 - b1);
			if (abs(bdiff) < pow(10, -6))
			{
				//两条直线重合,而且不存在垂直方向的重合
				vector<Point> sorta = { {x1,y1}, {x2,y2},{m1, n1}, {m2,n2} };
				sort(sorta.begin(), sorta.end(), [](auto& a, auto& b) {return a.first < b.first; });

				if (abs(x2 - x1) + abs(m2 - m1) >= ((max(max(x1, x2), max(m1, m2))) - (min(min(x1, x2), min(m1, m2)))))
				{
					//证明重合且有交点
					return { sorta[1].first, sorta[1].second };
				}
				else {
					return {};
				}
			}
			else {
				//两条直线平行
				return {};
			}
		}
		else {
			double outx = 0;
			double outy = 0;
			if (k1 == 0)
			{//bug2其中有一条是水平线，就使用另一条计算

				outy = y1;
				outx = (outy - b2) / k2;
			}
			else if (k2 == 0) {

				outy = n1;
				outx = (outy - b1) / k1;
			}
			else {
				outy = (b1 * k2 - b2 * k1) / (k2-k1);
				outx = (outy-b1)/k1;
			}
			//bug:60>=60这个可能会输出false
			double y_min = min(y1, y2);
			double y_max = max(y1, y2);
			double n_min = min(n1, n2);
			double n_max = max(n1, n2);
			double x_min = min(x1, x2);
			double x_max = max(x1, x2);
			double m_min = min(m1, m2);
			double m_max = max(m1, m2);
			if (outy - y_min >= pow(10, -6) && y_max-outy>=pow(10, -6) && outy - n_min >= pow(10, -6) && n_max - outy >= pow(10, -6)
				&& outy - x_min >= pow(10, -6) && x_max - outy >= pow(10, -6) && outy - m_min >= pow(10, -6) && m_max - outy >= pow(10, -6))
			{
				return { outx, outy };
			}
			else {
				return {};
			}
		}

	}
};

int main()
{
	Solution solution;
	vector<int> start1, end1, start2, end2;
	/*
	start1 = { 0,0 };
	end1 = { 3,3 };
	start2 = { 1,1 };
	end2 = { 2,2 };
	*/

	/*
	start1 = { 0,0 };
	end1 = { 1,0 };
	start2 = { 1,1 };
	end2 = { 0,-1 };
	*/
	
	//start1 = { 0,0 };
	//end1 = { 0,1 };
	//start2 = { 0,2 };
	//end2 = { 0,3 };

	
	/*start1 = { 0,3 };
	end1 = { 0,6 };
	start2 = { 0,1 };
	end2 = { 0,5 };*/
	
	start1 = { 12, -55 };
	end1 = { 59, -60 };
	start2 = { 4, -55 };
	end2 = { 81, -62 };

	auto rval = solution.intersection(start1, end1, start2, end2);
	if(!rval.empty())
	cout << rval[0] << " " << rval[1] << endl;
	cout << "pow(10, -6):" << pow(10, -6) << endl;
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
