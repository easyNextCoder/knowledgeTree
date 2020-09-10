#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	/**
	 *
	 * @param person int整型一维数组
	 * @param personLen int person数组长度
	 * @return int整型
	 */
	int house(int* person, int personLen) {
		// write code here
		vector<int> f(personLen, 1);
		vector<int> f1(personLen, 1);

		for (int i = 1; i < personLen; ++i)
		{
			if (person[i] > person[i - 1])
				f[i] = f[i - 1] + 1;

		}

		for (int i = personLen - 2; i >= 0; --i)
		{
			if (person[i] > person[i + 1])
				f1[i] = f1[i + 1] + 1;
		}

		int sum = 0;
		for (int i = 0; i < personLen; ++i)
			sum += max(f[i], f1[i]);
		return sum;
	}
};
int main()
{
	const int N = 3;
	int person[N] = { 3,2,4 };
	Solution so;
	cout<<so.house(person, N) << endl;
	return 0;
}