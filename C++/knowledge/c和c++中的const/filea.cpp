/*
 * 作用：测试C++中const和static对链接属性带来的影响
 * 作者：xyk
 * 日期：2020.5.23
 */
#include <iostream>

using namespace std;
const int vara = 99;
//测试结果：加入static和const之后，外部文件无法再链接
//const已经将其链接属性改成了只能内部链接了
extern int sum(int a, int b);
int main()
{
    int result = sum(10, 10);
    cout<<"the result is: "<<result<<endl;
    cout<<"value vara is: "<<vara<<endl;

}