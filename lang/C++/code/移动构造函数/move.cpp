#include <iostream>
#include <vector>

using namespace std;

class A{

public:
    A(){};
    A(const A& a)
    {
        this->vec = a.vec;
        this->value = a.value;
    }
    void operator=(const A& a)
    {
        
        this->vec = a.vec;
        this->value = a.value;
        delete this->name;
        this->name = a.name;
        cout<<"invoke operator="<<endl;
    }
    void operator=(A&&a)
    {
        cout<<"invoke void operator=(const A&&a){}";
        this->name = a.name;
        a.name = NULL;
        this->vec = a.vec;
        this->value = a.value;
    }
    void push(int val){vec.push_back(val);};
    void pop(){vec.pop_back();};
    void setValue(int a){value = a;};
    int getValue(){return value;};
    int size(){return vec.size();};
    int operator[](int n){
        if(n < vec.size())
            return vec[n];
        else
        {
            const char * out = "overflow.";
            throw(out);
        }
    }
    void setName(string _name){
        name = new char[_name.size()];
        memcpy(name, _name.c_str(), _name.size());
        name[_name.size()-1] = 0;
    }
    const char* getName()
    {
        return name;
    }
    ~A()
    {
        delete name;
        cout<<"invoke deconstructor function."<<endl;
    }
private:
    vector<int> vec;
    int value;
    char* name;
};

int main()
{
    A b;
    {
        A a;
        a.setValue(10);
        cout<<a.getValue()<<endl;
        a.push(99);
        a.push(109);
        a.setName("xuyongkang");
        cout<<a.getName()<<endl;
        b = std::move(a);//延长生命周期，盗取堆上的内存
        //b = a;//得到的结果完全不一样！
        //cout<<"after move:"<<a.getName()<<endl;这里再输出就是NULL了会发生错误
    }
    cout<<"here"<<endl;
    cout<<b.getValue()<<endl;
    cout<<b[1]<<endl;
    cout<<b.getName()<<endl;
    //cout<<a.getName()<<endl;
    return 0;
}