/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* mergeSort(ListNode* head, int len)
    {
        if(len<=1)return head;

        int leftLen = len/2;
        int rightLen = len-leftLen;
        ListNode* headl = head;
        ListNode* headr = head;
        leftLen -= 1;
        while(leftLen--)
        {
            headr = headr->next;
        }
        ListNode* tmp = headr->next;
        headr->next = NULL;
        headr = tmp;

        headl = mergeSort(headl, len/2);
        headr = mergeSort(headr, rightLen);

        ListNode* headOut = new ListNode();
        ListNode* headOutc = headOut;
        while(headl || headr)
        {
            if(headl && headr)
            {
                if(headl->val < headr->val)
                {
                    headOut->next = headl;
                    headOut = headOut->next;
                    headl = headl->next;
                }else{
                    headOut->next = headr;
                    headOut = headOut->next;
                    headr = headr->next;
                }
            }else if(headl){
                headOut->next = headl;
                headOut = headOut->next;
                headl = headl->next;
            }else if(headr){
                headOut->next = headr;
                headOut = headOut->next;
                headr = headr->next;
            }
        }
        headOut->next = NULL;
        return headOutc->next;

    }
    ListNode* sortList(ListNode* head) {
        //计数之后，采用归并排序
        int  count = 0;
        ListNode* headc = head;
        while(headc != NULL)
        {
             count++;
             headc = headc->next;
        }
           
        return mergeSort(head, count);
    }
};