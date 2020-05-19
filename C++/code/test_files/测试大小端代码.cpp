#include <iostream>

using namespace std;

int main()
{
	unsigned int a = 0x1234;
	unsigned char* p = (unsigned char*)&a;
	for(int i = 0; i<4; i++){
		cout<<*(p--)<<endl;
	}
	return 0;
}
