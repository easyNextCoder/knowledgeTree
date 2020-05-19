#include <iostream>
#include <vector>
//https://en.cppreference.com/w/cpp/io/basic_ios/operator_bool
using namespace std;

int main()
{
	vector<int> a;

	int n = 0;
	while(cin>>n)
	{
		cout<<n<<endl;
	} 
	cout<<"out of while(cin>>)."<<endl;

	return 0;
}
