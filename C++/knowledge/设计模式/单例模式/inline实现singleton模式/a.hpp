//如果下面的一行删掉inline就会出现链接错误，
//因为已经变成了工厂模式，重复引用出错
inline int asum(int a, int b)
{
    return a+b;
}
