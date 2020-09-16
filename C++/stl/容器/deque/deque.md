## deque性质

deque是一个分段连续的数据结构，连续是假象， 分段是事实

## deque函数

1. insert函数
iterator insert(iterator position, const value_type& x)
{
    if(插入在头部)

    else if(插入在尾部)

    else
        insert_aux();//辅助函数
}

iterator insert_aux(iterator position, const value_type& x)
{
    //判断插入点距离头端近还是尾端近，哪一端近移动哪一端
}

