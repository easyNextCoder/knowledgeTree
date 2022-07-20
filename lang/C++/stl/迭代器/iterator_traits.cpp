#include <iostream>

using namespace std;

template <class T>
struct iterator_traits<const T*>
{
	typedef T value_type;
}

int main()
{
	
	return 0;	
} 
