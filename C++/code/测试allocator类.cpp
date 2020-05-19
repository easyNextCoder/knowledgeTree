#include <iostream>
#include <memory>

using namespace std;

int main()
{
	allocator<string> alloc;
	auto const p = alloc.allocate(20);
	auto end = p+20;
	auto q = p;
	while(q != end)
	{
		alloc.construct(q++, "name");
	}
	auto z = p;
	while(z != end)
	{
		cout<<*z++<<endl;
		//alloc.destroy(z);
	}
	auto zz = p;
	while(zz != end)
	{
		alloc.destroy(zz++);
	} 
	//¶þ´ÎÏú»Ù 
	alloc.deallocate(p, 20);
	return 0;
}
