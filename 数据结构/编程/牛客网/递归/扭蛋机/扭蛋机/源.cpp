#include <iostream>
#include <string>
#include <unordered_map>
#include <list>

using namespace std;

/*
int times = 0;
int minTimes = 1e8;
int target = 0;
string tmp;
string out;
unordered_map<int, string> con;
//使用记忆搜索来加速这个方法使用不上！
string dfs(int n)
{
	
    if(n == target)
    {
		if (minTimes > times)
		{
			out = tmp;
			minTimes = times;
		} 
        
    }else if(n > target){
        return false;
    }else{
        
        times++;
        tmp.push_back('2');
		dfs(2*n+1);
        tmp.pop_back();
        times--;
        
        times++;
        tmp.push_back('3');dfs(2*n+2);
        tmp.pop_back();
        times--;
        
    }
}
*/
list<char> out;
int target = 0;
void dfs(int target)
{
	if (target == 0)
	{
		return;
	}

	if (target % 2 == 0)
	{
		out.push_front('3');
		dfs((target - 1) / 2);
	}
	else {
		out.push_front('2');
		dfs((target) / 2);
	}
}

int main()
{
    int n;
    cin>>n;
    target = n;
    dfs(target);
	for(auto item:out)
		cout<<item;
	cout << endl;
    return 0;
}