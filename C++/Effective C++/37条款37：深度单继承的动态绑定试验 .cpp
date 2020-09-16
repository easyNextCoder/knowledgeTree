//深度单继承的动态绑定试验 
#include <iostream>

using namespace std;

class Base{
	public:
		virtual void get(){
			cout<<"Base get()."<<endl;
		}
		
		void set(){
			cout<<"Base set()."<<endl;
		}
}; 

class Derived:public Base{
	public:
		virtual void get(){
			cout<<"Derived get()."<<endl;//Derive
		}
		
		void set(){
			cout<<"Derived set()."<<endl; 
		}
};

class DeepDerived:public Derived{
	public:
		virtual void get(){
			cout<<"DeepDerived get()."<<endl;//Derive
		}
		
		void set(){
			cout<<"DeepDerived set()."<<endl; 
		}
};

int main()
{
	Base * agenta = new Derived();
	Base * agentb = new DeepDerived();
	
	agenta->get();
	agentb->get();
	
		
	return 0;
} 


