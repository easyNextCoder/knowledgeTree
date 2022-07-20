#pragma once
#include <new>
#include <cstddef>//Ptrdiff_t 机器相关保存两个指针的差 signed类型 可能为负 
#include <climits>
#include <iostream>

namespace JJ
{
	template <class T>
	inline T* _allocate(ptrdiff_t size, T*)
	{
		set_new_handler(0);
		T* tmp = (T*)(::operator new(size_t));
	}
}
