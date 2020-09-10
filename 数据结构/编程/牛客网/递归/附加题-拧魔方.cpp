#include <iostream>

using namespace std;

const int N = 24;
int nero[N];
//20:40
int top[4] = {22,23,21,20};
int toplayer[8] = {0,1,9,15,19,18,10,4};
int bottom[4] = {7,6,12,13};
int bottomlayer[8] = {3,2,5,11,16,17,14,8};

int before[4] = {18,19,17,16};
int beforelayer[8] = {20,21,15,14,13,12,11,10};
int after[4] = {1,0,2,3};
int afterlayer[8] = {23,22,4,5,6,7,8,9};

int mleft[4] = {4,10,11,5};
int leftlayer[8] = {22,20,18,16,12,6,2,0};
int mright[4] = {15,9,8,14};
int rightlayer[8] = {21,23,1,3,7,13,17,19};

int * layer[6][2] = {
    {top, toplayer},
    {bottom, bottomlayer},
    {before, beforelayer},
    {after, afterlayer},
    {mleft, leftlayer},
    {mright, rightlayer},
    
};

void rotateCW(int n, int len, int* arr, int* rarr)
{//顺时针
    int* tmp = new int[n];
    int count = 0;
    for(int i = len-n; i<len; i++)
        tmp[count++] = rarr[arr[i]];
    for(int i = len-n; i>=0; i-=n)
        for(int j = 0; j<n; j++)
            if(i != 0)
                rarr[arr[i+j]] = rarr[arr[i-n+j]];
        
    for(int i = 0; i<n; i++)
        rarr[arr[i]] = tmp[i];
}

void rotateCCW(int n, int len, int* arr, int* rarr)
{//逆时针
    int* tmp = new int[n];
    int count = 0;
    for(int i = 0; i<n; i++)
        tmp[count++] = rarr[arr[i]];
    for(int i = 0; i<len; i+=n)
        for(int j = 0; j<n; j++)
            if(i != len-n)
                rarr[arr[i+j]] = rarr[arr[i+n+j]];
    count = 0;
    for(int i = len-n; i<len; i++)
        rarr[arr[i]] = tmp[count++];
}

int cal(int *arr)
{
    int rval = 0;
    rval+=arr[0]*arr[1]*arr[2]*arr[3];
    rval+=arr[4]*arr[5]*arr[10]*arr[11];
    rval+=arr[6]*arr[7]*arr[12]*arr[13];
    rval+=arr[8]*arr[9]*arr[14]*arr[15];
    rval+=arr[16]*arr[17]*arr[18]*arr[19];
    rval+=arr[20]*arr[21]*arr[22]*arr[23];
    return rval;
}
int rvalMax = 0;
void neroSwitch(int u, int n)
{
    if(u == n)return;
    else{
        rvalMax = max(rvalMax, cal(nero));
    }
    for(int i = 0; i<6; i++)
    {
        
        //这一面顺时针旋转
        rotateCW(1, 4, layer[i][0], nero);
        rotateCW(2, 8, layer[i][1], nero);
        neroSwitch(u+1, n);
        rotateCCW(1, 4, layer[i][0], nero);
        rotateCCW(2, 8, layer[i][1], nero);
        //也有可能是逆时针
        rotateCCW(1, 4, layer[i][0], nero);
        rotateCCW(2, 8, layer[i][1], nero);
        neroSwitch(u+1, n);
        rotateCW(1, 4, layer[i][0], nero);
        rotateCW(2, 8, layer[i][1], nero);
        
    }
}

int main()
{
    for(int i = 0; i<N; i++)
        cin>>nero[i];
    //如果是转上面一层
    //cout<<cal(nero);
    
     for(int i = 0; i<6; i++)
    {
        
        //这一面顺时针旋转
        rotateCW(1, 4, layer[i][0], nero);
        rotateCW(2, 8, layer[i][1], nero);
        rotateCCW(1, 4, layer[i][0], nero);
        rotateCCW(2, 8, layer[i][1], nero);
        //也有可能是逆时针
        rotateCCW(1, 4, layer[i][0], nero);
        rotateCCW(2, 8, layer[i][1], nero);
        rotateCW(1, 4, layer[i][0], nero);
        rotateCW(2, 8, layer[i][1], nero);
        
    }

    //for(int i = 0; i<24; i++)
        //cout<<nero[i]<<endl;

    neroSwitch(0, 6);
    cout<<rvalMax<<endl;
    
    
    return 0;
}