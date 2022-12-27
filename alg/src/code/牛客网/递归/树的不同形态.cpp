#include <iostream>
#include <vector>
#include <deque>

using namespace std;

struct TreeNode{
    TreeNode(int value):val(value){};
    TreeNode * left;
    TreeNode * right;
    int val;
};

const int N = 2048;
int con[N];

int main()
{
    int count = 0;
    while(cin>>con[count]){count++;}
    
    int* level = &con[0];
    int* mid = &con[count/2];
    
    deque<vector<int>> childrens;
    childrens.push_back({0, count/2-1});
    
    deque<TreeNode**> father;
    TreeNode* grandFather = new TreeNode(0);
    father.push_back(&grandFather->right);
      
    int midCnt = 0;
    while(!childrens.empty())
    {
        auto reg = childrens.front();
        childrens.pop_front();
        int head = level[midCnt++];
        int index = -1;
        for(int i = reg[0]; i<=reg[1]; i++)
            if(head == mid[i])
            {
                index = i;
                break;
            }
        
        
        TreeNode* root = new TreeNode(head);
        *father.front() = root;
        father.pop_front();
cout<<index<<endl;
cout<<reg[0]<<" "<<reg[1]<<endl;
        if(index == reg[0])
        {
            //只有右孩子，左孩子为空，将右孩子压入栈
            root->left = NULL;
            if(reg[0]+1 <= reg[1])
            {
                father.push_back(&root->right);    
                childrens.push_back({reg[0]+1, reg[1]});
            }
            continue;
        }

        if(index == reg[1])
        {
            //只有左孩子，右孩子为空，将左孩子压入栈
            root->right = NULL;
            if(reg[0] <= reg[1]-1)
            {
                father.push_back(&root->left);
                childrens.push_back({reg[0], reg[1]-1});
            }
            continue;
            
        }


        //有左孩子又有右孩子
        father.push_back(&root->left);
        father.push_back(&root->right);
        childrens.push_back({reg[0], index-1});
        childrens.push_back({index+1, reg[1]});
        
    }
    cout<<grandFather->right->val<<endl;
    
    return 0;
}