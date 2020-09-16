* 成员函数是const意味着什么？有两个流派
  * bitwise constness(physical constness)

    成员函数只有在不更改对象内的任何成员变量（static除外）时才可以说是cosnt的

  * logical constness

    一个const成员函数可以修改它所处理的对象内的某些bits，但是只有在客户端侦测不出的情况下才能如此。

* 如何解决想要在一个const对象中修改对象成员？

    在对象的成员前面加上mutable