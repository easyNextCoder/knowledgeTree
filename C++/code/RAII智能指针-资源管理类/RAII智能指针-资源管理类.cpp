#include <iostream>
#include <memory>

using namespace std;

class Mea {
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

class  MeaAlloctor
{
	//using namespace std;
	
};

void MeaDeleter(Mea* meaObj)
{
	cout << "I am MeaDeleter()." << endl;
}

shared_ptr<Mea> createMea()
{
	return make_shared<Mea>();
}

void sendMea_to(weak_ptr<Mea>&obj, shared_ptr<Mea>&obj1)
{
	obj = obj1;
}
int main()
{
	/*测试 weak 和 shared ,这两个对象是息息相关的*/
	weak_ptr<Mea> wpo;
	{
		shared_ptr<Mea> spo(new Mea, MeaDeleter);
		spo.get()->set_name("xuxiaokang");
		cout << spo.get()->get_name() << endl;
		shared_ptr<Mea> spo1 = spo;
		cout << spo1.get()->get_name() << endl;
		sendMea_to(wpo, spo1);
	}
	cout <<( wpo.lock()  == nullptr )<< endl;
	//cout << (bw).lock()->get_name() << endl;//由于源对象b已经被析构了所以当lock的时候会失败
	

	/* 测试unique指针，这个指针是有自己独特的性质 */
	shared_ptr<Mea> spo2(new Mea);
	//unique_ptr<Mea> upo(spo2.get());//unique_ptr<Mea>不支持拷贝,所以进行拷贝的时候会出错
	unique_ptr<Mea> upo1;
	//upo1.reset(spo2.get());//虽然编译可以通过，但是reset失败
	upo1.reset(new Mea);

	cout << (upo1.release())->get_name() << endl;


	return 0;
}