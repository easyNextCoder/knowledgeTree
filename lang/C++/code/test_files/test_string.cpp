#include <iostream>
#include <string>
#include <string.h>
using namespace std;

int main(){
	string s = "ms";
	char sa[4];
	//for(int i = 0; i<s.size(); i++)
	//	sa[i] = s[i];
	cout<<s<<endl;
	strcpy(sa,s);
	cout<<sa<<endl;
	return 0;
}
