#include <iostream>
#include <vector>

using namespace std;

class Window {
public:
	
	virtual ~Window()
	{
		cout << "I am deconstructor of Window." << endl;
	};
};

class SpecialWindow :public Window {
public:
	virtual void blink() {
		cout << "I am blinking." << endl;
	}
	 void dosomething(){}
	~SpecialWindow() {
		cout << "I am deconstructor of SpecialWindow." << endl;
	}
};


int main() {
	vector<shared_ptr<Window> >vec;
	vec.push_back(shared_ptr<Window>(new SpecialWindow()));
	vec.push_back(shared_ptr<Window>(new SpecialWindow()));
	vec.push_back(shared_ptr<Window>(new SpecialWindow()));

	for (auto iter = vec.begin(); iter != vec.end(); iter++) {
		SpecialWindow* psw =  dynamic_cast<SpecialWindow*>(iter->get());
		psw->blink();
	}


	return 0;
}
