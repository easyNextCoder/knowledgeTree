#include <iostream>
#include <vector>

using namespace std;

class Window{
}; 

class SpecialWindow:publc Window{
	void blink(){
		cout<<"I am blinking."<<endl;
	}
	~SpecialWindow(){
		cout<<"I am deconstructor of SpecialWindow."<<endl;
	}
};


int main(){
	vector<shared_ptr<Window> >vec;
	vec.push_back(new SpecialWindow());
	
	
		
	return 0;
}

