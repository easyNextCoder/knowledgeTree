## 以对象管理资源

作者的含义就是使用类似于只能指针的对象来管理资源，这样能够做到最终能够得到释放，而不用最后自己去delete，而且有可能还无法delete掉，因为delete之前有可能发生return或者抛出异常等其他行为。

* STL容器要求其元素发挥正常的复制行为。（能复制成两份，而不是复制得到一份之后，被复制者变为空）
* shared_ptr都在其析构函数内做delete而不是delete[]动作，这意味着在动态分配而得到的array身上使用shared_ptr是个馊主意，但是其仍然能够通过编译

## 条款15：在资源管理类中提供对原始资源的访问

* **疑问点：资源管理类与资源之间的隐式转化写法**

## 条款21：必须返回对象时，不要返回其引用

## 条款23：尽量使用非成员、非友元函数来替换成员函数

* 越多的东西被封装，越少的人可以看到它。而越少的人看到它，我们就有越大的弹性去改变它。
* 将非成员函数和类放入到一个命名空间中，**命名空间和类并不相同，命名空间可以跨越多个文件而类不行。**

## 条款25：考虑写出一个不抛出异常的swap函数

* 在这里使用了模板的特例化，模板特例化可以用来优化函数对某些特定类的操作

	namespace std{
		template<>//这一行表示swap的一个全特化版本，之后的Widget表示对Widget这个版本的特殊优化代码
		void swap<Widget>(Widget&a, Widget& b)
		{
			swap(a.pImpl, b.Impl);
		}

* C++只允许对class templates偏特化，在function templates身上偏特化是行不通的

## 条款25：考虑写出一个不抛出异常的swap函数

* 对std空间中的template进行特例化是鼓励的，但是不要尝试往空间中加入新的东西。


# 6 继承与面向对象设计

## 条款32：确定你的public继承塑造出is-a的关系

* virtual函数意味着接口必须被继承
* non-virtual函数意味着接口和实现都必须被继承

## 条款33：
* 不同作用域中相同名称不同类型变量之间会发生遮掩，内层变量会遮掩外层变量。
* 对于类的继承来讲，继承类中的中与基类名字相同但是类型不同的函数仍然会发生遮掩情况，继承类的函数会遮掩基类的函数。

## 条款35

function<>中只要加入相同的函数签名，就可以代表任何运行函数，函数对象等等。

## 条款37

静态绑定叫前期绑定

动态绑定叫后期绑定（运行时绑定）；

virtual函数动态绑定而来，意思是调用一个virtual函数时，究竟调用哪一份函数实现代码，取决于发出调用的那个对象的动态类型。

**vitual函数是动态绑定的，而缺省参数值确实静态绑定的**意思是：可能会在调用一个定义与derived class内的virtual函数的同时，却使用base class为它所指定的缺省参数值。对于引用这个问题仍然存在。

## 条款38

* 中间有部分语句自己并不理解
* 使用模板的方法做一个底层是list的set

## 条款39：明智而审慎的使用private继承

* **如果class之间的继承关系是private，编译器不会自动将一个derrived class对象转换为一个base class对象**

* 由private base class继承而来的所有成员，在derived class 中都会变成private属性，纵使它们在base class中原本是protected或public属性**？？**。

* private继承意味着implemented-in-terms-of（根据某物实现出）。如果你让class D以private形式继承class B，你的用意是为了采用class B内已经准备妥当的某些特性，**不是因为B对象和D对象存在任何观念上的关系。**

* 想尽方法让继承类有避免重写virtual函数的能力，C++11中已经有了final字段。

* **即使是对于private继承，仍然可以在集成类中重新定义基类的私有的virtual函数，即使你不能访问他**。

## 条款40：小心使用多重继承

* 多重继承的确有正当用途。其中一个情节涉及"public继承某个Interface class"和"private 继承某个协助实现class"的两相组合。

* **注意：仔细思考，private继承的含义是表述，继承类以何种方式向外对客户展示基类成员**

## 条款51：编写new和delete时需固守常规

* 针对class X而设计的operator new,其行为很典型地只为大小刚好为sizeof(X)的对象设计，然而一旦被继承下去，有可能base class的operator new被调用用以分配derived class对象。

* 设计new时候的准则
	* operator new应该内含一个无穷循环，并在其中尝试分配内存，如果它无法满足内存需求，就该调用new-handler。
	* 它也应该有能力处理0 bytes申请。class专属版本则还应该处理“比正确大小更大的错误申请”
* 在设计operator new[]需要遵守的规则：
	* base class中的operator new[]有可能经由继承被调用，将内存分配给“元素为derived class对象”的array使用，而derived class对象通常比base class对象大。而且动态分配的arrays可能包含额外空间用来存放元素个数。

* 写delete时需要注意的准则：
	* 处理好delete NULL
	* 使用与new对应的delete.

* 关于placement new 和placement delete
	* placement delete只有在“伴随placement new”调用而触发的构造函数出现异常时才会被调用。对着一个指针施行delete绝对不会导致调用placement delete（即使这个内存是通过placement new申请的）。
	* 为用户定义的placement new 会掩盖系统中自带的new函数，所以用户再使用new申请内存的时候会出错（**用相应的代码测试**）