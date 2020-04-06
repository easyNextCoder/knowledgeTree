  #include <iostream>
 #include <vector>
 #include <string>

using namespace std;

class Solution {
public:
    void find(vector<vector<int>>& record, int row, int col, int k){
        
        if(row < 0 || col <0 || row>=(int)record.size() || col>=(int)record[0].size())
            return;
        int sum = 0;
        string srow = to_string(row);
        string scol = to_string(col);
        
        for(size_t i = 0; i<srow.size(); i++){
            sum+=srow[i] - '0';
        }
        for(size_t i = 0; i<scol.size(); i++){
            sum+=scol[i] - '0';
        }
        if(sum > k || record[row][col] == 1){
            return;
        }else{
            record[row][col] = 1;
            find(record, row, col+1, k);
            find(record, row, col-1, k);
            find(record, row+1, col, k);
            find(record, row-1, col, k);
        }
    }
    int movingCount(int threshold, int rows, int cols)
    {
        vector<vector<int>> record;
        for(int i = 0; i<rows; i++){
        	record.push_back(*(new vector<int>()));
            for(int j = 0; j<cols; j++){
                record[i].push_back(0);
            }
        }
        
        find(record, 0, 0, threshold);
        int val = 0;
        for(auto item:record){
            for(auto innerItem:item)
                val+=innerItem;
        }
        return val;
    }
};

int main(){
	Solution* solution = new Solution();
	cout<<solution->movingCount(18,100,100)<<endl;
	return 0;
}
