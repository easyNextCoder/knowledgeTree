#include <iostream>
#include <iterator>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

int main(){
	vector<string>vec;
	istream_iterator<string>eos;
	istream_iterator<string>out_iter(cin);
	
	copy(out_iter, eos, inserter(vec, vec.begin()));
	
	reverse(vec.begin(), vec.end());
	for(auto item:vec)
		cout<<item<<endl;
	//linux��ctrl+D��ʾ���������windows��ctrl+Z��ʾ������� 
	return 0;
}

