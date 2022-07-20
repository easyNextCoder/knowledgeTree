/*******************************************************************
Copyright(c) 2016, Harry He
All rights reserved.

Distributed under the BSD license.
(See accompanying file LICENSE.txt at
https://github.com/zhedahht/CodingInterviewChinese2/blob/master/LICENSE.txt)
*******************************************************************/

//==================================================================
// 《剑指Offer——名企面试官精讲典型编程题》代码
// 作者：何海涛
//==================================================================

// 面试题68：树中两个结点的最低公共祖先
// 题目：输入两个树结点，求它们的最低公共祖先。

#include <cstdio>
#include "..\Utilities\Tree.h"
#include <list>
#include <iostream>

using namespace std;

bool GetNodePath(const TreeNode* pRoot, const TreeNode* pNode, list<const TreeNode*>& path)
{
    if(pRoot == pNode)
        return true;
 
    path.push_back(pRoot);
 
    bool found = false;

    vector<TreeNode*>::const_iterator i = pRoot->m_vChildren.begin();
    while(!found && i < pRoot->m_vChildren.end())
    {
        found = GetNodePath(*i, pNode, path);
        ++i;
    }
 
    if(!found)
        path.pop_back();
 
    return found;
}

const TreeNode* GetLastCommonNode
(
    const list<const TreeNode*>& path1, 
    const list<const TreeNode*>& path2
)
{
    list<const TreeNode*>::const_iterator iterator1 = path1.begin();
    list<const TreeNode*>::const_iterator iterator2 = path2.begin();
    
    const TreeNode* pLast = nullptr;
 
    while(iterator1 != path1.end() && iterator2 != path2.end())
    {
        if(*iterator1 == *iterator2)
            pLast = *iterator1;
 
        iterator1++;
        iterator2++;
    }
 
    return pLast;
}

const TreeNode* GetLastCommonParent(const TreeNode* pRoot, const TreeNode* pNode1, const TreeNode* pNode2)
{
    if(pRoot == nullptr || pNode1 == nullptr || pNode2 == nullptr)
        return nullptr;
 
    list<const TreeNode*> path1;
    GetNodePath(pRoot, pNode1, path1);
 
    list<const TreeNode*> path2;
    GetNodePath(pRoot, pNode2, path2);
 
    return GetLastCommonNode(path1, path2);
}
/*
   @author：徐永康
   @date:2020.4.21-10:55
   @for:另外一种方法求任意结构树（不只针对二叉树和二叉搜索树）的公共节点（不借助外部存储）
   @err message:在test3中最近公共祖先的定义不同导致failed
 */
const TreeNode* XGetLastCommonParent(const TreeNode* root, const TreeNode* p, const TreeNode* q)
{
    static int find_num = 0;
    bool common_node = 0;
    if (root == NULL)
        return NULL;

    if (root != p && root != q) 
    {
        //从这里一直一条路径寻找没有分叉，一旦知道那个分叉点就得到了结果
        if (find_num == 0)common_node = 1;
    }else if (root == p || root == q) 
    {
        //上一层的点或许是个分叉点,也可能不是
        find_num++;
        if (find_num == 1)common_node = 1;
        if (find_num == 2)return NULL;
    }

    auto iter = root->m_vChildren.begin();
    for (iter; iter != root->m_vChildren.end(); iter++)
    {
        const TreeNode* rval = XGetLastCommonParent(*iter, p, q);
        if (rval != NULL)
            return rval;
        
        if (find_num == 2 && common_node == 0) { return NULL; }
        else if (find_num == 2 && common_node == 1) {
            find_num = 0;
            return root;
        }
    }

    return NULL;

}
/*
   @author：徐永康
   @date:2020.4.21-10:55
   @for:另外一种方法求任意结构树（不只针对二叉树和二叉搜索树）的公共节点（不借助外部存储）
   @err message:在test3中最近公共祖先的定义不同导致failed
 */
const TreeNode* YGetLastCommonParent(const TreeNode* root, const TreeNode* p, const TreeNode* q)
{
    if (root == NULL)
        return NULL;
    if (root == p || root == q)
        return root;
    
    auto iter = root->m_vChildren.begin();
    vector<const TreeNode*>tmpv;
    for (iter; iter != root->m_vChildren.end(); iter++)
    {
        const TreeNode* rval = XGetLastCommonParent(*iter, p, q);
        if (rval != NULL)
            tmpv.push_back(rval);
    }
    cout << "tmpv.size() is:"<<tmpv.size();
    if (tmpv.size() == 2)
        return root;
    else 
        return tmpv.empty() ? NULL : tmpv[0];//如果最终的节点不是根节点，则通过这个路径返回
    
}

// ====================测试代码====================

void Test(const char* testName, const TreeNode* pRoot, const TreeNode* pNode1, const TreeNode* pNode2, TreeNode* pExpected)
{
    if(testName != nullptr)
        printf("%s begins: ", testName);

    const TreeNode* pResult = YGetLastCommonParent(pRoot, pNode1, pNode2);
    cout << "pResult->value:"<<pResult->m_nValue << endl;
    cout << "pExpected->value:"<<pExpected->m_nValue << endl;
    if((pExpected == nullptr && pResult == nullptr) || 
        (pExpected != nullptr && pResult != nullptr && pResult->m_nValue == pExpected->m_nValue))
        printf("Passed.\n");
    else
        printf("Failed.\n");
}




// 形状普通的树
//              1
//            /   \
//           2     3
//       /       \
//      4         5
//     / \      / |  \
//    6   7    8  9  10
void Test1()
{
    TreeNode* pNode1 = CreateTreeNode(1);
    TreeNode* pNode2 = CreateTreeNode(2);
    TreeNode* pNode3 = CreateTreeNode(3);
    TreeNode* pNode4 = CreateTreeNode(4);
    TreeNode* pNode5 = CreateTreeNode(5);
    TreeNode* pNode6 = CreateTreeNode(6);
    TreeNode* pNode7 = CreateTreeNode(7);
    TreeNode* pNode8 = CreateTreeNode(8);
    TreeNode* pNode9 = CreateTreeNode(9);
    TreeNode* pNode10 = CreateTreeNode(10);

    ConnectTreeNodes(pNode1, pNode2);
    ConnectTreeNodes(pNode1, pNode3);

    ConnectTreeNodes(pNode2, pNode4);
    ConnectTreeNodes(pNode2, pNode5);

    ConnectTreeNodes(pNode4, pNode6);
    ConnectTreeNodes(pNode4, pNode7);

    ConnectTreeNodes(pNode5, pNode8);
    ConnectTreeNodes(pNode5, pNode9);
    ConnectTreeNodes(pNode5, pNode10);

    Test("Test1", pNode1, pNode6, pNode8, pNode2);
}

// 树退化成一个链表
//               1
//              /
//             2
//            /
//           3
//          /
//         4
//        /
//       5
void Test2()
{
    TreeNode* pNode1 = CreateTreeNode(1);
    TreeNode* pNode2 = CreateTreeNode(2);
    TreeNode* pNode3 = CreateTreeNode(3);
    TreeNode* pNode4 = CreateTreeNode(4);
    TreeNode* pNode5 = CreateTreeNode(5);

    ConnectTreeNodes(pNode1, pNode2);
    ConnectTreeNodes(pNode2, pNode3);
    ConnectTreeNodes(pNode3, pNode4);
    ConnectTreeNodes(pNode4, pNode5);

    Test("Test2", pNode1, pNode5, pNode4, pNode3);
}

// 树退化成一个链表，一个结点不在树中
//               1
//              /
//             2
//            /
//           3
//          /
//         4
//        /
//       5
void Test3()
{
    TreeNode* pNode1 = CreateTreeNode(1);
    TreeNode* pNode2 = CreateTreeNode(2);
    TreeNode* pNode3 = CreateTreeNode(3);
    TreeNode* pNode4 = CreateTreeNode(4);
    TreeNode* pNode5 = CreateTreeNode(5);

    ConnectTreeNodes(pNode1, pNode2);
    ConnectTreeNodes(pNode2, pNode3);
    ConnectTreeNodes(pNode3, pNode4);
    ConnectTreeNodes(pNode4, pNode5);

    TreeNode* pNode6 = CreateTreeNode(6);

    Test("Test3", pNode1, pNode5, pNode6, nullptr);
}

// 输入nullptr
void Test4()
{
    Test("Test4", nullptr, nullptr, nullptr, nullptr);
}

int main(int argc, char* argv[])
{
    Test1();
    Test2();
    Test3();
    Test4();

    return 0;
}

