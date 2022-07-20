#include <iostream>
#include <vector>
#include <algorithm>
#include <random>

using namespace std;

void processDown(vector<int>& vec, int father, int end)
{
    
    
    while(father*2+1 < end)
    {
        
        int son = father*2 + 1;
        //cout<<"son is:"<<son<<endl;
        if(son + 1< end && vec[son+1] > vec[son])
        {
            son = son+1;
        }

        if(vec[father] < vec[son])
        {
            swap(vec[father], vec[son]);
        }
        
        father = son;
    }
}

void heap_sort(vector<int>& vec, int first, int last)
{
    for(int i = vec.size()/2-1; i>=0; i--)
    {
        processDown(vec, i, vec.size());
    }

    for(int i = 0; i<vec.size(); i++)
    {
        swap(vec[0], vec[vec.size()-1-i]);
        
        processDown(vec, 0, vec.size()-i-1);//注意这里不能从vec.size()-i-1开始，因为一开始第一个就已经移动到后面了
        
    }
    cout<<endl;

}

int main()
{
    vector<int> vec = {9,8,7,6,5,4,3,2,1};
    
    //reverse(vec.begin(), vec.end());
    default_random_engine e;
    int count = 1000;
    while(count--)
    {
        vec.push_back(e()%(10000));
    }
    cout<<"before heap_sort:"<<endl;
    for(auto item:vec)
    {
        cout<<item<<" ";
    }
    cout<<endl;
    cout<<"after heap_sort:"<<endl;
    heap_sort(vec, 0, vec.size()-1);
    for(auto item:vec)
    {
        cout<<item<<" ";
    }
    cout<<endl;

}