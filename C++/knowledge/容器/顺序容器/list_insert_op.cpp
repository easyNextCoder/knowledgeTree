#include <iostream>
#include <list>
#include <random>

using namespace std;
void list_insert(list<int>& mylist, int x)
{
    list<int>::iterator iter= mylist.begin();
    while(iter != mylist.end())
    {
        if(*iter <= x )
        {
            iter++;
        }else{
            //插入的是当前迭代器的前面一个位置
            mylist.insert(iter, x);
            break;
        }
    }
    //下面这一行一定不能少，
    //即能处理向空list插入元素，又能处理插入值是当前list中最大的
    if(iter == mylist.end())mylist.insert(mylist.end(), x);
}
int main()
{
    list<int>mylist;
          
    default_random_engine e;
    int count = 100;
    while(count-->0)
    {
        list_insert(mylist, e()%200);
    }
    
    for(auto item:mylist)
    {
        cout<<item<<" ";
    }
    cout<<endl;

    return 0;
}