* 使用define A  B后，如果编译过程中出现关于A的错误会提示B，使得错误难以寻找

    因为一旦定义：
    const double AspectRatio  = 1.653,如果编译过程中出现错误。编译器会指出变量

* 使用类内的static常量来代替define可以定义类内的常量，也即限定了常量的作用域