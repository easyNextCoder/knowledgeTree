#include <iostream>
#include <string>

using namespace std;

class File {
public:
	string setFileName(string fn)
	{
		fileName = fn;
		return fileName;
	}
	string fileName;
};

class InputFile : virtual public File {};

class OutputFile : virtual public File {};

class IOFile :public InputFile, public OutputFile {

};

class Person {
public:
	Person() = default;
	Person(string name) :personName(name) { ; }
	virtual void walk() const = 0;//是不是纯虚函数对虚继承中的变量初始化没影响
	string personName;
};
void Person:: walk() const
{
	cout << "Person Walk." << endl;
}

class Men :public virtual Person {
public:
	Men(string name):Person(name){ 
	}
	virtual void walk()const//必须继承并实现虚类中的接口
	{
		cout << "Men walk." << endl;
	}
	virtual void pee() 
	{
		cout << "Men pee." << endl;
	};
};

class MenActor : public Men {
public:
	
	MenActor(string name) :Men(name){ 
	; }
	void Walk()const
	{
		cout << "MenActor walk." << endl;
	}
	void act()
	{
		cout << "MenActor act." << endl;
	}
};

int main(){
	cout << "======================TEST1======================" << endl;
	IOFile iof;
	//cout << iof.InputFile::setFileName("xiaohua") << endl;
	cout << iof.setFileName("xiaohua") << endl;


	cout << "======================TEST2======================" << endl;
	MenActor huge("胡歌");
	//当Men虚继承Person的时候，通过继承类无法间接初始化最基类（可以通过增删virtual关键词来测试）
	//所以说为了避免这个问题，虚基类中尽量不要放变量，而只放函数接口，来实现合理的接口和实现的配置
	
	//Person oldAgent;错误抽象类不允许生成对象
	huge.act();
	huge.walk();
	cout<<huge.personName << endl;

	return 0;
}