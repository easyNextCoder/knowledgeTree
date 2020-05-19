#include <iostream>

using namespace std;

int main(){
	for(int i = 0; i<20; i++){
		cout<<(i&1)<<endl;
	}
	
	for(int i = 0; i<20; i++){
		cout<<(i^1)<<endl; 
	}
	return 0;
}
