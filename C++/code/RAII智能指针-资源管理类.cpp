#include <iostream>
#include <memory>

using namespace std;

class Mea{
public:
	
	string get_name()
	{
		return name;
	}

	void set_name(string s)
	{
		name = s;
	}
	
private:
	string name = "xuyongkang";

};
void MeaDeleter(Mea* meaObj)
{
	cout<<"I am MeaDeleter()."<<endl;
}
shared_ptr<Mea> createMea()
{
	return make_shared<Mea>();
}
int main()
{
	shared_ptr<Mea> b(createMea(), MeaDeleter);
	b.get()->set_name("xuxiaokang");
	cout<<b.get()->get_name()<<endl;
	shared_ptr<Mea, MeaDeleter> bf = b;
	cout<<bf.get()->get_name()<<endl;
	
	
	return 0;
}
