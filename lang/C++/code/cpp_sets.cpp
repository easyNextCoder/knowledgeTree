#include <set>
//set multiset
#include <unordered_set>
//unordered_set
#include <iostream>

using namespace std;

int main(){
	cout<<"set"<<endl;
	set<int>nset;
	nset.insert(1);
	cout<<nset.count(1)<<endl;;
	nset.erase(1);
	cout<<nset.count(1)<<endl;;
	
	cout<<"unordered_set"<<endl;
	unordered_set<int>nrset;
	nrset.insert(2);
	cout<<nrset.count(2)<<endl;;
	nrset.erase(2);
	cout<<nrset.count(2)<<endl;
	
	cout<<"multiset"<<endl;
	multiset<int>mset;
	mset.insert(2);
	mset.insert(2);
	cout<<mset.count(2)<<endl;;
	mset.erase(2);
	cout<<mset.count(2)<<endl;
	
	
	return 0;
}
