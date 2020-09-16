#include <iostream>
#include <vector>
#include <algorithm>


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


void quicksort(vector<int>& vec, int first, int last){
    
    if(last-first+1 < 3)
    {
        if(last-first+1 == 1)
            return;
        else{
            if(vec[first] > vec[last])
                swap(vec[first], vec[last]);
        }
    }else{
        int provit = getprovit(vec, first, last);
        swap(vec[first+(last-first)/2], vec[last-1]);
        int i = first;//这里出现问题
        int j = last-1;
        for(;;)
        {
            while(vec[++i] < provit){;}
            while(vec[--j] > provit){;}
            if(i<j){
                swap(vec[i], vec[j]);
            }else{
                break;
            }
        }
        swap(vec[i], vec[last-1]);
        quicksort(vec, first, i-1);
        quicksort(vec, i+1, last);
    }

}

int main()
{
    vector<int> in = {4,5,6,7,8,9,10};//{10, 9, 8,7,6,5,4,3,2,1,0};
    vector<int> in1 = {10, 9, 8,7,6,5,4,3,2,1,0};
    //sort(in.begin(), in.end());
    quicksort(in1, 0, in1.size()-1);
    for(auto item:in1)
        cout<<item<<endl;

    return 0;
}


