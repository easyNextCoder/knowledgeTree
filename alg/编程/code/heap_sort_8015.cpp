#include <iostream>
#include <vector>

using namespace std;

void processDown(int father, int end, vector<int>& vec)
{
    //int end = vec.size();
    while(father*2 < end)
    {
        int son = father * 2;
        int son1 = father * 2 + 1;
        if(son1 < end)
            if(vec[son] < vec[son1])
                swap(vec[son], vec[son1]);
        if(vec[father] < vec[son])
        {
            swap(vec[father], vec[son]);
        }
        father*=2;
    }
}

int main()
{
    vector<int> vec = {0,1,2,3,4,5,6,7,8,9};
    for(int i = vec.size()/2; i>0; --i)
    {
        processDown(i, vec.size(), vec);
    }

    for(auto item:vec)
        cout<<item<<" ";
    cout<<endl;


    for(int i = vec.size()-1; i>1; --i)
    {
        swap(vec[1], vec[i]);
        processDown(1, i, vec);
    }

    for(auto item:vec)
        cout<<item<<" ";
    cout<<endl;
    return 0;
}