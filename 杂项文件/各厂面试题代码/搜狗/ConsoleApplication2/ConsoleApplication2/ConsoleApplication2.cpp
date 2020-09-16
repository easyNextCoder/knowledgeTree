#include <iostream>
#include <vector>

using namespace std;
vector<int> out;
bool get(vector<int>& con, int index, int k, int tmp, vector<int>& res)
{

	if (tmp == k)
	{
		if(out.empty())
			out = res;
		return true;
	}
	else if (index == con.size()) {
		return false;
	}
	else {
		if (tmp < k)
		{
			get(con, index + 1, k, tmp, res);
			res.push_back(con[index]);
			bool rval = get(con, index + 1, k, tmp + con[index], res);
			if (rval)
			{
				return true;
			}
			res.pop_back();
			
		}
		else {
			return false;
		}
	}
}

int main()
{
	int n, k;
	n = 4;
	k = 13;
	vector<int> con = { 1,2,4,7 };
	vector<int> res;
	bool rval = get(con, 0, k, 0, res);
	if (rval)
	{
		cout << "YES" << endl;
		for (int i = 1; i < out.size(); ++i)
			cout << out[i] << " ";
	}
	else {
		cout << "NO" << endl;
	}

	return 0;
}