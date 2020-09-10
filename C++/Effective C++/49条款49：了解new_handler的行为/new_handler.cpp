//当我们的像系统申请内存过大，系统无法提供时
//系统就会调用new_handler,这个new_handler的
//系统默认配置就是输出无法申请内存并abort
//实际上我们可以重新设置new_handler来自己处理申请
//大量内存的情况

/*
    namespace std{
        typedef void (*new_handler)();
        //安装用户自定义的handler处理函数，并返回旧的handler处理函数
        new_handler set_new_handler(new_handler p) throw();

    }

 */



#include <iostream>

using namespace std;

void outOfMem()
{
    std::cerr<<"Unable to satisfy request for memory\n";
    std::abort();
}

class big{
public:
    int arr[1000];
};

int main()
{
    std::set_new_handler(outOfMem);
    big* pBigArr = new big[100000000];
    cout<<"end."<<endl;

}
