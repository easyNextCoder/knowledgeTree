package xsort

//
//#include <iostream>
//#include <vector>
//
//using namespace std;
//
//
//void merge_sort(vector<int>&vec, int first, int last, vector<int>tmp)
//{
//	cout << first << ":" << last << endl;
//	if (last - first <= 0)return;
//	int mid = first + (last - first) / 2;
//	int ptr1 = first;
//	int ptr2 = mid + 1;
//	merge_sort(vec, first, mid, tmp);//这里的mid-1一定要与出口return 配合好
//	merge_sort(vec, mid + 1, last, tmp);
//	int tmp_ptr = first;
//	while (ptr1 <= mid || ptr2 <= last)
//	{
//		cout << "i" << endl;
//		if (ptr1 <= mid && ptr2 <= last)
//		{
//			if (vec[ptr1] < vec[ptr2])
//			{
//				tmp[tmp_ptr++] = vec[ptr1++];
//			}
//			else {
//				tmp[tmp_ptr++] = vec[ptr2++];
//			}
//
//		}
//		else if (ptr1 <= mid || ptr2 <= last) {
//			if (ptr1 <= mid) {
//				tmp[tmp_ptr++] = vec[ptr1++];
//			}
//			else {
//				tmp[tmp_ptr++] = vec[ptr2++];
//			}
//		}
//	}
//	for (int i = first; i <= last; i++)
//	{
//		cout << vec[i] << " ";
//		vec[i] = tmp[i];
//	}
//
//}
//
//int main()
//{
//	vector<int> vec = { 9,8,7,6,5,4,3,2,1 };
//	vector<int> tmp(vec.size());
//	cout << "here." << endl;
//	merge_sort(vec, 0, vec.size() - 1, tmp);
//	cout << "end merge_sort()." << endl;
//	cout << vec.size() << endl;
//	for (auto item : vec)
//	{
//		cout << item << " ";
//	}
//	cout << endl;
//
//	return 0;
//}
//
///*
//#include <iostream>
//#include <vector>
//
//using namespace std;
//
//
//void merge_sort(vector<int>vec, int first, int last, vector<int>tmp)
//{
//	cout<<first<<":"<<last<<endl;
//	if((first>=last))return;
//	int mid = first + (last - first)/2;
//	int ptr1 = first;
//	int ptr2 = mid;
//	merge_sort(vec, first, mid-1, tmp);//这里的mid-1一定要与出口return 配合好
//	merge_sort(vec, mid, last, tmp);
//	int tmp_ptr = first;
//	while(ptr1 < mid || ptr2 <= last)
//	{
//		cout<<"i"<<endl;
//		if(ptr1<mid && ptr2 <= last)
//		{
//			if(vec[ptr1] < vec[ptr2])
//			{
//				tmp[tmp_ptr++] = vec[ptr1++];
//			}else{
//				tmp[tmp_ptr++] = vec[ptr2++];
//			}
//
//		}else if(ptr1<mid || ptr2<=last){
//			if(ptr1<mid){
//				tmp[tmp_ptr++] = vec[ptr1++];
//			}else{
//				tmp[tmp_ptr++] = vec[ptr2++];
//			}
//		}
//	}
//	for(int i = first; i<=last; i++)
//	{
//		vec[i] = tmp[i];
//	}
//
//}
//
//int main()
//{
//	vector<int> vec = {9,8,7,6,5,4,3,2,1};
//	vector<int> tmp(vec.size());
//	cout<<"here."<<endl;
//	merge_sort(vec, 0, vec.size()-1, tmp);
//	cout<<vec.size()<<endl;
//	for(auto item:vec)
//	{
//		cout<<item<<" ";
//	}
//	cout<<endl;
//
//	return 0;
//}
//*/
