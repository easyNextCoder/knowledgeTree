// 剑指offer-圆圈中最后剩下的数字.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <array>

using namespace std;

class Solution {
public:
    struct Node {
        int val = 0;
        Node* before = NULL;
        Node* next = NULL;
    };
    //测试vector中删除的效率，测试array的使用方法
    //array n(10);

    int lastRemaining(int n, int m) {
        Node* head = new Node[n];
        for (int i = 0; i < n; i++) {
            head[i].val = i;
            if (i - 1 >= 0) {
                head[i - 1].next = &head[i];
                head[i].before = &head[i - 1];
            }
        }
        head[n - 1].next = head;
        head[0].before = &head[n - 1];

        int leftSize = n;
        Node* tmpHead = head;
        int rgap;
        while (leftSize > 1) {
            rgap = m % leftSize;
            if (rgap == 0) {
                tmpHead = tmpHead->before;
            }
            else {
                rgap--;
                while (rgap--) {
                    tmpHead = tmpHead->next;
                }
            }

            Node* first = tmpHead->before;
            Node* second = tmpHead->next;
            first->next = second;
            second->before = first;
            tmpHead = second;
            leftSize--;
        }
        return tmpHead->val;
    }
};


int main()
{
    Solution solution;
    cout<<solution.lastRemaining(70866,116922);
    std::cout << "Hello World!\n";
}
