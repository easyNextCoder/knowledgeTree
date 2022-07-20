#include <iostream>

using namespace std;
//普通的模板类从OriType类型中提取出OriType类型
template <typename OriType>
struct Base{
    OriType a;
    OriType* b;
    void set(OriType value){
        a = value;
        cout<<"invoke base."<<endl;    
    }
};
//使用特例化的类从OriType*类型中提取出OriType类型
template <typename T>
struct Base<T*>
{
    T a;
    T aa;
    T* b;
    void set(T value){
        a = value;
        cout<<"invoke special."<<endl;    
    }
};

int main()
{
    Base<int*> x;
    int* p = new int(100);
    x.set(*p);
    cout<<typeid(x.a).name()<<endl;;
    cout<<typeid(x.a).hash_code()<<endl;;
    
    cout<<typeid(p).name()<<endl;
    
    cout<<x.a<<endl;
    return 0;
}