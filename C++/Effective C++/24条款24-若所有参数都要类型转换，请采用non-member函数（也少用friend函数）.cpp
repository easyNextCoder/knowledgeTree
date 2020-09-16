/*非成员运算符重载*/
#include <iostream>

using namespace std;

class Rational{
	
	public:
		Rational() = default;
		//explicit Rational(int input):out(input){;}
		Rational(int input):out(input){;}
		Rational& setValueOut(int value)
		{
			out = value;
			return *this;
		}
		int getValueOut()
		{
			return out;
		}
	private:
		int out = 1;
}; 

const Rational operator*(const Rational& item1, const Rational& item2)
{
	Rational rval;
	
	return rval.setValueOut(item1.getValueOut()*item2.getValueOut());
}

int main()
{
	Rational result = 2*Rational(10);
	cout<<result.getValueOut()<<endl;
	return 0;
}


