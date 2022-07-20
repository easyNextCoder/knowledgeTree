//这道题自己不会做，看的题解区的第一个人写的代码
#include <iostream>
#include <vector>

using namespace std;
int m, n;
 

void link(int ori, int i, int j, int value, int vec[][1001])
{
    if(i-1>=0 && vec[i-1][j] == ori)
    {
        vec[i-1][j] = value;
        link(ori, i-1, j, value, vec);
    }
    if(i+1<m && vec[i+1][j] == ori)
    {
        vec[i+1][j] = value;
        link(ori, i+1, j, value, vec);
    }
    if(j-1>=0 && vec[i][j-1] == ori)
    {
        vec[i][j-1] = value;
        link(ori, i, j-1, value, vec);
    }
    if(j+1<n && vec[i][j+1] == ori)
    {
        vec[i][j+1] = value;
        link(ori, i, j+1, value, vec);
    }
}
const int N = 1001;
int vec[N][N];
int main()
{
    cin>>m>>n; 
    cin.ignore();
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
            cin>>vec[i][j];
        }
        cin.ignore();
    }
    /*
    //先从顶边找到一个水的起始点，使用dfs进行联通
    for(int j = 0; j<n; j++)
    {
        if(vec[0][j] == 0)
        {
            vec[0][j] = -1;
            link(0, 0, j, -1, vec);
            //联通图采用dfs进行联通图的连接
        }
    }
    //将陆地中的水岛消除
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
            if(vec[i][j] == 0)//消除水的孤岛
                vec[i][j] = 1;
        }
    }
    link(1, m-1, n-1, 2, vec);
    */
    //统计所有连接的陆地
    cout<<"we output:"<<endl;
    int count = 0;
    for(int i = 0; i<m; i++)
    {
        for(int j = 0; j<n; j++)
        {
            cout<<vec[i][j]<<" ";
            //if(vec[i][j] == 2)
                //count++;
        }
        cout<<endl;
    }
    
    cout<<count<<endl;
    return 0;
}
