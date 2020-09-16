/*
 *date:2020.8.3
 *for:验证使用自定义类型的大小堆数据结构，重载函数作为内部成员函数和友员函数的写法不同！
 */
#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>


//https://www.runoob.com/cplusplus/cpp-overloading.html
//函数重载学习
using namespace std;
struct ListNode{
    ListNode() = default;
    ListNode(int a):val(a){};
    int val;
    ListNode* next;
    // friend bool operator < (const ListNode a, const ListNode b)
    // {
    //     if(a && b)
    //     {
    //         return a.val > b.val;
    //     }
    // }
    //当重载作为成员函数的时候，虽然意义上需要两个参数，实际上也要写两个成员函数

    // bool operator < (const ListNode a)
    // {
    //     return this->val > a.val;
    // }
    //当重载作为成员函数的时候，虽然意义上需要两个参数，实际上只写一个，另一个是this
};


struct cmp{
    bool operator()(ListNode* a, ListNode* b)
    {
        if(a != NULL && b != NULL)
        {
            return a->val > b->val;
        }
    }
};

int main()
{
    ListNode *head = new ListNode(-1);
    ListNode *headc = head;
    //priority_queue<ListNode*, vector<ListNode*>, cmp> mq;
    priority_queue<ListNode*, vector<ListNode*>, cmp> mq;
    vector<int> input = {9,5,7,4,8,2,4,1,0};
    vector<int> input1 = {90,15,7,4,58,21,4,1,0};
    vector<int> input2 = {19,105,71,44,55,21,4,10,70};
    vector<vector<int>> intCon = {input, input1, input2};
    vector<ListNode*> lstCon;
    for(auto& vec: intCon)
    {
        head = headc;
        sort(vec.begin(), vec.end());
        for(auto item:vec)
        {
            cout<<item<<" ";
            head->next = new ListNode(item);
            head = head->next;
        }
        cout<<endl;
        head->next = NULL;
        lstCon.push_back(headc->next);
        mq.push(headc->next);
    }
    head = headc;
    ListNode* headOut = new ListNode(-1);
    ListNode* headOutc = headOut;
    while(!mq.empty())
    {
        ListNode* tmp = mq.top();
        mq.pop();
        
        headOut->next = tmp;
        if(tmp->next != NULL)
        {
            mq.push(tmp->next);
        }
        headOut = headOut->next;
    }
    headOut->next = NULL;
    headOut = headOutc->next;
    while(headOut)
    {
        cout<<headOut->val<<" ";
        headOut = headOut->next;
    }
    cout<<endl;

    return 0;
}