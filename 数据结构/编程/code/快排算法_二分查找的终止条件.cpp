#include <iostream>
#include <vector>

using namespace std;
//注意默写一下quick_sort和binary search
//都是最终的终止条件的判定；快速排序的终止条件的判定,二分法搜索的边界条件的判定 
void insert_sort(vector<int>&vec, int first, int last){
	 
	for(int i = first; i<last; i++){
		int tmp_min = vec[i];
		int tmp_min_index = i;
		for(int j = i+1; j<=last; j++){
			if(tmp_min > vec[j]){
				tmp_min = vec[j];
				tmp_min_index = j;
			}
		}
		
		swap(vec[i], vec[tmp_min_index]);
	}
}

int get_provit(vector<int>&vec, int first, int last){
	
	int mid = (first + last)/2;
	if(vec[first] > vec[last])
		swap(vec[first], vec[last]);
	if(vec[mid]<vec[first])
		swap(vec[first], vec[mid]);
	if(vec[mid]>vec[last])
		swap(vec[mid], vec[last]);
	return vec[mid];
}

void quick_sort(vector<int>&vec, int first, int last){
	if(last - first < 3){
		insert_sort(vec, first, last);
		return ;
	}
	int provit = get_provit(vec, first, last);
	swap(vec[(first+last)/2], vec[last-1]);
	int i = first;
	int j = last - 2;
	for(;;){
		while(vec[++i]<provit){}
		while(vec[--j]>provit){}
		if(i<j){
			swap(vec[i], vec[j]);
		}else{
			swap(vec[i], vec[last-1]);
			break;
		}
	}
	
	quick_sort(vec, first, i-1);
	quick_sort(vec, i+1, last);
}

int binary_search(vector<int>&vec, int target){
	int i = 0, j = vec.size() - 1;
	int mid = -1;
	while(i<=j){
		mid = (i+j)/2;
		if(vec[mid] < target){
			i = mid + 1;
		}else if(vec[mid] == target){
			return mid;
		}else{
			j = mid - 1;
		}
	}
	return -1;
}

int main(){
	
	vector<int> vec = {6,5,4,5,6,4,6,9,10,67,88,90,67,10};
	quick_sort(vec, 0, vec.size()-1);
	for (auto item : vec)
        cout << item << endl;
        
    int index = binary_search(vec, 67);
    cout<<"the index is: "<<index<<endl;
	return 0;
}
