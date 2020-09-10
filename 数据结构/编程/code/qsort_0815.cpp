#include <iostream>
#include <vector>
#include <algorithm>
#include <random>

using namespace std;

int get_provit(int first, int last, vector<int>& vec)
{
    int mid = first + (last-first)/2;
    if(vec[first] > vec[last])
        swap(vec[first], vec[last]);
    if(vec[mid] < vec[first])
        swap(vec[mid], vec[first]);
    if(vec[mid] > vec[last])
        swap(vec[mid], vec[last]);
    return vec[mid];
}

void qsort(int first, int last, vector<int>&vec)
{
    int len = last-first+1;
    if(len < 3)
    {
        if(len == 2 && vec[first] > vec[last])
            swap(vec[first], vec[last]);
        return;
    }

    int provit = get_provit(first, last, vec);
    int mid = first+(last-first)/2;
    swap(vec[mid],vec[last-1]);
    int i = first;
    int j = last-1;
    for(;;)
    {
        
        while(vec[++i] < provit){;}
        while(vec[--j] > provit){;}
        if(i < j)
            swap(vec[i], vec[j]);
        else
            break;
    }
    swap(vec[i], vec[last-1]);
    qsort(first, i-1, vec);
    qsort(i+1, last, vec);
    return;
}

int main()
{
    //vector<int> vec = {0,1,2,3,4,5,6,7,8,9};
    
    //reverse(vec.begin(), vec.end());
    vector<int> vec(1200);
    default_random_engine e;
    
    //generate(vec.begin(), vec.end(), e);
    for(int i = 0; i<1200; ++i)
    {
        vec[i] = e()%1200;
    }
    //sort(vec.begin(), vec.end());
    qsort(0, vec.size()-1, vec);


    for(auto item:vec)
        cout<<item<<" ";
    cout<<endl;

    for(int i = 0; i<vec.size()-1; ++i)
    {
        if(vec[i] > vec[i+1])
        {
            cout<<"false"<<" ";
        }
    }
    cout<<"sucess."<<endl;
    return 0;
}