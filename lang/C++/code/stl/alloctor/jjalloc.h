#pragma once
#include <new>
#include <cstddef>//Ptrdiff_t ������ر�������ָ��Ĳ� signed���� ����Ϊ�� 
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
