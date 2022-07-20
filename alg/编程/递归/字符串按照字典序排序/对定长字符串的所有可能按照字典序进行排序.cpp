
//https://www.nowcoder.com/questionTerminal/6fc8716ee33e4cc59d58d7e18712094e
// 假定一种编码的编码范围是a ~ y的25个字母，从1位到4位的编码，如果我们把该编码按字典序排序，
// 形成一个数组如下： a, aa, aaa, aaaa, aaab, aaac, … …, b, ba, baa, baaa, baab, baac
//  … …, yyyw, yyyx, yyyy 其中a的Index为0，aa的Index为1，aaa的Index为2，以此类推。 
//  编写一个函数，输入是任意一个编码，输出这个编码对应的Index.

// 输入
// baca
// 输出
// 16331


#include <iostream>
#include <vector>
#include <string>

using namespace std;

int f[5] = { ((26*25+1)*25)*25+1 ,(26*25+1)*25+1,26*25+1,26,1};

int dfs(int start, vector<int>& ins)
{
    int out = 0;
    if(start == 4 || ins[start] == ' ')return 0;
    else{
        //cout<<"tmp:"<<(ins[start]-'a')*f[start+1]<<endl;
        //cout<<"tmp:"<<(ins[start]-'a')*f[start+1];
        out = (ins[start]-'a')*f[start+1]+dfs(start+1, ins)+1;
    }
    return out;
}

int main()
{
    vector<int> ins(4, ' ');
    string str;
    cin>>str;
    for(int i = 0; i<str.size(); i++)
    {
        ins[i] = str[i];
    }
    
    int out = dfs(0, ins);
    cout<<out-1<<endl;
    

    
    return 0;
}

