//https://leetcode-cn.com/problems/largest-1-bordered-square/
class Solution {
public:
    int largest1BorderedSquare(vector<vector<int>>& grid) {
        
    for(int i = 0; i<grid.size(); i++)
    {
        int count = 1;
        for(int j = grid[i].size()-1; j>=0; j--)
        {
            if(grid[i][j] != 0)
            {
                grid[i][j] = count;
                count += 1;
            }else{
                count = 1;
            }
        }
    }
    for(int i = 0; i<grid[0].size(); i++)
    {
        int count = 1;
        
        for(int j = grid.size()-1; j>=0; j--)
        {
            if(grid[j][i] != 0)
            {
                grid[j][i] = min(count, grid[j][i]);
				//先统计出左上角的那个向下和向右能延申多远，找出那个最小值
                count += 1;
            }else{
                count = 1;
            }
        }
    }
       
        int maxout = 0;
        for(int i = 0; i<grid.size(); i++)
        {
            for(int j = 0; j<grid[i].size(); j++)
            {
                if(grid[i][j] != 0)
                {
                    int mini = min(i, j);
                    for(int k = 1; k<=mini; k++)
                    {    
                        if(grid[i-k][j] != 0 && grid[i][j-k] != 0)
                        {
                            if(grid[i-k][j-k] >= k)
                            {
                                maxout = max(maxout, k+1);
                            }
                        }else{
                            break;
                        }
                    }
                    maxout = max(1, maxout);
                }
            }
        }
        return maxout*maxout;
    }
};