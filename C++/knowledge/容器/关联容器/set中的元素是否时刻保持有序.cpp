#include <iostream>
#include <set>
#include <string>

using namespace std;

int main()
{

    set<string> con;

    con.insert("ceef");
    con.insert("ctef");
    
    con.insert("abcd");
    con.insert("bcde");
    con.insert("cdef");

    for(auto s:con)
        cout<<s<<endl;
    con.erase("abcd");
    cout<<"next"<<endl;
    for(auto s:con)
        cout<<s<<endl;
    cout<<endl;
    return 0;
}