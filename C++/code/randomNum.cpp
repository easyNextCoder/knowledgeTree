#include<random>
#include <iostream>

using namespace std;

int main(){
	default_random_engine e;
	for(int i = 0; i<10; i++){
		cout<<e()%54<<endl;
	}
	return 0;
}
