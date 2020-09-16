## C++中的new和delete操作符内存配置的方式

* new:先配置内存，然后构造对象
* delete:先析构对象，然后释放内存

## SGI STL（也就是标准STL采用的版本）的内存配置

* 将内存配置、内存释放与构造、析构分开
    即一个由：<stl_alloc.h>负责
      一个由：<stl_construct.h>负责
    这个两个头文件都是被放在<memory>中

    * STL alloctor中 alloc::allocate负责内存配置
    * STL alloctor中 alloc::deallocate负责内存释放
    * STL alloc::construct负责对象构造
    * STL alloc::destroy负责对象析构

## SGI对空间的配置与释放的设计哲学：
* 向system heap要求空间；
  * SGI底层用的malloc和free完成内存的配置和释放。
* 考虑多线程状态；
* 考虑内存不足时的应变措施；
    使用new_handler(C++语言自带)当内存不足时，让用户自定义函数去处理
* **考虑过多“小型区块”可能造成的内存碎片问题**
    * SGI设计了双层级配置器，第一级直接使用malloc和free
    * 第二级配置器则视情况采用不同的策略：
      * 当小于128bytes时，视为过小，为了减小额外负担而采用更复杂的memory pool方式，不再求助于第一级适配器


### SGI STL第一级配置器
 template<int inst>
 class _malloc_alloc_template{...};
 第一级配置器以malloc() , free()， realloc()函数实现，并模拟了C++中的new_handler因为它没有使用new所以C++中的new_handler机制并不能使用。

 ### SGI STL第二级配置器（也是默认调用的配置器）
 template<bool thread, int inst>
 class _default_alloc_template{...};
 其中：
 * 维护16个自由链表，负责小型区域的次配置能力，内存池以malloc()配置而得，如果内存不足，转而调用第一级配置器
 * 如果申请得内存大于128bytes，就转而调用第一级配置器

 最终typedef _default_alloc_template alloc;就是日常我们使用的默认的分配器
 ### SGI STL第二级分配器当中引发的两个问题

1. 在进行dealloc的时候，并没有将内存还给操作系统，所以占用的内存一直保持在最高峰值
2. 没有检查传入的要释放的指针的类型和大小，如果传入的大小不是8的倍数会引发危险

### 二级配置器的源代码

```
template <bool threads, int inst>
class __default_alloc_template {

private:
  // Really we should use static const int x = N
  // instead of enum { x = N }, but few compilers accept the former.
#if ! (defined(__SUNPRO_CC) || defined(__GNUC__))
    enum {_ALIGN = 8};
    enum {_MAX_BYTES = 128};
    enum {_NFREELISTS = 16}; // _MAX_BYTES/_ALIGN
# endif
  static size_t
  _S_round_up(size_t __bytes) 
    { return (((__bytes) + (size_t) _ALIGN-1) & ~((size_t) _ALIGN - 1)); }

__PRIVATE:
  union _Obj {
        union _Obj* _M_free_list_link;
        char _M_client_data[1];    /* The client sees this.        */
  };
private:
# if defined(__SUNPRO_CC) || defined(__GNUC__) || defined(__HP_aCC)
    static _Obj* __STL_VOLATILE _S_free_list[]; 
        // Specifying a size results in duplicate def for 4.1
# else
    static _Obj* __STL_VOLATILE _S_free_list[_NFREELISTS]; 
# endif
  static  size_t _S_freelist_index(size_t __bytes) {
        return (((__bytes) + (size_t)_ALIGN-1)/(size_t)_ALIGN - 1);
  }

  // Returns an object of size __n, and optionally adds to size __n free list.
  static void* _S_refill(size_t __n);
  // Allocates a chunk for nobjs of size size.  nobjs may be reduced
  // if it is inconvenient to allocate the requested number.
  static char* _S_chunk_alloc(size_t __size, int& __nobjs);

  // Chunk allocation state.
  static char* _S_start_free;
  static char* _S_end_free;
  static size_t _S_heap_size;

# ifdef __STL_THREADS
    static _STL_mutex_lock _S_node_allocator_lock;
# endif

    // It would be nice to use _STL_auto_lock here.  But we
    // don't need the NULL check.  And we do need a test whether
    // threads have actually been started.
    class _Lock;
    friend class _Lock;
    class _Lock {
        public:
            _Lock() { __NODE_ALLOCATOR_LOCK; }
            ~_Lock() { __NODE_ALLOCATOR_UNLOCK; }
    };

public:

  /* __n must be > 0      */
  static void* allocate(size_t __n)
  {
    void* __ret = 0;

    if (__n > (size_t) _MAX_BYTES) {
      __ret = malloc_alloc::allocate(__n);
    }
    else {
      _Obj* __STL_VOLATILE* __my_free_list
          = _S_free_list + _S_freelist_index(__n);
      // Acquire the lock here with a constructor call.
      // This ensures that it is released in exit or during stack
      // unwinding.
#     ifndef _NOTHREADS
      /*REFERENCED*/
      _Lock __lock_instance;
#     endif
      _Obj* __RESTRICT __result = *__my_free_list;
      if (__result == 0)
        __ret = _S_refill(_S_round_up(__n));
      else {
        *__my_free_list = __result -> _M_free_list_link;
        __ret = __result;
      }
    }

    return __ret;
  };

  /* __p may not be 0 */
  static void deallocate(void* __p, size_t __n)
  {
    if (__n > (size_t) _MAX_BYTES)
      malloc_alloc::deallocate(__p, __n);
    else {
      _Obj* __STL_VOLATILE*  __my_free_list
          = _S_free_list + _S_freelist_index(__n);
      _Obj* __q = (_Obj*)__p;

      // acquire lock
#       ifndef _NOTHREADS
      /*REFERENCED*/
      _Lock __lock_instance;
#       endif /* _NOTHREADS */
      __q -> _M_free_list_link = *__my_free_list;
      *__my_free_list = __q;
      // lock is released here
    }
  }

  static void* reallocate(void* __p, size_t __old_sz, size_t __new_sz);

} 
```



 

