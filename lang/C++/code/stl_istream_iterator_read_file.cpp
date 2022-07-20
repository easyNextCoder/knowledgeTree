#include <iostream>
#include <fstream>
#include <iterator>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

int main(){
	vector<int>vec(10,2);
	fstream File("stl_istream_iterator.cpp");
	istream_iterator<string>eos;
	istream_iterator<string>out_iter(File);
	while(out_iter != eos)
	{
		cout<<*out_iter<<endl;
		++out_iter;
	}
	
	return 0;
}

