#include <iostream>
#include <functional>

using namespace std;
/*
	测试在定义类的时候什么时候需要类的声明什么时候需要定义
 */

class A;
class B {
public:
	int get() { ; }

private:
	shared_ptr<A>Impl;//不需要使用定义式，只需要使用声明式就可以
	//A require_define;//必须要使用到定义式，因为要知道内存的大小用来申请内存
};

/*
	测试IDE是否会检查inline函数的各种语法错误	
 */

/*
inline void test_inline() {
	a = b + c;
	return 1;
}
*/

/*
   测试是否能返回类内的成员函数指针或者类内成员变量的指针？
   测试结果：不能返回类内的动态成员的指针，因为这些指针就是一个偏移量（自己想了各种方法返均不成功）
			 但是能够返回类内的静态函数成员
 */

class A;
typedef  int (*funt)();
typedef int (A::* ref_funt)();
class A {

public:
	static int hit() {
		cout << "we are in hit." << endl;
		return 0;
	}
	
	int hit_ref() {
		cout << "we are in hit_ref." << endl;
		return 0;
	}
	int hit_test() {
		cout << "we are in hit_test." << endl;
		return 0;
	}
	funt get_hit_function(){
		auto rval = ((A::hit));
		return rval;
		//funt* p = (&(A::hit));
		//return rfun;//(&(A::hit));
		//return rval;
	}
	
	/*
	ref_funt get_hit_ref_function() {
		return &(A::hit_ref);
	}
	*/
private:
	
	int a = -1;
	//function<int()>f = []()->int {}
	//bind(&A::hit);
	//funt rfun = &(A::hit);
};

/*
	直接测试相当于c语言中的函数指针的使用情况，能正常使用
 */

int have() {

	return 0;
}
typedef int havet();


int main() {

	havet* another_have = have;
	cout << another_have() << endl;

	(A().get_hit_function())();

	/*//这个是无法执行的，因为在类内部取类内成员的地址是非法的
	ref_funt p = (A().get_hit_ref_function());
	A test;
	(test.*p)();
	*/

	ref_funt p2 = &(A::hit_ref);
	A test2;
	(test2.*(p2))();
	return 0;
}