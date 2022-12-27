package xsort

//#include <iostream>
//#include <vector>
//#include <random>
//
//using namespace std;
//
//void merge_sort(vector<int>& vec, int first, int last)
//{
//    if(last - first + 1 < 3)
//    {
//        if(last == first)
//            return ;
//        else if(vec[first] > vec[last]){
//            swap(vec[first], vec[last]);
//        }
//        return ;
//    }else{
//        cout<<first<<":"<<last<<endl;
//        merge_sort(vec, first, first + (last-first)/2);
//        merge_sort(vec, first + (last-first)/2+1, last);
//        int first1 = first;
//        int last1 = first + (last-first)/2;
//        int first2 = last1+1;
//        int last2 = last;
//        vector<int> tmp(last - first + 1);
//        int count = 0;
//        while(first1 <= last1 || first2 <= last2)
//        {
//            if(first1 <= last1 && first2 <= last2)
//            {
//                if(vec[first1] < vec[first2])
//                {
//                    tmp[count++] = vec[first1++];
//                }else{
//                    tmp[count++] = vec[first2++];
//                }
//            }else{
//                if(first1 <= last1){
//                    tmp[count++] = vec[first1++];
//                }else{
//                    tmp[count++] = vec[first2++];
//                }
//            }
//        }
//        for(int i = first; i<=last; i++)
//            cout<<tmp[i]<<" ";
//        cout<<endl;
//        count = 0;
//        for(int i = first; i <= last; i++)
//        {
//            vec[i] = tmp[count++];
//        }
//    }
//}
//
//int main()
//{
//
//    vector<int> vec = {9,8,7,6,5,4,3,2,1};
//    /*
//    default_random_engine e;
//    int count = 1000;
//    while(count--)
//    {
//        vec.push_back(e()%(10000));
//    }
//    cout<<"original array is:"<<endl;
//    for(auto item:vec)
//    {
//        cout<<item<<" ";
//    }
//    */
//    merge_sort(vec, 0, vec.size()-1);
//    cout<<"sorted array is:"<<endl;
//    for(auto item:vec)
//    {
//        cout<<item<<" ";
//    }
//    cout<<endl;
//    return 0;
//}
