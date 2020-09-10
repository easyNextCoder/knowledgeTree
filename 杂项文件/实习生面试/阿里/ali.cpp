#include <iostream>
#include <set>
#include <vector>
#include "BinaryTree.h"
using namespace std;

struct BinaryTreeNode
{
    int                    m_nValue;
    BinaryTreeNode* m_pLeft;
    BinaryTreeNode* m_pRight;
};

BinaryTreeNode* CreateBinaryTreeNode(int value)
{
    BinaryTreeNode* pNode = new BinaryTreeNode();
    pNode->m_nValue = value;
    pNode->m_pLeft = nullptr;
    pNode->m_pRight = nullptr;

    return pNode;
}

void ConnectTreeNodes(BinaryTreeNode* pParent, BinaryTreeNode* pLeft, BinaryTreeNode* pRight)
{
    if (pParent != nullptr)
    {
        pParent->m_pLeft = pLeft;
        pParent->m_pRight = pRight;
    }
}

BinaryTreeNode* SearchHead(BinaryTreeNode* root, set<int>& leaves)
{
    if (root == NULL)
    {
        return NULL;
    }
    else {
        if (leaves.find(root->m_nValue) != leaves.end())
           return root;								 

        auto leftVal = SearchHead(root->m_pLeft, leaves);
        if (leftVal != NULL)
            return leftVal;

        auto rightVal = SearchHead(root->m_pRight, leaves);
        if (rightVal != NULL)
            return rightVal;
    }
}

void SearchAndRemove(BinaryTreeNode* head, set<int>& leaves)
{
    if (head == NULL) 
	{
        return;
    }
    else {
        if (leaves.find(head->m_nValue) != leaves.end()) 
		{
            leaves.erase(head->m_nValue);
            SearchAndRemove(head->m_pLeft, leaves);
            SearchAndRemove(head->m_pRight, leaves);
        }
        else {
            return;
        }
    }
}

bool isLeafConnected(BinaryTreeNode* root, set<int>& leaves)
{
    BinaryTreeNode* head = SearchHead(root, leaves);
    if (head == NULL)
	{ 
        return false;
    } else {
        //从已经找到的起始节点开始寻找叶子并删除 
        SearchAndRemove(head, leaves);
        return leaves.empty();
    }
}


template<typename testFuncT>
bool Test(vector< pair<set<int>, bool> >& testSets, testFuncT testFunc, BinaryTreeNode* root)
{
	for(auto item: testSets)
	{
		if(testFunc(root, item.first) != item.second)
		{
			return  false;
		}
	}
	return true;
}

int main()
{
    BinaryTreeNode* pNode1 = CreateBinaryTreeNode(1);
    BinaryTreeNode* pNode2 = CreateBinaryTreeNode(2);
    BinaryTreeNode* pNode3 = CreateBinaryTreeNode(3);
    BinaryTreeNode* pNode4 = CreateBinaryTreeNode(4);
    BinaryTreeNode* pNode5 = CreateBinaryTreeNode(5);
    BinaryTreeNode* pNode6 = CreateBinaryTreeNode(6);
    BinaryTreeNode* pNode7 = CreateBinaryTreeNode(7);
    BinaryTreeNode* pNode8 = CreateBinaryTreeNode(8);
    BinaryTreeNode* pNode9 = CreateBinaryTreeNode(9);
    BinaryTreeNode* pNode10 = CreateBinaryTreeNode(10);
    BinaryTreeNode* pNode11 = CreateBinaryTreeNode(11);
    BinaryTreeNode* pNode12 = CreateBinaryTreeNode(12);
    BinaryTreeNode* pNode13 = CreateBinaryTreeNode(13);
    BinaryTreeNode* pNode14 = CreateBinaryTreeNode(14);
    BinaryTreeNode* pNode15 = CreateBinaryTreeNode(15);

    //build left
    ConnectTreeNodes(pNode1, pNode2, pNode3);
    ConnectTreeNodes(pNode2, pNode4, pNode5);
    ConnectTreeNodes(pNode5, pNode6, pNode7);
    //build right
    ConnectTreeNodes(pNode3, NULL, pNode8);
    ConnectTreeNodes(pNode8, pNode9, pNode10);
    ConnectTreeNodes(pNode9, pNode11, pNode12);
    ConnectTreeNodes(pNode12, pNode13, pNode14);
    ConnectTreeNodes(pNode14, pNode15, NULL);

    
    //pair<input leaves, expected result>
    vector< pair<set<int>, bool> > testSets= {
    	{{}, false},
		{{1},true},
		{{1,2,3},true},
		{{1,2,3,4,5},true},
		{{1,2,5},true},
		{{1,2,6},false},
		{{1,3,8,9},true},
		{{1,3,10},false},
		{{9,12,14},true},
		{{12,13,14,15},true},
		{{11,12,13},false},
		{{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},true},
		{{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}, false}, 
		{{-1}, false},
		{{-1,-2,-100}, false}
	};
	
	typedef bool testFuncT(BinaryTreeNode* , set<int>&);
	cout << "Test begin:"<<endl;
    cout << Test<testFuncT>(testSets, isLeafConnected, pNode1) <<endl;
    //泛型函数以后可以拓展测试不同类型的函数 
    
    return 0;
}
