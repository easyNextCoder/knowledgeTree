const int N = 520;
int f[8][520];

class Solution {
public:
    /*
    vector<vector<char>> cseats;
    vector<vector<int>> state;
    vector<vector<int>> path;
    int maxOut = 0;
    void getMax(int row, int col, int tmp)
    {
        if(row == cseats.size()-1 && col == cseats.front().size())
        {
            maxOut = max(maxOut, tmp);
            return ;
        }
        
        if(col == cseats.front().size())
        {
            getMax(row+1, 0, tmp);
        }else{
            //先看座位能不能坐，如果不能坐就不坐
            if(cseats[row][col] == '#')
            {
                getMax(row, col+1, tmp);
            }else{
                //正常的点要审视周围能不能看到
                int unsafe = 0;

                if( row-1 < 0 || col-1 < 0 )
                {
                    if(col-1>=0 && state[row][col-1])unsafe++;
                }
                else{
                    if(state[row][col-1])unsafe++;
                    if(state[row-1][col-1])unsafe++;
                } 

                if(row-1 < 0 || col+1 >= cseats.size()){
                    ;//if(col+1<cseats.size() && state[row][col+1])unsafe++;
                }
                else{
                    //if(state[row][col+1])unsafe++;
                    if(state[row-1][col+1])unsafe++;
                }

                if(unsafe > 0)
                {//这个位置不能坐
                    getMax(row, col+1, tmp);
                }
                else{

                    //不坐
                    getMax(row, col+1, tmp);
                    //坐
                    state[row][col] = 1;
                    path.push_back({row, col});
                    getMax(row, col+1, tmp+1);
                    path.pop_back();
                    state[row][col] = 0;
                }
            }
        }
    }
    
    long long int vstate = 0;
    long long int badState = 0;
    int width = 0;
    int len = 0;

    int getMax1(int index, long long int vstate, int tmp)
    {
        if(index == len)
        {
            maxOut = max(maxOut, tmp);
            return maxOut;
        }
        else{
            if(badState&((long long int)1<<index))
            {//这个位置不能坐人
                getMax1(index+1, vstate, tmp);
            }else{
                int lu = index-width-1;
                int ru = index-width+1;
                int l = index-1;
                int unsafe = 0;
                if(lu >= 0 && (vstate & (long long int)1<<lu))unsafe++;
                if(ru >= 0 && (vstate & (long long int)1<<ru))unsafe++;
                if(l >= 0 && (vstate & (long long int)1<<l))unsafe++;
                if(unsafe)
                {//这个位置不能用
                    getMax1(index+1, vstate, tmp);
                }else{
                    //不占用这个位置
                    getMax1(index+1, vstate, tmp);
                    //占用这个位置
                    getMax1(index+1, vstate|((long long int)1<<index), tmp+1);
                }
            }
            
        }
        return maxOut;
    }
    */

    int mem[8][1<<8];
    vector<int> compressed;
    int memGet(int index, int state, int width)
    {
        if(index >= compressed.size())return 0;
        if(mem[index][state] != -1)return mem[index][state];
        int out = 0;
        for(int scheme = 0; scheme<(1<<width); ++scheme)
        {
            //学生不可以坐，或者坐上就连续的位置
            if(scheme & ~state || scheme & (scheme<<1))
                continue;
            //屏蔽后面的某些座位
            int ans = 0;
            for(int i = 0; i<width; i++)
                if(scheme & 1<<i)ans++;

            if(index+1 == compressed.size())
                out = max(out, ans);
            else{
                unsigned char next_state = compressed[index+1];
                next_state &= ~(scheme<<1);//右边下脚的去除
                next_state &= ~(scheme>>1);//左下角的去除
                out = max(out, ans+memGet(index+1, next_state, width));    
            }
        }
        mem[index][state] = out;
        return out;
    }

    int maxStudents(vector<vector<char>>& seats) {
        
        //采用记忆化递归的方法
        
        for(auto vchar:seats)
        {
            int count = 0;
            int ans = 0;
            for(auto ch:vchar)
            {
                if(ch == '.')
                {
                    ans |= (int)1<<count;
                }
                count++;
            }
            compressed.push_back(ans);
        }
        memset(mem, -1, sizeof(mem));
        return memGet(0, compressed[0], seats.front().size());
        /*
            //递归的方式
            maxOut = 0;
            cseats = seats;
            vector<vector<int>> cstate(seats.size()+1, vector<int>(seats.front().size()+1, 0));
            state = cstate;
            getMax(0, 0, 0);
        */

        /*
            //采用状态压缩的方法->定义一个long long int 型变量就可以表示全部的状态
            width = seats.front().size();
            len = seats.size()*seats.front().size();
            
            int count = 0;
            for(auto vchar:seats)
            {
                for(auto ch:vchar)
                {
                    if(ch == '#')
                    {
                        badState |= (long long int)1<<count;
                    }
                    count++;
                }
            }
            getMax1(0, vstate, 0);
         */

        
            
            

         

        /*
            采用动态规划的方法?
         */
        
        

        //return maxOut;
    }
};