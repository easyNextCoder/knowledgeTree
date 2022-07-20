  #include <iostream>
 #include <queue>
 #include <stdlib.h>


 using namespace std;
 
 class T
 {
 public:
     int x,y,z;
     T(int a,int b,int c):x(a),y(b),z(c)
     {
     }
 };
 bool operator<(const T&t1,const T&t2)
 {
     return t1.z<t2.z;
 }
 int main(void)
 {
     priority_queue<int, vector<int>, greater<int>>q;
     queue<int> myQueue;
     q.push(10);
     q.push(200);
     q.push(8);
     q.push(1);
     myQueue.push(10);
     
     while(!q.empty())
    {
         int t=q.top();
         q.pop();
         cout<<t<<endl;
		 //cout<<t.x<<" "<<t.y<<" "<<t.z<<endl;
     }
     system("Pause");
     return 1;
 }
