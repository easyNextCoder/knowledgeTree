#include <iostream>
#include <vector>
#include <algorithm>


using namespace std;


void processDown(vector<int>& in, int n, int end)
{
    int father = n;
    //这里再次忘了+1导致的最终的失误
    while(father*2+1 < end)
    {
        //这里再次忘了+1导致的错误
        int son = father*2+1;
        if(son+1<end)
        {
            if(in[son] < in[son+1])
                son = son+1;
        }
        
        if(in[father] < in[son]){
            swap(in[father], in[son]);
            father = son;
        }else{
            break;
        } 
    }
}

int main()
{
    vector<int> in = {4,5,6,7,8,9,10};
	//{10, 9, 8,7,6,5,4,3,2,1,0};

    for(int i = in.size()/2-1; i>=0; --i)
    {
        processDown(in, i, in.size());
    }
    for(auto item:in)
        cout<<item<<endl;
    int end = in.size();
    for(int i = 0; i<in.size(); i++)
    {
        swap(in[0], in[--end]);
        processDown(in, 0, end);
    }

    for(auto item:in)
        cout<<item<<endl;


    return 0;
}