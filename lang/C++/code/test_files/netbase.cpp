#include <iostream>
#include <vector>
#include <math.h>
#include <algorithm>

using namespace std;
int cal_length(vector<vector<int>>& con, int x, int y, int depth, ){
	int left_up_px = x-1>=0?x-1:0;
	int left_up_py = y-1>=0>y-1:0;
	int right_dn_px = x+1<con.size()?x+1:con.size()-1;
	int right_dn_py = y+1<con.size()?y+1:con.size()-1;
	//上横行
	vector<pair<double,int>> tmp;
	for(int i = left_up_py; i<=right_dn_py; i++){
		double dis = sqr(double((x-left_up_px)*(x-left_up_px)) + 
						 double((y-i)*(y-i)))
		int thing = con[left_up_px][i];
		if(thing == 0){
			continue;
		}else{
			tmp.push_back(make_pair(dis, thing));
		}
	}
	sort(tmp.begin(), tmp.end(), [](pair<double,int>&item1, pair<double,int>&item2){
		return item.first<item2.first;
	})
	for(auto item:tmp){
		
	}
	//下横行
	
	//左竖行
	
	//右竖行 
	return 0;
}
int main(){
	int n = 0;
	cin>>n;
	cin.ignore();
	while(n--){
		vector<vector<int>> con;
		int size = 0;
		int knife = 0;
		cin>>size>>knife;
		for(int i = 0; i<size; i++){
			vector<int> in_con;
			int value = -1;
			for(int j = 0; j<size; j++){
				cin>>value;
				in_con.push_back(value);
			}
			con.push_back(in_con);
		} 
		int x = 0, y = 0;
		cin>>x>>y;
		cin.ignore();
		/*
		cout<<"output begin()."<<endl;
		for(auto item:con){
			for(auto in_item:item){
				cout<<in_item<<"-";
			}
			cout<<endl;
		}
		*/
		cal_length(con, x, y, 1);
	}
	
	return 0;
}
