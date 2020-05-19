#include <iostream>
#include <windows.h> 
#include <time.h>
using namespace std;

int main(){
	
	size_t start = clock();
	int a;
	while(a<1000000000){
		a++;
	}
	size_t end = clock();
	cout<<end - start<<endl;
	
	return 0;
}
