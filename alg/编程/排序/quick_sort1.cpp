/*
    data:2020.8.26
    for:检验自己写quicksort的出错点：1.vec[++i]  vec[--j]  2. if(len<3){return;}
 */


#include <iostream>
#include <vector>

using namespace std;

int getprovit(vector<int>& vec, int first, int last)
{
    int mid = first+(last-first)/2;
    if(vec[first]>vec[last])
        swap(vec[first], vec[last]);
    if(vec[mid] > vec[last])
        swap(vec[mid], vec[last]);
    if(vec[mid] < vec[first])
        swap(vec[mid], vec[first]);
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
    }else{
    
        if(first == 6 && last == 10)
        {
            cout<<"610"<<endl;
            for(auto item:vec)
                cout<<item<<" ";
            cout<<endl;
        }
        int provit = getprovit(vec, first, last);
        int mid = first + (last-first)/2;
        if(first == 6 && last == 10)
        {
            cout<<"610"<<endl;
            for(auto item:vec)
                cout<<item<<" ";
            cout<<endl;
        }
        swap(vec[mid], vec[last-1]);
        if(first == 6 && last == 10)
        {
            cout<<"after swap610"<<endl;
            for(auto item:vec)
                cout<<item<<" ";
            cout<<endl;
        }
        int i = first;
        int j = last-1;
        
        while(i<j)
        {
            cout<<i<<":"<<j<<endl;
            if(vec[i] < provit)i++;
            if(vec[j] > provit)j--;
            if(i<j)
            {
                swap(vec[i], vec[j]);
            }
        }
        cout<<"i == "<<i<<endl;
        // for(;;)
        // {
        //     while(vec[++i]<provit){;}
        //     while(vec[--j]>provit){;}
        //     if(i<j)
        //         swap(vec[i], vec[j]);
        //     else
        //     {
        //         break;   
        //     }
        // }
        for(auto item:vec)
            cout<<item<<" ";
        cout<<endl;
        cout<<first<<":"<<last<<endl;
        swap(vec[i], vec[last-1]);
        
        qsort(first, i-1, vec);

        qsort(i+1, last, vec);
    }
}


int main()
{

    vector<int> in1 = {10, 9, 8,7,6,5,4,3,2,1,0};
    //vector<int> in1 = {1,3,2};
    //sort(in.begin(), in.end());
    qsort(0, in1.size()-1, in1);
    for(auto item:in1)
        cout<<item<<endl;

    return 0;
}