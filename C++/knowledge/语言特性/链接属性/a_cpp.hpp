
static int a = 10;
//int a = 10;
/*
    如果不加上static整个Cpp系统运行就会有问题
 */

int sum(int a, int b);
int sum2a(void);
int getVarb();
/*
    在C++文件系统中：
        没有这几行的话整个cpp软件系统就无法运行
    在C文件系统中：
        没有这几行照样能运行
    由此可见：
        在C++系统中是默认内部链接属性的，而在C系统中是默认外部链接属性的
 */