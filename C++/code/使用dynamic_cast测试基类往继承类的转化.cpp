#include <iostream>
#include <vector>

using namespace std;

class Window {
public:
	
	Window(int ia, int ib){
		ia = a;
		ib = b;
	}
	
	virtual void blink() {
		; cout << "I am Window blinking." << endl;
	}
	virtual ~Window()
	{
		cout << "I am deconstructor of Window." << endl;
	};
private:
	int a;
	int b;
};

class SpecialWindow :public Window {
public:
	SpecialWindow(int a, int b):Window(a,b){

	}
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
	VerySpecialWindow(int a, int b) :Window(a, b){

	}
	
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
	vec.push_back(shared_ptr<Window>(new SpecialWindow(1,2)));
	vec.push_back(shared_ptr<Window>(new VerySpecialWindow(2,3)));
	make_shared<Window>(1, 3);
	for (auto iter = vec.begin(); iter != vec.end(); iter++) {
		(iter->get())->blink();
		//psw->blink();
	}


	return 0;
}
