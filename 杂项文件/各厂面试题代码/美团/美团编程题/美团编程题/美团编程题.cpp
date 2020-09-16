#include <iostream>
#include <vector>
#include <algorithm>
#include <set>
#include <map>
#include <string>

using namespace std;



/*
	if (n * m > 100000000)
	{
		cout << 1 << endl;
		return 0;
	}
	*/

#include <iostream>
#include <vector>
#include <algorithm>
#include <set>
#include <map>
#include <string>

using namespace std;

int main()
{
	int n, m;
	cin >> n >> m;
	vector<int> v(n + 3, 0);
	vector<int> w(n + 3, 0);
	
	vector<vector<int>> f(n + 3, vector<int>(m + 3, 0));

	for (int i = 0; i < n; ++i)
		cin >> v[i + 1] >> w[i + 1];

	for (int i = n; i >= 1; --i)
	{
		for (int j = 1; j <= m; ++j)
		{
			f[i][j] = max(f[i + 1][j], f[i + 1][j - 1] + v[i] + w[i] * 2);
		}
	}
	int maxValue = f[1][m];
	int cnt = m;
	for (int i = 1; i <= n; ++i)
	{
		if (cnt >= 1 && f[i][cnt] == f[i + 1][cnt - 1] + v[i] + w[i] * 2)
		{
			cout << i << " ";
			cnt -= 1;
		}
		if (cnt == 0)
			break;
	}

	return 0;
}



/*
5 5
1 2
3 4
5 6
6 7
3 8


5 4
1 2
3 4
5 6
6 7
3 8

67
2 3 4 5

5 2
5 10
8 9
1 4
7 9
6 10


5 4
1 1
1 1
1 1
1 1
1 1


*/

/*
5 5
1 2
3 4
5 6
6 7
3 8


5 4
1 2
3 4
5 6
6 7
3 8

5 4
1 1
1 1
1 1
1 1
1 1


*/