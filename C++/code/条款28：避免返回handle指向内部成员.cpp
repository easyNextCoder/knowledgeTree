#include <iostream>

using namespace std;


class A{
	public:
		int get(){
			return a;
		}
		typedef int funt(void);
		funt get_get_function(){
			return get;
		} 
	private:
		int a = -1;
};

int have(){
	
	return 0;
}
typedef int havet();
int main(){
	havet* another_have = have;
	cout<<another_have()<<endl;
	return 0;
}
