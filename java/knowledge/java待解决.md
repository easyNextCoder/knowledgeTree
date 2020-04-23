在委托事件模型中，源生成事件并把它发送给一个或多个监听器，每个监听器必须向事件源注册。

访问权限控制从最大权限到最小权限依次为：public、 包访问权限、protected和private



讲一下redis的主从复制怎么做的？

请谈谈你对Javaweb开发中的监听器的理解？ 


/*
public class TreeNode {
    int val = 0;
    TreeNode left = null;
    TreeNode right = null;

    public TreeNode(int val) {
        this.val = val;

    }

}
*/
import java.util.*;
public class Solution {
    void travel(TreeNode root, List<Integer> out){
        if(root == null){
            return ;
        }else{
            travel(root.left, out);
            out.add(root);
            travel(root.right, out);
        }
    }
    TreeNode KthNode(TreeNode pRoot, int k)
    {
        
        List<TreeNode> out = new ArrayList<>();
        travel(pRoot, out);
        return out.get(k-1);
    }


}