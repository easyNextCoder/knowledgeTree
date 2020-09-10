#include <iostream>
#include <string>
#include <vector>

using namespace std;

class A{

public:
    A(){
        cout<<"invoke A()."<<endl;
    };
    A(const A& inputa)
    {
        cout<<"invoke A(const A& inputa)."<<endl;
        s = inputa.s;
    }
    A(A&& inputa)
    {
        cout<<"invoke A(A&& inputa)."<<endl;
        s = inputa.s;
        //inputa.s = "";
    }
    // A(const A inputa)
    // {
    //     cout<<"invoke A(const A inputa)."<<endl;
    //     s = inputa.s;
    // }
    A(const string& inputs)
    {
        cout<<"invoke A(const string& inputs)."<<endl;
        s = inputs;
    }
    // A(string&& inputs)
    // {
    //     cout<<"invoke A(const string&& inputs)."<<endl;
    //     s = inputs;
    // }
    // A(const string inputs)
    // {
    //     cout<<"invoke A(const string inputs)."<<endl;
    //     s = inputs;
    // }

    string s;
};

int main()
{
    A a;
    A b("hello");
    A c(string("world"));
    string sa("name");
    A d(sa);
    A e(std::move(sa));
    cout<<"test vector."<<endl;
    vector<A> vec;
    vec.emplace_back(b);
    cout<<"emplace_back()"<<endl;
    vec.push_back(a);
    
    

    return 0;
}