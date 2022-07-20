#include <map>
//map multimap
#include <unordered_map>
//unordered_map
#include <iostream>
#include <string>

using namespace std;

int main(){
	cout<<"map"<<endl;
	map<string, int>nmap;
	nmap.insert(make_pair("xiaohong", 1));
	
	cout<<nmap.count("xiaohong")<<endl;
	nmap.erase("xiaohong");
	cout<<nmap.count("xiaohong")<<endl;;
	
	cout<<"unordered_map"<<endl;
	unordered_map<string, int>nrmap;
	nrmap.insert(make_pair("xiaohong", 1));
	
	cout<<nrmap.count("xiaohong")<<endl;;
	nrmap.erase("xiaohong");
	cout<<nrmap.count("xiaohong")<<endl;
	
	
	cout<<"multimap"<<endl;
	multimap<string, int>mmap;
	mmap.insert(make_pair("xiaohong", 1));
	mmap.insert(make_pair("xiaohong", 1));
	
	cout<<mmap.count("xiaohong")<<endl;;
	mmap.erase("xiaohong");
	cout<<mmap.count("xiaohong")<<endl;
	
	
	return 0;
}
