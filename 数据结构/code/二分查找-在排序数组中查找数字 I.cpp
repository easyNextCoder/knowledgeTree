#include <iostream>
#include <vector>
#include <queue>
using namespace std;


class Solution {
public:
	int binary_search_n(vector<int>&nums, int target, int start, int end){
		/*
        while(start<end){
        	cout<<start<<end<<endl;
            int mid = (start+end)/2;
            if(target > nums[mid]){
                binary_search_n(nums, target, mid+1, end);
                start = mid+1;
            }else if(target < nums[mid]){
                binary_search_n(nums, target, start, mid-1);
                end = mid-1;

            }else{
            	cout<<"here"<<endl;
                int count = 0;
                int i = mid;
				while(i>=0 && nums[i--] == target){
                    count++;
                }
                i = (start+end)/2+1;
                while(i<nums.size()&&nums[i++] == target){
                    count++;
                }
                return count;
            }
        } 
		*/
		int left =start;
    int right=end;
    while(left<=right)
    {
    	cout<<left<<right<<endl;
        int mid =(left+right)/2;
        if(target==nums[mid])
        {
        	 cout<<"here"<<endl;
            return mid;
        }
        if(target>nums[mid])
        {
            left=mid+1;
        }
        else
        {
            right =mid-1;
        }
    }
    return -1;//Î´ÕÒµ½x  
        return 0;
    }
	/*
    int binary_search_n(vector<int>&nums, int target, int start, int end){
        if(end == start){
            return 0;
        }else{
            int mid_value = nums[(start+end)/2];
            if(mid_value == target){
                
                int count = 0;
                int i = (start+end)/2;
				while(i>=0 && nums[i--] == target){
                    count++;
                }
                i = (start+end)/2+1;
                while(i<nums.size()&&nums[i++] == target){
                    count++;
                }
                return count;
            }else if(mid_value<target){
            	cout<<(start+end)/2+1<<end<<endl;
                return binary_search_n(nums, target, (start+end)/2+1, end);
            }else{
            	
                return binary_search_n(nums, target, start, (start+end)/2-1);
            }
        }
    }
    */
    int search(vector<int>& nums, int target) {
        //binary search
        if(nums.size() == 0){
            return 0;
        }
        int count = 0;
        while(count < nums.size() && nums[count] == target){count++;}
        if(count > 0)
            return count;
        count = 0;
		int count_i = nums.size()-1;
		while(count_i >=0 && nums[count_i--] == target ){count++;}
        if(count >0)
            return count;
        cout<<"binary_search."<<endl;
        return binary_search_n(nums, target, 0, nums.size()-1);
    }
};

class A {
public:
    A();
};
class B {
public:
    explicit B(int a){
	};

};
void dosomething(B obj) {
    ;
}

int main(){
	priority_queue<int> max_queue;
	max_queue.push(1);
	max_queue.push(10);

	Solution solution;
	vector<int> vec = {0,1,2,3,4,4,4};
	cout<<"the result is:"<<solution.search(vec, 2);
	dosomething(14);
	return 0;
}
