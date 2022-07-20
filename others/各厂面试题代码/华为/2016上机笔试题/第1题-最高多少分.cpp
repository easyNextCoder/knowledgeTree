//https://www.nowcoder.com/questionTerminal/3897c2bcc87943ed98d8e0b9e18c4666
//这题坑的地方在于：文字描述中说了有多组数据，但是样例中只给了一组数据！浪费30分钟

#include <iostream>
#include <vector>

using namespace std;

int main()
{
	int n, m;
	while (cin >> n >> m)
	{
		vector<int>grade(n, 0);
		for (int i = 0; i < n; i++)cin >> grade[i];
		while (m--)
		{
			char c;
			int first = -1;
			int last = -1;
			cin >> c >> first >> last;
			if (c == 'Q')
			{
				int max1 = -1;
				if (first > last)swap(first, last);
				for (int i = first - 1; i < last; i++)
				{
					max1 = max(grade[i], max1);
				}
				cout << max1 << endl;
			}
			else {
				grade[first - 1] = last;
			}
		}

	}

	return 0;
}