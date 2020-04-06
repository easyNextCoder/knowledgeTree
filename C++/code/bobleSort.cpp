#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void bobleSort(vector<int>& con){
	
	for(int i = con.size() - 1; i>0; i--){
		for(int j = 0; j<i; j++){
			if(con[j]>con[j+1]){
				/*
				int tmp = con[j];
				con[j] = con[j+1];
				con[j+1] = tmp;
				*/
				swap(con.at(j), con.at(j+1));
			}
		}
	}
}

int main(){
	
	vector<int> con = {7,6,5,4,3,3,2,1};
	bobleSort(con);
	for(auto item:con){
		cout<<item<<endl;
	}
	return 0;
}
