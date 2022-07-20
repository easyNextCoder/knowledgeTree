#include <iostream>

using namespace std;


class Base {
	friend void testFunc(Base& input);
public:
	int pub = 0;
protected:
	int pro = 0;

private:
	int pri = 0;
};

void testFunc(Base& input)
{
	input.pri = 0;
}


class Derivedpub:public Base {
public:
	
	void setBasepub(int value)
	{
		pub = value;
	}
	void setBasepro(int value)
	{
		pro = value;
	}
	void setBasepri(int value)
	{
		pri = value;
	}
};

class Derivedpro :protected Base {
public:
	void setBasepub(int value)
	{
		pub = value;
	}
	void setBasepro(int value)
	{
		pro = value;
	}
	void setBasepri(int value)
	{
		pri = value;
	}
};

class Derivedpri :private Base {
public:
	void setBasepub(int value)
	{
		pub = value;
	}
	void setBasepro(int value)
	{
		pro = value;
	}
	void setBasepri(int value)
	{
		pri = value;
	}
};

int main()
{
	Derivedpub puba;
	Derivedpro proa;
	Derivedpri pria;

	/*
	cout << "其实protected有意义吗？" << endl;

	puba.pub = 0;
	puba.pro = 0;

	proa.pub = 0;
	proa.pro = 0;
	*/



}