include <iostream>
#include <iterator>
#include <vector>
#include <algorithm>

using namespace std;

int main(){
	vector<int>vec(10,2);
	istream_iterator<int>eos;
	istream_iterator<int>out_iter(cin);
	while(out_iter != eos)
	{
		cout<<*out_iter<<endl;
		++out_iter;
	}
	
	return 0;
}

