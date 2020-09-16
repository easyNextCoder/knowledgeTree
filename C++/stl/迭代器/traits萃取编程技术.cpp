#include <iostream>
#include <algorithm>
#include <vector>
//https://blog.csdn.net/wutao1530663/article/details/64922389
namespace xstd{

    template<class Category,class T,class Distance = ptrdiff_t,class Pointer=T*,class Reference=T&>
    class iterator{
        typedef Category iterator_category;
        typedef T        value_type;
        typedef Distance difference_type;
        typedef Pointer  pointer;
        typedef Reference reference;
    };

    struct input_iterator_tag{};
    struct output_iterator_tag{};
    struct forward_iterator_tag:public input_iterator_tag{};
    struct bidirectional_iterator_tag:public forward_iterator_tag{};
    struct random_access_iterator_tag:public bidirectional_iterator_tag{};


    template <class Iterator>
    struct iterator_traits{
        typedef typename Iterator::iterator_category iterator_category;
        typedef typename Iterator::value_type value_type;
        typedef typename Iterator::pointer pointer;
        typedef typename Iterator::reference reference;
        typedef typename Iterator::difference_type difference_type;

    };
    template <class T>
    struct iterator_traits<T*>{
        typedef std::random_access_iterator_tag iterator_category;
        typedef T value_type;
        typedef T* pointer;
        typedef T& reference;
        typedef ptrdiff_t difference_type;
    };

    template <class T>
    struct iterator_traits<const T*>{
        typedef std::random_access_iterator_tag iterator_category;
        typedef T value_type;
        typedef const T* pointer;
        typedef const T& reference;
        typedef ptrdiff_t difference_type;
    };

    template<class InputIterator>
    inline typename xstd::iterator_traits<InputIterator>::difference_type __distance(InputIterator first, InputIterator last, std::input_iterator_tag){
        typename xstd::iterator_traits<InputIterator>::difference_type n = 0;
        while (first != last){
            ++first; ++n;
        }
        return n;
    }

    template<class InputIterator>
    inline typename xstd::iterator_traits<InputIterator>::difference_type \
        __distance(InputIterator first, InputIterator last, xstd::random_access_iterator_tag){
            return last - first;
    }

    template<class InputIterator>
    inline typename xstd::iterator_traits<InputIterator>::difference_type distance(InputIterator first, InputIterator last){
        std::cout<<"using xstd::distance()"<<std::endl;
        return __distance(first, last, typename xstd::iterator_traits<InputIterator>::iterator_category());
    }
    

};




template<class Item>
class vecIter{
    Item *ptr;
public:
    typedef std::random_access_iterator_tag iterator_category;
    typedef Item value_type;
    typedef Item* pointer;
    typedef Item& reference;
    typedef std::ptrdiff_t difference_type;
public:
    vecIter(Item *p = 0) :ptr(p){}
    Item& operator*()const{
        return *ptr;
    }
    Item* operator->()const{
        return ptr;
    }
    //pre
    vecIter& operator++(){
        ++ptr;
        return *this;
    }
    vecIter operator++(int){
        vecIter tmp = *this;
        ++*this;
        return tmp;
    }

    bool operator==(const vecIter &iter){
        return ptr == iter.ptr;
    }
    bool operator!=(const vecIter &iter){
        return !(*this == iter);
    }
};
int main(){

    int a[] = { 1, 2, 3, 4 };
    std::cout<<*vecIter<int>(a)<<std::endl;
    std::vector<int> vec(5, 99);
    int len = xstd::distance(vec.begin(), vec.end());
    std::cout<<len<<std::endl;
    //std::cout << std::accumulate(vecIter<int>(a), vecIter<int>(a + 4), 0);//输出 10

}

/*
int main(){
    int a[] = { 1, 2, 3, 4 };
    std::vector<int> vec{ 1, 2, 3, 4 };
    std::list<int> lis{ 1, 2, 3, 4 };
    std::cout<<"vec distance:"<<WT::distance(vec.begin(), vec.end())<<std::endl;
    std::cout << "list distance:" << WT::distance(lis.begin(), lis.end())<<std::endl;
    std::cout << "c-array distance:" << WT::distance(a,a + sizeof(a) / sizeof(*a)) << std::endl;
        //输出 vec distance:4
        //    list distance:4
        //    c-array distance:4
}
*/
