#include <iostream>
#include <vector>
#include <algorithm>
#include <set>

using namespace std;

int main()
{
    vector<vector<int>> infirst;
    vector<vector<int>> regin;
    int start, end, id;
    int index = 0;
    while(cin>>start>>end>>id)
    {
        regin.push_back({start, end, id});
    }
    vector<vector<int>> out;
    
    sort(regin.begin(), regin.end(), [](vector<int>& a, vector<int>& b){
        return a[0]<b[0];
    });
    
    for(;;)
    {
        int first = -1;
        int count = 0;
        for(auto re:regin)
        {
            if(re[0]<=re[1])
            {
                first = re[0];
                break;
            }
            count++;
        }
        if(count == regin.size())
            break;//已经处理完毕
        else{
            //先建立起来的第一个区间
            out.push_back({first, first});
            
            set<int> tmpId;
            int diedLinesCount = 0;
            int living = 0;
            //一共n行，死亡行数一变化，就要新开一个区间重新统计
            
            for(int i = 0; i<regin.size(); i++)
            {
                if(regin[i][0] > regin[i][1])
                {
                    diedLinesCount++;
                }
                if(regin[i][0] <= regin[i][1] && regin[i][0] == first)
                {
                    living++;
                }
            }
            int nowDiedLinesCount = diedLinesCount;
            int nowliving = living;
            do
            {//死亡行数如果没有变化就一直在本区间前进
                for(int i = 0; i<regin.size(); i++)
                {
                    if(regin[i][0] <= regin[i][1] && regin[i][0] == first)
                    {
                        regin[i][0]++;
                        tmpId.insert(regin[i][2]);
                    }
                }
                
                nowDiedLinesCount = 0;
                nowliving = 0;
                first++;
                for(int i = 0; i<regin.size(); i++)
                {
                    if(regin[i][0] > regin[i][1])
                    {
                        nowDiedLinesCount++;
                    }
                    if(regin[i][0] <= regin[i][1] && regin[i][0] == first)
                    {
                        nowliving++;
                    }
                }
                if(nowDiedLinesCount == diedLinesCount && nowliving == living)
                    out.back()[1]++;
            }while(nowDiedLinesCount == diedLinesCount && nowliving == living);
            //死亡行和存活行有一个发生变动，就要重新开一个新的空间
            
            //死亡行数一发生变化，就要重新开一个区间重新统计  
            //-> oriD > dete
            for(auto item:tmpId)
            {
                out.back().push_back(item);
            }
            //发生了变化就应该回到头去重新寻找新的first然后重新开始,同时清空setid
            tmpId.clear();
            
        }

    }
    for(auto item:out)
    {

        for(int i = 0; i<item.size(); i++)
        {
            if(i == item.size()-1)
                cout<<item[i];
            else
                cout<<item[i]<<" ";
        }
        cout<<endl;
    }
    
    return 0;
}

/*
用例:
2 2 1
1 2 2

对应输出应该为:

1 1 2
2 2 1 2

你的输出为:

1 1 2
2 2 1 2
*/


链接：https://www.nowcoder.com/questionTerminal/f9daa556baa84997b3482755aaf63ec5?f=discussion
来源：牛客网

// 调输入调了2个小时，这道题的实际输入一堆问题。样例中没有给n,直接给的是版本号
// 另外最后的样例中通配符可能出现不只是*的情况，可能有多余的空格之类的，所以还需要catch异常
// 由于题目的样例中只有输出0的情况，c++的cin或是scanf即使是没有读对也不会报错，所以拿c++写的一通乱算算出来都是0，正好就对了 (-。 -)
// 更新，优化了运行速度


链接：https://www.nowcoder.com/questionTerminal/01cb04dc53f54625834f2a86c519dce9?f=discussion
来源：牛客网

//这道题目可以定义一个结构体用于存储每个分段的起始时间，截止时间，价格，然后对所
//有相邻且价格一致的区间进行merge，这也是题目本身想要考察的。但是这样的做法不适
//合考试时候时间有限的情况下，虽然单纯的merge并不复杂，但是注意题目中提到了酒店
//的价格可能不一致，如果不一致，按后面的价格为准，如果后面记录的价格和前面的记录
//不一样，这就可能导致原来的一个区间段更新价格后分裂成为两个或者三个，略显复杂。
//所以可以直接用一个数组存储每天的价格，这样做的好处就是即使前后数据不一致，后面
//的会直接覆盖前面的完成价格更新，缺点就是需要一个数组来记录每天的价格，空间复杂
//度较高，对于这个实际问题而言，不会出现特别多的天数，所以不必考虑大数问题，当然
//也所幸内存够用。特别需要注意输出格式要满足题目要求。