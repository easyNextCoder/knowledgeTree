#include <iostream>
#include <vector>
using namespace std;
int main() {
    //int a;
    //cin >> a;
    vector<vector<int>> vec = {
        /*
        {1,2,3,4,5},
        {6,7,8,9,10},
        {11,12,13,14,15}
        */
       {1,2},
       {3,4}
    };
    if(vec.empty())return 0;
    int nrow = vec.size();
    int ncol = vec.front().size();
    for(int i = 0; i<min(nrow/2, ncol/2)+1; ++i)
    {
        // cout<<"left-right"<<endl;
        for(int j = i; j<ncol-i-1; ++j)
        {
            cout<<vec[i][j]<<" ";
        }
        // cout<<endl;
        // cout<<"up-down"<<endl;
        for(int j = i; j<nrow-i-1; ++j)
        {
            cout<<vec[j][ncol-i-1]<<" ";
        }
        // cout<<endl;
        // cout<<"right-left"<<endl;
        for(int j = ncol-i-1; j>i; --j)
        {
            // cout<<vec[nrow-i-1][j]<<" ";
            if(nrow-i-1 == i)
                break;
        }
        // cout<<endl;
        // cout<<"down-up"<<endl;
        

        for(int j = nrow-i-1; j>i; --j)
        {
            // cout<<vec[j][i]<<" ";
            if(ncol-i-1 == i)
                break;
        }
        cout<<endl;
    }
    cout<<endl;
    cout << "Hello World!" << endl;
}