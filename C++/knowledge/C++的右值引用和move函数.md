## 右值引用不同于右值

有名字的右值引用是一个左值
而没有名字的右值引用是一个右值，所以std：：move()函数就是一个右值

## move函数要解决的问题（将一个左值转换为对应的右值引用类型，move调用告诉编译器：**我们有一个左值，但我们希望像一个右值一样处理它**）
>c++ primer 473页
假设class X包含一个指向某资源的指针或句柄m_pResource。这里的资源指的是任何需要耗费一定的时间去构造、复制和销毁的东西，比如说以动态数组的形式管理一系列的元素的std::vector。逻辑上而言X的赋值操作符应该像下面这样：

```cpp
X& X::operator=(X const & rhs)
{
  // [...]
  // 销毁m_pResource指向的资源
  // 复制rhs.m_pResource所指的资源，并使m_pResource指向它
  // [...]
}
```

最后一行有如下的操作：

销毁x所持有的资源
复制foo返回的临时对象所拥有的资源
销毁临时对象，释放其资源
上面的过程是可行的，但是**更有效率的办法是直接交换x和临时对象中的资源指针，然后让临时对象的析构函数去销毁x原来拥有的资源**。换句话说，当赋值操作符的右边是右值的时候，我们希望赋值操作符被定义成下面这样：

```cpp
// [...]
// swap m_pResource and rhs.m_pResource
// [...]
```

move函数的出现就是为了解决这个问题。

### 没有使用move的swap库函数

```cpp
template<class T>
void swap(T& a, T& b)
{
  T tmp(a);
  a = b;
  b = tmp;
}
 
X a, b;
swap(a, b);
```

### 使用move函数之后能显著提升move函数的效率

这个函数只会做一件事就是：把它的参数转换为一个右值并返回。
现在swap使用了move语义，但是对于那些没有实现move语义的类型来说（没有针对右值引用重载拷贝构造函数和赋值操作符），新旧swap还是一样的

```cpp
template<class T>
void swap(T& a, T& b)
{
  T tmp(std::move(a));
  a = std::move(b);
  b = std::move(tmp);
}
 
X a, b;
swap(a, b);
```

