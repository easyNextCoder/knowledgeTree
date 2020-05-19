#include <iostream>

using namespace std;
class Empty { };

void f()
{
	Empty a, b;
	if (&a == &b) cout << "impossible: report error to compiler supplier";
	//cout<<((unsigned int)(reinterpret_cast<unsigned char*>(&a))|(unsigned int)(-1))<<endl;;
	//×Ô¼ºµÄ²âÊÔ 
	Empty* p1 = new Empty;
	Empty* p2 = new Empty;
	
	if (p1 == p2) cout << "impossible: report error to compiler supplier";
}

struct X : Empty {
	int a;
	// ...
};

void f(X* p)
{
	void* p1 = p;
	void* p2 = &p->a;
	if (p1 == p2) cout << "nice: good optimizer";
}

int main()
{
	f();
	f(new X);
	return 0;
}
