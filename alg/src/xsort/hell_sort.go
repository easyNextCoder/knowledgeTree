package xsort

/*
#include <iostream>
#include <vector>

using namespace std;

void hell_sort(vector<int>&vec, int first, int last)
{
    for(int i = first; i<last; ++i)
    {

        for(int N = (last-first+1)/2; N>=1; N/=2)
        {
            int tmp_min_index = i;
            for(int j = i+N; j<=last; j+=N)
            {
                if(vec[i] > vec[j])
                {
                    tmp_min_index = j;
                }
            }

            swap(vec[i], vec[tmp_min_index]);
        }

    }
}

int main()
{
    vector<int> vec = {9,8,7,6,5,4,3,2,1,0};
    hell_sort(vec, 0, vec.size()-1);
    for(auto item:vec)
    {
        cout<<item<<endl;
    }
    return 0;
}
*/
