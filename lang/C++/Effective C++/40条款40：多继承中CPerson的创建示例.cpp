// ����40����̳��е�CPerson�ഴ��ʾ��.cpp : ���ļ����� "main" ����������ִ�н��ڴ˴���ʼ��������
//
#include <iostream>
#include <string>
#include <string.h>

using namespace std;

class IPerson {
public:
    virtual ~IPerson(){};
    virtual string name() const = 0;
    virtual string birthDate() const = 0;
};

class PersonInfo {
public:
    explicit PersonInfo(const char* inName,const char* inBirthData) {
        strcpy(name, inName);
        strcpy(birthDate, inBirthData);
    }

    virtual ~PersonInfo(){};
    virtual const char* theName() const;
    virtual const char* theBirthDate() const;
    virtual const char* valueDelimOpen() const;
    virtual const char* valueDelimClose() const;


private:
    static const int NAME_LEN = 50;
    static const int BD_LEN = 20;
    
    char name[NAME_LEN];
    char birthDate[BD_LEN];
    char value[NAME_LEN + BD_LEN];
};

const char* PersonInfo::valueDelimOpen() const
{
	return "[";
}

const char* PersonInfo::valueDelimClose() const
{
	return "]";
}

const char* PersonInfo::theName() const
{
	strcat(value, valueDelimOpen());
	strcat(value, name);
	strcat(value, valueDelimClose());
	return value;
}

const char* PersonInfo::theBirthDate() const
{
	strcat(value, valueDelimOpen());
	strcat(value, birthDate);
	strcat(value, valueDelimClose());
	return value;
}

class CPerson : public IPerson, private PersonInfo {
public:
    virtual ~CPerson(){};
    CPerson(string name, string birthDate) :
        PersonInfo(name.c_str(), birthDate.c_str()) {

    }
    virtual string name() const
    {
        return theName();//����������õ������ص� ���շ��ź��� 
    };
    virtual string birthDate() const
    {
        return theBirthDate();//����������õ������ص� ���շ��ź��� 
    }
    const char* valueDelimOpen() const
	{
		return "(";
	}
	
	const char* valueDelimClose() const
	{
		return ")";
	}
};
//�������˵���˿������ز�����ֺ��������ҿ���ʹ�û�����δ���صĺ����������ڼ̳����������˵ĺ��� 
int main()
{
	CPerson xiaohua("xiaohua", "2000-10-12");
	cout<<xiaohua.name()<<endl;
	cout<<xiaohua.birthDate()<<endl;

    std::cout << "Hello World!\n";
}


