package xsort

/*
#include <iostream>
#include <vector>

using namespace std;

void bobule_sort(vector<int>&vec, int first, int last)
{
    for(int i = last; i>=first; --i)
    {
        for(int j = first; j<i; ++j)
        {
            if(vec[j] > vec[j+1])
                swap(vec[j], vec[j+1]);
        }
    }
}

int main()
{
    vector<int> vec = {9,8,7,6,5,4,3,2,1,0};
    bobule_sort(vec, 0, vec.size()-1);
    for(auto item:vec)
    {
        cout<<item<<endl;
    }
    return 0;
}

*/
