# git命令的测试和使用

## reset(soft mix hard)复位到哪个位置节点
- soft比mix模式更进一步，还将add记录了下来（也就是已经stage了之后直接commit就算完成了）

-mix   使用混合模式，使用reset复位到了某个commit，之前的commit都被删除了，但是以往commit的文件都被保存到原来的位置并没有动，也就是可以再次提交

-hard 如果使用hard，所有的commit也会消失，而且所有commit期间的文件也都消失了，所以之前历史的commit增加的文件无法再次提交了。

-->所以一个小总结：soft保存着你add的历史（也就是已经帮你stage过了），-mix还保存着你的文件，-hard连文件都给你删除了。


注意：不管你什么时候reset，没有add和commit的文件会在切换分支和resetcommit的过程中一直保持文件空悬着


## rebase && merge

就是将非当前分支重新档位base，这样就把自己提交的commit放到非当前分支修改的commit之前，并合并成一条线。这个方法可以消除merge对以后版本回退带来的困难。但是注意如果当前分支和非当前分支分别提交的commit中有对相同部分的内容的修改的话，就会无法rebase。如自己的截图

* merge 和rebase的使用
> https://www.zhihu.com/question/32163005/answer/76577586

![tupian](./can_not_rebase_situation.PNG)

## revert commit

其实际的含义就是对某个commit反悔了，可以修改后重新提交


## checkout ...branch

迁出到某个分支

* 注意点：

commit和commit && push之间的区别是什么？

## git命令

git add .

git commit -m ""(windows下是单引号)

git checkout -b mybranch

删除分支： $ git branch -d mybranch
强制删除分支： $ git branch -D mybranch
列出所有分支： $ git branch

更新远程库到本地： $ git fetch origin
推送分支： $ git push origin mybranch
取远程分支合并到本地： $ git merge origin/mybranch
取远程分支并分化一个新分支： $ git checkout -b mybranch origin/mybranch

更新master主线上的东西到该分支上：$git rebase master

切换到master分支：$git checkout master

>https://blog.csdn.net/carolzhang8406/article/details/49757101


学习git的最直观的方法：打游戏
> https://learngitbranching.js.org/?locale=zh_CN

游戏中使用的命令
git branch -f branchname fe30  
    fe30(是head的名称，这个命令强制将branchname这个分支向前或者向后移动到fe30这个HEAD处)

git rebase branchname

    将branchname处的提交拿到自己的提交下面，使提交更加线性

git rebase -i HEAD~3

    改变从当前HEAD处开始包括这个结点的前面三个结点的顺序，可以排列他们的顺序，并且可以删除
    这些操作是git会给提供一个ui来让我们操作

git cherry-pick fe30 fec3 ...

    将非当前分支的提交，pick到当前分支

git commit --amend

    就是反悔当前的提交，用现在的提交代替之前的那个提交（但是实际上生成了一个新的HEAD）