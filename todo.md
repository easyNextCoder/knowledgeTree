# todo

1. mysql binlog的三种模式

   1. row
   2. statement
   3. mixed

   一共三种模式各有优缺点

2. mvcc具体原理

3. binlog的组提交

   binlog是mysql server层的，所有的mysql都会有

   undolog 和 redolog是innodb层的

   https://juejin.cn/post/6987557227074846733

4. 复制

5. 高可用

   1. 如果没有使用mysql cluster NDB，实现高可用性的成本比较低，不象传统的高可用方案一样需要共享的存储设备和专用的软件才能实现

   2. 使用mysql cluster NDB，但是是全部内存操作（具体的速度跟网络环境有很大关系）的（目前也支持硬盘操作了）

      https://blog.51cto.com/ty1992/1534313
      https://dev.mysql.com/doc/refman/5.7/en/mysql-cluster.html

6. 联合索引失效条件

7. NDB cluster

8. 学习文档
   https://dev.mysql.com/doc/sakila/en/

git最终的总结git commit --amend. HEAD^与HEAD~的用法 git branch -f main HEAD~2

linux中的正则表达式和其他语言中使用的正则表达式

* git stash pop的时候出现冲突的时候，会出现的情况：
  * ”delete by us“的情况，但是实际上不是自己删的，而[是这个stash加上来的](https://groups.google.com/g/git-users/c/zKpi8QoKy00?pli=1)
  * pop出来之后这个stash并不会消失



* Git fetch && git rebase origin/master && git push **能减少两次git checkout的切换，一定要试一试**

9. 将go中的位操作总结下来
10. mysql中show master status能显示数据库当前在写的binlog文件名和当前的position
11. mysql使用auto_increment是效率最高的，但是也会产生坏的影响，例如产生热点id争用现象
12. grant all privileges on *.* to 'root'@'localhost' identified by 'root' with grant option; 权限操作学习





go: github.com/stnc/pongo2gin invalid version: git fetch -f origin refs/heads/*:refs/heads/* refs/tags/*:refs/tags/* in /go/pkg: exit status 128: fatal: could not read Username for 'https://github.com': terminal prompts disabled
