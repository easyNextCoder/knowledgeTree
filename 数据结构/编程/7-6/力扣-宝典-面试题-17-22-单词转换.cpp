#include <iostream>
#include <vector>
#include <string>
#include <set>

using namespace std;
class Solution {
public:
    bool isearch(string & a, const string & b)
    {
        int count = 0;
        for(int i = 0; i<a.size(); ++i)
            if(a[i] != b[i])count++;
        return count == 1;
    }
    //再进行递归是，多个参数引用传递并不会降低速度，反而使用全局变量涉及初始化更加损失速度
    bool dfs(string& start, string& end, vector<string>& wl, vector<string>& result, vector<bool>& visited)
    {
        if(start == end)return true;
        for(int i = 0; i<wl.size(); ++i)
        {
            if( visited[i] || !isearch(start, wl[i]))continue;
            visited[i] = true;
            result.push_back(wl[i]);
            bool ret = dfs(wl[i], end, wl, result, visited);
            if(ret)return true;
            //这里是个最奇葩的,本来是按照dfs的写法，应该把visited[i]恢复成false
            //但是恢复之后竟然会超时，不恢复却不会超时！此路不同，故不再行此路！
            //visited[i] = false;
            result.pop_back();
        }
        return false;
    }
    
    vector<string> findLadders(string beginWord, string endWord, vector<string>& wordList) {
        
        vector<string> result;   
        result.push_back(beginWord);
        vector<bool> visited(wordList.size(), false);
        int ret = dfs(beginWord, endWord, wordList, result, visited);
        if(ret)return result;
        else return {};
    }
};

int main()
{
    Solution so;
    vector<string> in = {"qa","ba","ca","fa","ga","ha","la","ma","na","pa","ra","ta","ya","yb","db","mb","nb","pb","rb","sb","tb","tc","sc","se","be","fe","ge","he","le","me","ne","re","ye","yo","co","go","ho","io","lo","mo","no","po","so","to","th","ph","rh","sh","uh","ur","ar","br","cr","er","fr","kr","lr","mr","or","sr","si","bi","ci","di","hi","li","mi","ni","pi","ti","tm","am","cm","fm","pm","sm","sn","an","ln","mn","mt","lt","pt","st","sq"};
    auto ret = so.findLadders("qa", "sq", in);
    for(auto item: ret)
    {
        cout<<item<<endl;
    }
    return 0;
}


    



