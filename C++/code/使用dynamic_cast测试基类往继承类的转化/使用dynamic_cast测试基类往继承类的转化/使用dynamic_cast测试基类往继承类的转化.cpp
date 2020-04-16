#include <iostream>
#include <vector>

using namespace std;

class Window {
public:
	
	virtual void blink() {
		; cout << "I am Window blinking." << endl;
	}
	virtual ~Window()
	{
		cout << "I am deconstructor of Window." << endl;
	};
};

class SpecialWindow :public Window {
public:
	virtual void blink() {
		cout << "I am SpecialWindow blinking." << endl;
	}
	 void dosomething(){}
	~SpecialWindow() {
		cout << "I am deconstructor of SpecialWindow." << endl;
	}
};

class VerySpecialWindow :public Window {
public:
	virtual void blink() {
		cout << "I am VerySpecialWindow blinking." << endl;
	}
	void dosomething() {}
	~VerySpecialWindow() {
		cout << "I am deconstructor of VerySpecialWindow." << endl;
	}
};

int main() {
	vector<shared_ptr<Window> >vec;
	//vec.push_back(shared_ptr<Window>(new SpecialWindow()));
	vec.push_back(shared_ptr<Window>(new SpecialWindow()));
	vec.push_back(shared_ptr<Window>(new VerySpecialWindow()));

	for (auto iter = vec.begin(); iter != vec.end(); iter++) {
		(iter->get())->blink();
		//psw->blink();
	}


	return 0;
}
