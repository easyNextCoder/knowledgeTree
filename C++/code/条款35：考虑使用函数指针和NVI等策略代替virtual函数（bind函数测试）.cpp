#include <iostream>
#include <algorithm>
#include <functional>

using namespace std;

using namespace std;

class Fake{
	
};

void func(Fake& fake, int num)
{
	cout<<"this is the "<<num<<endl;	
} 

int main()
{
	Fake fk;
	function<void (int)> fbody = bind(func, fk, placeholders::_1);
	
	//void (*pbody)(int) = bind(func, fk, placeholders::_1);
	auto pbody = bind(func, fk, placeholders::_1);
	//只能使用auto 或者function类来进行bind绑定 
	for(int i = 0; i<20; i++)
	{
		fbody(i);
		pbody(i);
	}
	
	return 0;
}
