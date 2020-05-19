#include <iostream>
#include <iterator>
#include <vector>
#include <algorithm>

using namespace std;

int main(){
	vector<int>vec(10,2);
	ostream_iterator<int>out_iter(cout,",");
	copy(vec.begin(), vec.end(), out_iter);
	
	return 0;
}

