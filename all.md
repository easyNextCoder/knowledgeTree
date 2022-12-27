https://portal.qiniu.com/kodo/bucket/resource-v2?bucketName=xyk-web-data

# Mysql

## 运行原理

![](http://www.xykxykxyk.top/KnowledgeTree%2F%E6%95%B0%E6%8D%AE%E5%BA%93%2Fmysql%E8%BF%90%E8%A1%8C%E5%8E%9F%E7%90%86.JPG)

![](https://learning.oreilly.com/api/v2/epubs/urn:orm:book:9781492080503/files/assets/hpm4_0102.png)

* 查询缓存
  Mysql8以后查询缓存已经废除了，因为比较鸡肋

* 词法分析器
  组成：

  * 词法分析器MysqlLex(mysql自己完成)
  * 语法分析器Bison(linux底层框架)

  工作步骤：

  1. 词法分析
  2. 语法分析
  3. 语义分析
  4. 构造执行树
  5. 生成执行计划
  6. 计划的执行

  生成的结果图：

  ![](http://www.xykxykxyk.top/KnowledgeTree%2F%E6%95%B0%E6%8D%AE%E5%BA%93%2Fmysql%E8%AF%AD%E6%B3%95%E6%A0%91.JPG)

* 优化器

  select * from user where id=19 and name="xu"

  where的后面既有id又有name，其中id是主键，优化器的作用是在执行过程中选择走主键

* 执行器
  执行器去调用数据库引擎的api接口（已经提前定义好了相应的接口）

## 数据库三大范式

* 第一范式：原子性
* 第二范式：唯一性
* 第三范式：无冗余性

### 范式和反范式

* 范式
  * 范式的更新操作通常比反范式的要快
  * 当数据较好的范式化时候，就只有很少或者没有重复数据，所以只需要修改更少的数据
  * 范式化的表通常更小，可以更好的放在内存里，所以执行操作会更快
  * 很少有多余的数据意味着检索列表数据时候更少需要DISTINCT或者GRPOUP BY语句

反范式的schema因为所有的数据都在一张表中，可以很好的避免关联。如果不需要关联表，则对大部分查询最差的情况-即使表没有使用索引-是全表扫描。当数据比内存大的时候，这可能比关联要快的多，因为这样避免了随机IO。

## ACID特性

### A(atomic)原子性

​	原子性的实现依靠的是undo log 

### C(consistency)一致性

### I(Isolation)隔离性

* 实现原理
  MVCC多版本并发控制（读取数据的时候，通过一种类似快照的方式将数据保存下来，这样读锁和写锁就不冲突了，不同的事务session会看到自己特定版本的数据，版本链。）
* 隔离级别
  * 读未提交
    可能会读到其他事务未提交的数据，也叫做脏读
  * 读已提交
    两次读取结果不一致，叫做不可重复读。这个隔离级别在**每次select都会生成一个新的Read View**，也就意味着，事务期间的多次读取同一条数据，前后两次读的数据会出现不一致，因为可能这期间另一个事务修改了该记录并提交了事务。
  * 可重复读
    mysql的默认级别，就是每次读取结果都一样，但是有可能产生幻读。这个隔离级别是**启动事务的时候生成一个Read View，然后整个事务期间都在用这个Read View**，这样保证了在事务期间读到的数据都是事务启动前的记录。
    * 幻读
      * 幻读产生的原因
        RR读不会受到其他事务update、insert的影响，但是自己执行了update就会把其他事务insert的数据更新成自己的版本号，下一次读取就会读到了
      * 解决方法
        幻读可以通过MVCC+间隙锁来进行解决
  * 串行化
    最高的隔离级别，强制事务排序，使之不可能发生冲突，足够安全，但是效率足够低，因此生产环境中是不会使用的。但是什么问题都不会有
* 读的方式
  * 当前读
    select lock in share mode(共享锁), select for update, update, insert, delete(排他锁)这些操作都是一种当前读，就是**读的是记录的最新的版本**，读取时候还要保证其他并发事务不能修改当前记录，会对当前记录加锁。
  * 快照读
    像不加锁的select操作就是快照读，即不加锁的非阻塞读。快照读的前提是隔离级别不是未提交读和串行化级别，因为未提交读总是读取最新的数据行，而不是符合当前事务版本的数据行。而串行化则会对所有读取的行都加锁

### D(Durablity)持久化

​	持久化依靠的是redo log

## 引擎

### InnoDB

#### 特点

* 索引
  只有一个聚簇索引（主键索引-一级索引）和多个非聚簇索引（辅助索引-二级索引）组成。

  索引的数据结构：B+树

* 文件存储形式

  有两个个文件

  * xxx.frm
    表结构信息

  * xxx.ibd
    表数据和索引信息

  数据和索引是存储到一块的。都是存在idb（数据文件和索引文件）文件中frm（表结构）

* 使用要求
  Innodb存储引擎在进行数据插入的时候，数据必须要跟索引放在一起，如果有主键就使用主键，没有主键就使用唯一键，没有唯一键就是用自生成的6字节的row_id（对于用户不可见），因此跟数据绑定在一起的就是聚簇索引，而为了避免数据冗余存储，其他的索引的叶子节点中存储的都是聚簇索引的key值，因此innodb中既有聚簇索引也有非聚簇索引，而myisam中只有非聚簇索引。

* 优势
  支持事物、行级锁、外键

  

#### [内存结构](https://dev.mysql.com/doc/refman/5.7/en/innodb-in-memory-structures.html)

- [缓冲池](https://dev.mysql.com/doc/refman/5.7/en/innodb-buffer-pool.html)
- [更改缓冲区](https://dev.mysql.com/doc/refman/5.7/en/innodb-change-buffer.html)
- [自适应哈希索引](https://dev.mysql.com/doc/refman/5.7/en/innodb-adaptive-hash.html)
- [日志缓冲区](https://dev.mysql.com/doc/refman/5.7/en/innodb-redo-log-buffer.html)

![](https://oss-emcsprod-public.modb.pro/wechatSpider/modb_20210918_ec08ea90-1816-11ec-8281-38f9d3cd240d.png)

* Buffer

  innodb 引擎默认的数据页是16kb，而buffer pool启动的时候是默认的128M，所以是有8192个数据页的。而磁盘的数据管理也是用数据页为单位来管理的，所以每次查找数据的时候，先请求buffer pool，buffer pool中没有的话会到磁盘中找到对应的数据页，然后copy到buffer pool中给客户端返回。

  * [Buffer Pool](https://dev.mysql.com/doc/refman/5.7/en/innodb-buffer-pool.html)

    * Buffer Pool

      * Free list
        空闲页链表

        ##### 1. 怎么知道数据页是否被缓存？
  
        数据库中有一个**`数据页缓存哈希表`**我，用**`表空间号+数据页号`，作为一个key，然后缓存页的地址作为value**表空间号+数据页号 = 缓存页地址
  
      * Flush list
        脏页链表

      * LRU list
  
        数据链表（冷热数据链表-用以解决缓冲池污染问题）
  
        * 冷数据块3/8
        * 热数据块5/8
  
    ![](https://dev.mysql.com/doc/refman/5.7/en/images/innodb-buffer-pool-list.png)
    

    * Change Buffer

      [链接1](https://juejin.cn/post/6844903874172551181)

      [链接2](https://juejin.cn/post/6844903875271475213)

      对于为非唯一索引，辅助索引的修改操作并非实时更新索引的叶子页，而是把若干对同一页面的更新缓存起来做，合并为一次性更新操 作，减少IO，转随机IO为顺序IO,这样可以避免随机IO带来性能损耗，提高数据库的写性能

      具体流程：

      先判断要更新的这一页在不在缓冲池中
  
      a、若在，则直接插入；
  
      b、若不在，则将index page 存入Change Buffer，按照Master Thread的调度规则来合并非唯一索引和索引页中的叶子结点
      原文链接：https://blog.csdn.net/MortShi/article/details/122506516
  
      * 为什么只对辅助索引有作用？
        如果数据库都是唯一索引，那么在每次操作的时候都需要判断索引是否有冲突，势必要将数据加载到缓存中对比，因此也用不到 Change Buffer。
  
    * 如何解决buffer pool污染
      新读取的块被插入到 LRU 列表的中间。所有新读取的页面都插入到默认情况下`3/8`位于 LRU 列表尾部的位置。当页面在缓冲池中第一次被访问时，它们被移动到列表的前面（最近使用的一端）。因此，从未访问过的页面永远不会进入 LRU 列表的前面部分，并且比使用严格的 LRU 方法更快地“老化” 。这种安排将 LRU 列表分为两部分，其中插入点下游的页面被认为是“旧的”并且是 LRU 驱逐的理想受害者。
  
  * Log Buffer

[buffer原理讲解1](https://www.modb.pro/db/111341)

[buffer原理讲解2](https://blog.csdn.net/weixin_35952290/article/details/115906914)

#### [文件存储结构](https://zhuanlan.zhihu.com/p/429567830)

InnoDB的物理文件有很多种，包括：

1. 系统表空间（system tablespace）。文件以 ibdata1、ibdata2 等命名，包括元数据数据字典（表、列、索引等）、double write buffer、插入缓冲索引页（change buffer）、系统事务信息（sys_trx）、默认包含 undo 回滚段（rollback segment）。
2. 用户表空间。innodb_file_per_table=true 时，一个表对应一个独立的文件，文件以 db_name/table_name.ibd 命名。行存储在这类文件。另外还有 5.7 之后引入 General Tablespace，可以将多个表放到同一个文件里面。
3. redo log。文件以 ib_logfile0、ib_logfile1 命名，滚动写入。主要满足ACID特性中的 Durablity 特性，保证数据的可靠性，同时把随机写变为内存写加文件顺序写，提高了MySQL的写吞吐。
4. 另外还可能存在临时表空间文件、undo 独立表空间等。

分为一个ibd数据文件-->Segment（段）-->Extent（区）-->Page（页）-->Row（行）

一般情况下一个段管理256个区，每个区1MB大小，如果设置的page是16K那么就有64个，否则个数随着page的大小变化而变化

![](https://pic1.zhimg.com/80/v2-330ad504926d516ebc57acd8cba3c590_1440w.jpg)





![](https://pic2.zhimg.com/80/v2-b0c81c6ad80d3a28be6d226645d693a1_1440w.jpg)





#### 监控命令

InnoDB实时监控
mysql> show engine innodb status\G





### MyIsam

* 特点

* 文件存储形式

  有三个文件

  * xxx.frm
    表结构信息

  * xxx.MYD
    表数据信息

  * xxx.MYI
    表索引

  * 

* 数据结构

### Memory等

### 区别与联系

​	innodb（叶子节点直接放的是data）和myisam（叶子节点放的是对应的数据行的地址）的存储结构图

​	![树结构](C:\Users\xyk\Desktop\knowledgeTree\数据库\mysql\images\innodb和myisam的结构图.JPG)

## 索引

### 索引的分类

#### 从逻辑的角度

* 单列

  * 普通索引（MUL辅助索引）
    仅加速索引，内部的值可以重复（可以有null）

  * 唯一索引（UNI）
    加速查询+列值唯一（不能有重复的值，可以有null)

  * 主键索引（PK）
    加速查询+列值唯一（不能有重复值，不可以有null）+表中只有一个


* 多列
  * 联合索引
    多列值组成一个索引，专门用于组合搜索，效率大于索引合并
  * 索引合并
    使用多个单列索引组合搜索
  * 索引覆盖
    select的数据只用从索引中就能够取得，不必读取数据行，换句话说查询列要被所建的索引覆盖
  

#### 从物理存储角度

##### 聚簇索引和非聚簇索引

* 如何区分？
  看看数据和索引是否存储在了一块
* 存储位置
  索引都是直接存到磁盘上的
* 都是B+树
  * 聚簇索引
    数据和索引存储到了一块的，找到了索引也就找到了数据
    * 优点
      * 对范围查找效果非常好
      * 适合排序
      * 效率高不需要二次查询
    * 劣势
      * 维护索引代价高，插入数据导致索引移动，也会造成内存碎片
      * 存储稀疏因为使用UUID随机ID生成数据，扫描速度会很慢
  * 非聚簇索引
    叶子节点不存储数据，存储的是数据行地址，根据数据行地址再去磁盘中查找数据
* InnoDB一定有主键，主键一定是聚簇索引

### 索引的数据结构

* B+树索引
  目前99.99%的使用场景

* Hash索引

  Hash算法
  优点：查询效率高
  缺点：不支持范围查询

### 存储引擎

#### innodb

![](C:\Users\xyk\Desktop\knowledgeTree\数据库\mysql\images\索引结构.JPG)

B+Tree(B-Tree变种)

* 非叶子节点不存储data,只存储索引，可以放更多的索引
* 叶子节点包含所有索引字段
* 叶子节点用指针连接，提高区间访问的性能

#### myIsam

![](C:\Users\xyk\Desktop\knowledgeTree\数据库\mysql\images\myIsam的索引的数据结构.JPG)



### [联合索引和最左匹配原则](https://blog.csdn.net/qq_39408664/article/details/118889666)

#### 覆盖索引

​	就是select的数据列只从索引中就能得到，不必读取数据行，也就是只需扫描索引就可以得到查询结果。当一个查询使用了索引覆盖，在查询分析器explain的Extra列可以看到Using index

#### 回表

#### 索引下推-索引条件

英文名称：[using index condition](https://www.bilibili.com/video/BV1ta411C7xq?p=6)

#### [索引截断](https://zhuanlan.zhihu.com/p/471209432)

### 索引问题

* 优秀讲解视频
  [链接](https://www.bilibili.com/video/BV1ta411C7xq?p=6)

* <>查询的时候会走索引嘛？
  <>查询的时候会走索引

* like查询的时候会走索引吗？

  like的时候，如果不是%开头的就会走索引，查询出全部的以xx开头的信息

* **索引失效**的几种情况？

  * 总结
    模糊匹配、使用函数、隐式类型转换、没有遵循最左匹配原则、OR后边跟的不是索引项
  * 当我们使用左或者左右模糊匹配的时候，like %xx和like %xx%都会造成索引失效
  * 当我们在查询条件中对索引列使用函数，就会导致索引失效
  * 当我们在查询条件中对索引列进行表达式计算，也会导致索引失效
  * MySQL 在遇到字符串和数字比较的时候，会自动把字符串转为数字，然后再进行比较。如果字符串是索引列，而条件语句中的输入参数是数字的话，那么索引列会发生隐式类型转换，由于隐式类型转换是通过 CAST 函数实现的，等同于对索引列使用了函数，所以就会导致索引失效。
  * 联合索引要能正确使用需要遵循最左匹配原则，也就是按照最左优先的方式进行索引的匹配，否则就会导致索引失效。
  * 在 WHERE 子句中，如果在 OR 前的条件列是索引列，而在 OR 后的条件列不是索引列，那么索引会失效。

* 主键索引和唯一索引哪个更快？
  主键索引更快，因为主键索引不用回表查询。唯一索引里面存放的就是对应的主键值

* 主键的设置规则
  主键一般是使用整形的自增型变量来充当的，最好不要使用uuid因为uuid既不是整形的也不是自增的

## 锁

[独占锁、共享锁、意向锁和记录锁](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)

SELECT…FOR SHARE是 MySQL 8.0 的一项功能，它取代了SELECT……LOCK IN SHARE MODE以前的版本。

### 按照属性划分

- [Shared and Exclusive Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-shared-exclusive-locks)
- [Intention Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-intention-locks)
- [Record Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-record-locks)
- [Gap Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-gap-locks)
- [Next-Key Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-next-key-locks)
- [Insert Intention Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-insert-intention-locks)
- [AUTO-INC Locks](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-auto-inc-locks)
- [Predicate Locks for Spatial Indexes](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-predicate-locks)

* 共享锁

  共享锁又称读锁，一个事务为数据加上了读锁之后，其他事务只能对该数据加读锁，而不能加写锁

* 排他锁
  排他锁又称写锁，当一个事务为数据加上写锁时，其他请求讲不能再为数据加任何锁，知道该锁释放之后，其他事务才能对数据进行加锁。不允许其他人读数据，也不允许写数据。

### 按照锁的粒度划分

* 行级锁innodb
  粒度小，加锁比表锁麻烦，不容易冲突，相比表锁支持并发要高。
* 表级锁 innodb、myisam
  一个表锁加上了之后，其他事务只能等待。粒度大加锁简单，容易冲突
* 页级锁(innodb引擎）
* 记录锁
* 间隙锁
* 临键锁

### 按照锁的状态分类

* 意向共享锁
* 意向排他锁

## 复制

### 主从同步

* 原理

  * 主从同步过程中，主库会将所有的操作事件记录在binlog中，从库通过开启一个I/O线程保持与主库的通信，并在一定事件间隔内探测binlog日志文件是否发生改变
  * 如果binlog日志发生了变化，主库生成一个binlog dump 线程向从库I/O线程传送binlog
  * 从库上的I/O线程将binlog复制到自己的relay log中
  * 最终由从库中的SQL线程读取relay log中的事件从放到从库上

* [搭建过程](https://blog.csdn.net/poizxc2014/article/details/123522473)
  主要通过CHANGE MASTER TO命令来设置从节点同步于主节点，之间通过TCP/IP进行通信的

* 主从延时原因

  * 随机重放
    mysql主库的binlog的操作是顺序写的，磁盘的顺序读写速度很快。从库sql线程重放过程中是随机读取的
  * 主库高并发
    大量请求打到主库上，意味着要不断对binlog进行写入，此时从库中的sql就会应接不暇
  * 锁等待
    如果某一个时刻从库因为查询产生了锁等待的情况，此时只有当前的操作完成了才会进行下面的操作

* 主从延迟解决方案

  * 半同步复制
  介于异步复制和全同步复制之间，主库在执行万客户端提交的事务后不是立刻返回给客户端，而是等待至少一个从库接到并写到relay log中才返回给客户端。相当于异步复制，半同步复制提高了数据的安全性，同时它也造成了一定程度的延迟，这个延迟最少是一个TCP/IP的往返时间，所以半同步复制最好在低延迟的网络中使用
  
* 并行复制
  
* 降低主库并发（使用redis）
  
  * 读主库
  
    对一些实时性要求比较高的数据，选择读取主库数据
  
* 备库变成其他库主库的命令
  Log_slave_update

## Mysql命令

[所有功能学习网址](https://www.mysqlzh.com/doc/116.html)

### DML-DDL-DCL-TCl

* DML数据库操纵语言

  由DBMS提供，用于让用户或程序员使用，实现对数据库中数据的操作。
  DML分成交互型DML和嵌入型DML两类。
  依据语言的级别，DML又可分成过程性DML和非过程性DML两种。**需要commit**.
  
  * SELECT
  * INSERT
  * UPDATE
  * DELETE
  * MERGE
  * CALL
  * EXPLAIN PLAN
  * LOCK TABLE

* DDL数据库定义语言

  DDL是**SQL**语言的四大功能之一。
  用于定义数据库的三级结构，包括外模式、概念模式、内模式及其相互之间的映像，定义数据的完整性、安全控制等约束
  **DDL不需要commit**.

  * CREATE
  * ALTER
  * DROP
  * TRUNCATE
  * COMMENT
  * RENAME

* **DCL**（**Data Control Language**）**数据库控制语言** 授权，角色控制等
  
  * GRANT 授权
  * REVOKE 取消授权
  
* **TCL**（**Transaction Control Language**）**事务控制语言**
  
  * SAVEPOINT 设置保存点
  * ROLLBACK 回滚
  * SET TRANSACTION

**SQL主要分成四部分**：
（1）数据定义。（SQL DDL）用于定义SQL模式、基本表、视图和索引的创建和撤消操作。
（2）数据操纵。（SQL DML）数据操纵分成数据查询和数据更新两类。数据更新又分成插入、删除、和修改三种操作。
（3）数据控制。包括对基本表和视图的授权，完整性规则的描述，事务控制等内容。
（4）嵌入式SQL的使用规定。涉及到SQL语句嵌入在宿主语言程序中使用的规则。

### 查索引

show index from tb_name;

### 删表

* 删表可用drop、delete、truncate
  * delete
    * DML语言，操作的事务记录在日志中保存可以回滚，需要手动提交事务之后才会生效，可以带上Where条件
  * truncate
    * DDL语言，操作会立即生效，改操作不会走DML，原数据不会放到rollback segment中不可会滚
* 删表为什么要用truncate
  delete语句是dml语句会记录到bin_log中会占用磁盘的ipos，而且bin-log不断增长会造成数据删除后磁盘占用空间增大
* [三种删除方法的更深层次影响](https://www.modb.pro/db/411522)

### 函数

#### rank() over

* rank() over
  1，1，3，4，5，5，7
* dense_rank() over
  有并列排名1，1，2，2，3，4
* Row_number() over
  无论是否重复都是1，2，3，4
  https://leetcode.cn/problems/rank-scores/solution/dense_rank-overpai-ming-de-shi-yong-by-q-mq4s/

#### date_sub()

* 意义
  从当前日期获得前n个时刻的日期

* 用法
  date_sub(b.recordDate, interval 1 day)

  获得b.recordDate的前一天的日期

* 使用示例

  [197. 上升的温度](https://leetcode.cn/problems/rising-temperature/)

#### count

##### count(condition)

Count()内加条件的用法：

* count(IF(id>1, true, NULL))
* count(id>1 or NULL)   or NULL这个条件不可缺少否则无法统计出来

##### count(1)

用来计算第一列所有元素的和（也可以是group by之后的相应的元素的和）



#### group_concat()

​		[1484. 按日期分组销售产品](https://leetcode.cn/problems/group-sold-products-by-the-date/)

​		[group_concat讲解](https://dev.mysql.com/doc/refman/8.0/en/aggregate-functions.html#function_group-concat)

#### ifnull()

IFNULL(expr1,expr2)的用法：

假如expr1 不为 NULL，则 IFNULL() 的返回值为 expr1; 否则其返回值为 expr2。IFNULL()的返回值是数字或是字符串，具体情况取决于其所使用的语境。

#### isnull()

isnull(expr) 的用法：

 如expr 为null，那么isnull() 的返回值为 1，否则返回值为 0。 mysql> select isnull(1+1); -> 0 mysql> select isnull(1/0); -> 1 使用= 的null 值对比通常是错误的。

```mysql
mysql> SELECT IFNULL(1,0);
-> 1
mysql> SELECT IFNULL(NULL,10);
-> 10
mysql> SELECT IFNULL(1/0,10);
-> 10
mysql> SELECT
IFNULL(1/0,'yes');
->   'yes
```

### like

#### [1527. 患某种疾病的患者](https://leetcode.cn/problems/patients-with-a-condition/)

### update ... set ... case ... when...end

(case when end) as ...用法

```mysql
sum(CASE operation
WHEN 'Buy' THEN -price
WHEN 'Sell' THEN price
END) as total
```



[1393. 股票的资本损益](https://leetcode.cn/problems/capital-gainloss/)

[1158. 市场分析 I](https://leetcode.cn/problems/market-analysis-i/)

### union

#### [1873. 计算特殊奖金](https://leetcode.cn/problems/calculate-special-bonus/)

### select

### insert

### update

### delete

### [alter](https://dev.mysql.com/doc/refman/8.0/en/alter-table.html)

​	MYSQL中的alter table操作的性能对于大表来说是个大问题。MySQL执行大部分修改表结构操作的方法是用新的结构创建一个空表，从旧表中查出所有数据插入新表，然后删除旧表。这种操作可能话费很长时间，如果内存不足而表又很大，而且还有很多索引的情况下尤其如此。许多人都有这样的经验，alter table操作需要花费数个小时甚至数天才能完成。

* 更改字符集

  ```mysql
  alter table Student change Sname Sname char(10) character set utf-8;
  ```

### create

* create table 创建数据表

  ```mysql
  create table Course (CID varchar(10), Cname nvarchar(10) , TId varchar(10)) engine=InnoDB default charset=utf8;
  ```
  
  创建一个默认字符集为utf8的索引
  
* create index 创建索引

  create index idx_t1_bcd on t1(b, c, d);

  创建一个联合索引

### group by

* group by xxx 
   作用：对xxx进行聚合
  
  ![group by](C:\Users\xyk\Desktop\knowledgeTree\数据库\mysql\images\group by 函数的运行原理图.JPG)
  
* 语法注意事项
  select的时候只能选择group by的xxx字段或者count、sum等计算函数
  
  ```
  select a.type, count(a.xid) from user_message a inner join user_message b on a.xid=b.xid group by a.type;
  //正确
  
  select * from user_message a inner join user_message b on a.xid=b.xid group by a.type;
  //错误
  ```
  
  

### variables

show global variable like 'innodb_page_size'  16384目前一页是16KB

show variables like '%log_bin%'  查看是否开启了bin log日志

### 权限设置

* 创建用户create
  create user 'ua'@'%' identified by 'pa';

  表示创建的用户名为ua密码是pa，其中%表示用户可以通过任何ip地址以该身份登录到这个数据库

* 删除用户delete
  delete from mysql.user where user='ua';flush privileges;

* 赋予权限grant

  赋予用户ua最高的权限
  grant all privileges on \*.\* to 'ua'@'%' with grant option;

* 收回权限revoke

  ```mysql
  revoke all privileges on *.* from 'ua'@'%';
  ```

* 权限控制的颗粒度

  * 对某个库的权限单独操作

  ```mysql
  grant all privileges on db1.* to 'ua'@'%' with grant option;
  ```

  * 表权限和列权限

  ```mysql
  create table db1.t1(id int, a int);
  grant all privileges on db1.t1 to 'ua'@'%' with grant option;
  GRANT SELECT(id), INSERT (id,a) ON mydb.mytbl TO 'ua'@'%' with grant option;
  ```

  

### 临时表视图

### 外键



## 日志

### [Undo Log](https://dev.mysql.com/doc/refman/5.7/en/innodb-undo-logs.html)

回滚日志，由**引擎层**来实现

* 功能

  保证事务的原子性(回滚)和MVCC

* 特点
  
  1. 每次undo在写入磁盘之前，会先将该动作记录到redo上。是**逻辑日志**，可以认为当delete一条日志的时候，undo log中会对应insert一条记录。
  
  2. Undo的磁盘结构并不是顺序的，而是像数据一样按Page管理
  
     Undo写入时，也像数据一样产生对应的Redo Log
  
     Undo的Page也像数据一样缓存在Buffer Pool中，跟数据Page一起做LRU换入换出，以及刷脏。
     Undo Page的刷脏也像数据一样要等到对应的Redo Log 落盘之后
  
     [链接](https://www.zhihu.com/question/357887214/answer/2204930465)
  
  3. undolog写到redolog中，mysql在进行recover的时候，所用到的undo日志是从redo log里恢复的。msql为了降低复杂度，是将鞋undolog也看做普通的数据写入

![](https://img-blog.csdnimg.cn/20210613084841967.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1dlaXhpYW9odWFp,size_16,color_FFFFFF,t_70)



回滚日志中记录着row_id、事务id、回滚指针。如果事务的第一条语句是插入语句，那么他的回滚指针就是null，update之后生成的undo log的回滚指向上一条语句进而形成一条undo log回滚链。



### Redo Log 

重做日志（物理日志），由**引擎层**实现。是为了实现事务的持久性而出现的产物。防止在发生故障的时间点，尚有脏页未写入表的IBD文件中，在重启Mysql服务的时候，根据Redo Log进行重做，从而达到事务的未入磁盘数据进行持久化这一特性。

![](https://img-blog.csdnimg.cn/c999e31d8277467087615ba2260da330.png)



### Bin Log 

归档日志（逻辑日志），由**服务层**实现，所以只要是mysql无论什么引擎都会有

#### 特点

* 二进制日志
* Binlog在MySQL的Server层实现（引擎共用）
* Binlog为逻辑日志，记录的是一条语句的原始逻辑
* Binlog不限大小，追加写入，不会覆盖以前的日志

#### 功能

主从复制、crash-safe

#### 高级

* XA事务（两阶段提交）实现系统的crash-safe
* 支持组提交降低磁盘的iops压力

#### binlog参数

* 查询是否开启binlog记录功能
  show variables like 'log_bin'

* sync_binlog
  设置为1，表示每次事务binlog都将持久化到磁盘

* binlog_format
  binlog一共三种格式Row（记录每行修改信息）、Statement（语句）、Mixed（混合模式默认的）

  查看binlog的格式：Show variables like 'binlog_format'

  * row

    日志中会记录成每一行数据被修改成的形式，然后在slave端再对相同的数据进行修改，只记录要修改的数据，只有value，不会有sql多表关联的情况。

    优点：在row模式下，bin-log中可以不记录执行的[sql语句](https://so.csdn.net/so/search?q=sql语句&spm=1001.2101.3001.7020)的上下文相关信息，仅仅需要记录哪一条记录被修改了，修改成什么信样了，所以row的日志内容会非常清楚的记录下每一行数据修改的细节，非常容易理解。而且不会出现在某些特定情况下的存储过程和function，以及trigger的调用和处罚无法被正确问题。

* max_binlog_size
  查看最大binlogsize的命令：Show variable like 'max_binlog_size'

  显示binlog文件的最大大小，默认是1GB，当文件大于1GB的时候会生成新的文件， 使用flush log也可新生成一个binlog文件

* log_bin_basename

  查看日志存储地址

  Show variables like 'log_bin_basename'

#### binlog命令

##### 操作binlog

* flush logs
  会多一个最新的bin-log日志
* show master status
  查看最新的binlog日志的相关信息
* reset master
  清空所有的bin-log日志
* reset slave
  删除slave的中继日志
* /usr/local/mysql/bin/mysqlbinlog  -no-defaults /usr/local/mysql/data/mysql-bin.000001
  **查看bin-log内容**
* 查看此时的binlog日志记录
  show binlog events in 'bingo.000002'

##### 数据恢复

* [恢复指定位置段数据](https://dev.mysql.com/doc/refman/5.7/en/point-in-time-recovery-binlog.html)  
  mysqlbinlog --start-position=38543 --stop-position=70770 /usr/local/mysql/data/binlog.000002 |mysql -u root -pmM13137276827_

  建表的position=38543，删除整张表之前的终止position=70770

  将某个表会复制删除其中内容之前

* 恢复指定日期段数据  
  -stop-date="2019-03-02 12:00:00"  --start-date = "2018-03-02 11：55：00"

  * 例子：/usr/local/mysql/bin/mysqlbinlog  --no-defaults /usr/local/mysql/data/mysql-bin.000001 | mysql -u root -p tuling  (指定恢复全部tuling数据库)

### Relay Log

中继日志，在主从复制的过程中复制的从主机拷贝看过来的主机上的binlog，由**服务器层**实现

### Slow Query Log

慢查询日志，记录超过指定查询事件的SQL，由**服务器**层实现

* 什么是慢查询日志
  当查询超过一定时间没有返回结果的时候，才会记录进慢查询日志

  慢查询日志可以帮助DBA找出执行效率缓慢的SQL语句

  慢查询日志默认是不开启的，也没有必要一直开启

  当需要进行采样分析的时候手工开启

* 使用命令开启方式

  slow_query_log=on|off  是否开启慢查询日志

  slow_query_log_file=filename 指定慢查询日志保存路径以及文件名。不设置则使用默认值。默认存放位置为数据库文件所在的目录下，名称为hostname-slow.log

  long_query_time=2指定多少秒未返回结果的查询语句属于慢查询

  long-queries-not-using-indexs记录所有没有使用到索引的查询语句

  min_examined_row_limit=1000 记录那些由于查找了多于1000次而引发的慢查询。
  log-slow-admin-statements 记录那些慢的OPTIMIZE TABLE，ANALYZE TABLE，ALTER TABLE语句。
  log-slow-slave-statements 记录由slave所产生的慢查询。

  * 真实使用示例

    ```mysql
    SET @@global.slow_query_log=1 或者 SET global slow_query_log=1
    SET @@global.long_query_time=3
    其他参数可以通过以下命令查阅：
    SHOW VARIABLES LIKE '%slow%';
    ```

### Error Log

错误日志，由**服务器**层实现

### 其他

#### Bin和Redo log之间的区别

逻辑日志记录的是整个的运算的逻辑过程，物理日志记录的是最后的结果

## Crash-Safe

两阶段提交利用了redo-log和bin-log，正是两阶段提交保证了mysql的Crash-Safe的能力



## Mysql优化

### 参数优化

#### binlog日志

* [**log_bin**](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_log_bin) = off|on

  Whether the binary log is enabled. If the [`--log-bin`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#option_mysqld_log-bin) option is used, then the value of this variable is `ON`; otherwise it is `OFF`. This variable reports only on the status of binary logging (enabled or disabled); it does not actually report the value to which [`--log-bin`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#option_mysqld_log-bin) is set.

* [**log_bin_basename**](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_log_bin_basename) = basename

  Holds the base name and path for the binary log files, which can be set with the [`--log-bin`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#option_mysqld_log-bin) server option. The maximum variable length is 256. In MySQL 5.7, the default base name is the name of the host machine with the suffix `-bin`. The default location is the data directory.

* [**sync_binlog**](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_sync_binlog)

  Controls how often the MySQL server synchronizes the binary log to disk.

  - [`sync_binlog=0`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_sync_binlog): Disables synchronization of the binary log to disk by the MySQL server. Instead, the MySQL server relies on the operating system to flush the binary log to disk from time to time as it does for any other file. This setting provides the best performance, but in the event of a power failure or operating system crash, it is possible that the server has committed transactions that have not been synchronized to the binary log.
  - [`sync_binlog=1`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_sync_binlog): Enables synchronization of the binary log to disk before transactions are committed. This is the safest setting but can have a negative impact on performance due to the increased number of disk writes. In the event of a power failure or operating system crash, transactions that are missing from the binary log are only in a prepared state. This permits the automatic recovery routine to roll back the transactions, which guarantees that no transaction is lost from the binary log.
  - [`sync_binlog=N`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_sync_binlog), where *`N`* is a value other than 0 or 1: The binary log is synchronized to disk after `N` binary log commit groups have been collected. In the event of a power failure or operating system crash, it is possible that the server has committed transactions that have not been flushed to the binary log. This setting can have a negative impact on performance due to the increased number of disk writes. A higher value improves performance, but with an increased risk of data loss.

  For the greatest possible durability and consistency in a replication setup that uses `InnoDB` with transactions, use these settings:

  - [`sync_binlog=1`](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_sync_binlog).
  - [`innodb_flush_log_at_trx_commit=1`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_flush_log_at_trx_commit).

* **binlog_format** = STATEMENT|ROW|MIXED

  说明：**日志格式**

  1. STATEMENT模式（SBR）

     每一条会修改数据的sql语句会记录到binlog中。优点是并不需要记录每一条sql语句和每一行的数据变化，减少了binlog日志量，节约IO，提高性能。缺点是在某些情况下会导致master-slave中的数据不一致(如sleep()函数， last_insert_id()，以及user-defined functions(udf)等会出现问题)。


    2. ROW模式（RBR）
    
       不记录每条sql语句的上下文信息，仅需记录哪条数据被修改了，修改成什么样了。而且不会出现某些特定情况下的存储过程、或function、或trigger的调用和触发无法被正确复制的问题。缺点是会产生大量的日志，尤其是alter table的时候会让日志暴涨。


    3. MIXED模式（MBR）
    
       以上两种模式的混合使用，一般的复制使用STATEMENT模式保存binlog，对于STATEMENT模式无法复制的操作使用ROW模式保存binlog，MySQL会根据执行的SQL语句选择日志保存方式。

* **Max_binlog_size** = GB/MB

  说明：Max_binlog_size: 1073741824=1G ，**binlog的最大值**，一般设置为512M或1G,一般不能超过1G。此参数不能非常严格控制binlog的大小，特别是在遇到大事务时，而binlog日志又到达了尾部，为了保证事务完整性，不切换日志，把所有sql都写到当前日志。

* **expire_logs_days** = N

  说明:设置**binlog老化日期**；有大致三种情况引发日志切换：binlog大小超过max_binlog_size；手动执行flush logs；重新启动时( MySQL将会new一个新文件用于记录binlog)

* **binlog_cache_size** = MB

  说明：默认大小是37268即32K.根据事务需要调整大小。该参数表示在事务中容纳二进制日志sql语句的缓存大小。二进制日志缓存，是服务器支持事务存储引擎并且服务器启用了二进制日志(-log-bin选项)的前提下为每个客户端分配的内存，是每个client都可以分配设置大小的binlog cache空间。

  

#### slowquerylog日志

* slow_query_log = 0|1
  说明:**开关慢查询日志**。
* slow_query_log_file=为存放路径；
* long_query_time =记录超过的时间，默认为10s。



#### innodb-redo日志

##### 刷新规则

* **innodb_flush_log_at_trx_commit** = 0｜1｜2
  通常都取1，默认配置也是1
  
  * The default setting of 1 is required for full ACID compliance. Logs are written and flushed to disk at each transaction commit.（事务提交就写入磁盘并刷新，保证了完整的ACID属性）
  * With a setting of 0, logs are written and flushed to disk once per second. Transactions for which logs have not been flushed can be lost in a crash.（每隔一秒钟写入磁盘并刷新，msyql挂了可能会丢失1S数据）
  * With a setting of 2, logs are written after each transaction commit and flushed to disk once per second. Transactions for which logs have not been flushed can be lost in a crash.（每次都写入磁盘，但是每隔一秒钟刷新一次，操作系统挂了可能会丢失最多1秒钟的事务）
  
  ![](https://img-blog.csdnimg.cn/6a37689c9e16446d831679228f1ddd4c.png)
  
* [**innodb_flush_log_at_timeout**](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_flush_log_at_timeout)

  默认1 最大27000

  每秒写入和刷新日志*`N`* 。 [`innodb_flush_log_at_timeout`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_flush_log_at_timeout) 允许增加刷新之间的超时时间，以减少刷新并避免影响二进制日志组提交的性能。默认设置为 [`innodb_flush_log_at_timeout`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_flush_log_at_timeout) 每秒一次。



#### session/global参数

* join_buffer_size = MB

  One join buffer is allocated for each full join between two tables. For a complex join between several tables for which indexes are not used, multiple join buffers might be necessary.官方建议不要在全局增大这个空间，只在巨量的表，多次join的时候只在session级别设置这个变量

* Sort_Buffer_Size = MB
  主要是用来加速order by 和group by的速度的。最好也是在session中设定。

  说明:Sort_Buffer_Size 是一个**connection级参数**，每个connection第一次需要使用这个buffer的时候，一次性分配设置的内存。Sort_Buffer_Size 并不是越大越好，由于是connection级的参数，过大的设置+高并发可能会耗尽系统内存资源。官网文档说“On Linux, there are thresholds of 256KB and 2MB where larger values may significantly slow down memory allocation”

* innodb_file_per_table = 0|1

  说明：参数值为1，表示对每张表使用单独的 innoDB 文件

* [innodb_flush_method](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_buffer_pool_load_at_startup)

  * 说明:设置InnoDB同步IO的方式：

    * **O_SYNC** 使每次write等待物理I/O操作完成，包括由write操作引起的文件属性更新所需的I/O。
      
    * **O_DSYNC** 使每次write等待物理I/O操作完成，但是如果该写操作并不影响读取刚写入的数据，则不需等待文件属性被更新。
  
    * ***O_DIRECT** 每次*读/写操作 都会跳过OS Cache，直接在device上读/写。当应用有自己的缓存机制，那么O_DIRECT更可取，例如MySQL的InnoDB使用Buffer Pool机制。
  
      ```text
      fsync: InnoDB uses the fsync() system call to flush both the data and log files. fsync is the default setting.
      
      O_DSYNC: InnoDB uses O_SYNC to open and flush the log files, and fsync() to flush the data files. InnoDB does not use O_DSYNC directly because there have been problems with it on many varieties of Unix.
      
      littlesync: This option is used for internal performance testing and is currently unsupported. Use at your own risk.
      
      nosync: This option is used for internal performance testing and is currently unsupported. Use at your own risk.
      
      O_DIRECT: InnoDB uses O_DIRECT (or directio() on Solaris) to open the data files, and uses fsync() to flush both the data and log files. This option is available on some GNU/Linux versions, FreeBSD, and Solaris.
      
      O_DIRECT_NO_FSYNC: InnoDB uses O_DIRECT during flushing I/O, but skips the fsync() system call after each write operation.
      
      			Prior to MySQL 5.7.25, this setting is not suitable for file systems such as XFS and EXT4, which require an fsync() system call to synchronize file system metadata changes. If you are not sure whether your file system requires an fsync() system call to synchronize file system metadata changes, use O_DIRECT instead.
      
      			As of MySQL 5.7.25, fsync() is called after creating a new file, after increasing file size, and after closing a file, to ensure that file system metadata changes are synchronized. The fsync() system call is still skipped after each write operation.
      
      			Data loss is possible if redo log files and data files reside on different storage devices, and an unexpected exit occurs before data file writes are flushed from a device cache that is not battery-backed. If you use or intend to use different storage devices for redo log files and data files, and your data files reside on a device with a cache that is not battery-backed, use O_DIRECT instead.
      ```
  
      [链接](https://zhuanlan.zhihu.com/p/453978775)
  
  
  ![](https://img-blog.csdnimg.cn/img_convert/f3e35edf290ea44770157fd2e738d916.png)
  
  * linux中的fsync
  
    - **sync** 只是将所有修改过的块缓冲区加入写队列，然后就返回，它并不等待实际写磁盘操作结束。所以调用了sync函数，并不意味着已安全的送到磁盘文件上。通常称为update的系统守护进程会周期性地（一般每隔30秒）调用sync函数。这就保证了定期冲洗内核的块缓冲区。
  
    - **fsync** 函数只针对单个文件，只对由文件描述符fd指定的单一文件起作用，并且等待写磁盘操作结束，然后同步返回。fsync不仅会同步更新文件数据，还会同步更新文件的属性（比如atime,mtime等）。fsync可用于数据库这样的应用程序，这种应用程序需要确保将修改过的块立即写到磁盘上。
      fdatasync的功能与fsync类似，但是仅仅在必要的情况下才会同步metadata，因此可以减少一次IO写操作(因为文件的数据和metadata通常存在硬盘的不同地方)
  
      > “fdatasync does not flush modified metadata unless that metadata is needed in order to allow a subsequent data retrieval to be corretly handled.”
  
      举例来说，文件的尺寸（st_size）如果变化，是需要立即同步的，否则OS一旦崩溃，即使文件的数据部分已同步，由于metadata没有同步，依然读不到修改的内容。而最后访问时间(atime)/修改时间(mtime)是不需要每次都同步的，只要应用程序对这两个时间戳没有苛刻的要求，基本无伤大雅。
  
    - **fdatasync** 当初设计是考虑到有特殊的时候一些基本的元数据比如atime，mtime这些不会对以后读取造成不一致性，因此少了这些元数据的同步可能会在IO性能上有提升。该函数类似于fsync，但它只影响文件的数据部分，如果该写操作并不影响读取刚写入的数据，则不需等待文件属性被更新。



#### mysql-server参数

* character-set-server = utf8|utf8mb4
  说明:设定字符集，utf8存3个字节，utf8mb4存4个字节。

* max_connections = xxxx
  默认值：150

  最小值：1 

  最大值：100000
  最大连接数，当数据库面对高并发时，这个值需要调节为一个合理的值，才满足业务的并发要求，避免数据库拒绝连接。

  The maximum permitted number of simultaneous client connections. The maximum effective value is the lesser of the effective value of [`open_files_limit`](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_open_files_limit)` - 810`, and the value actually set for `max_connections`.

* max_user_connections=xxxx
  设置单个用户的连接数。用来对用户来进行限制 



#### Innodb引擎

* transaction_isolation = READ-UNCOMMITTED | READ-COMMITTED |REPEATABLE-READ | SERIALIZABLE(全局变量)
  说明:设定事务隔离级别

  1)未提交读(Read Uncommitted)：允许脏读，也就是可能读取到其他会话中未提交事务修改的数据

  2)提交读(Read Committed)：只能读取到已经提交的数据。Oracle等多数数据库默认都是该级别 (不重复读)

  3)可重复读(Repeated Read)：可重复读。在同一个事务内的查询都是事务开始时刻一致的，InnoDB默认级别。在SQL标准中，该隔离级别消除了不可重复读，但是还存在幻象读

  4)串行读(Serializable)：完全串行化的读，每次读都需要获得表级共享锁，读写相互都会阻塞

* 

* innodb_read_io_threads = xxxx

  | Default Value | `4`  |
  | :------------ | ---- |
  | Minimum Value | `1`  |
  | Maximum Value | `64` |

  The number of I/O threads for read operations in `InnoDB`. Its counterpart for write threads is [`innodb_write_io_threads`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_write_io_threads).

* innodb_write_io_threads = xxx
  数据库写操作时的线程数，用于并发。

 

* innodb_file_per_table= 1

  | Default Value | `ON` |
  | :------------ | ---- |

  When [`innodb_file_per_table`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_file_per_table) is enabled, tables are created in file-per-table tablespaces by default. When disabled, tables are created in the system tablespace by default. 

* innodb_stats_on_metadata={ OFF|on}
  是否动态收集统计信息，开启时会影响数据库的性能(一般关闭，找个时间手动刷新，或定时刷新）如果为关闭时，需要配置数据库调度任务，定时刷新数据库的统计信息。

 

* innodb_spin_wait_delay=xxxxx
  控制CPU的轮询时间间隔，默认是6,配置过低时，任务调度比较频繁，会消耗CPU资源。

  | 默认值              | `6`       |
  | :------------------ | --------- |
  | 最小值              | `0`       |
  | 最大值（64 位平台） | `2**64-1` |
  | 最大值（32 位平台） | `2**32-1` |

  [自旋](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_spin)锁 轮询之间的最大延迟 。该机制的底层实现因硬件和操作系统的组合而异，因此延迟不对应于固定的时间间隔。

 

* innodb_lock_wait_timeout=xxxx
  控制锁的超时时间，默认为50，这个值要注意，如果有特殊业务确实要耗时较长时，不能配置太短。

  | 默认值 | `50`         |
  | :----- | ------------ |
  | 最小值 | `1`          |
  | 最大值 | `1073741824` |
  | 单元   | 秒           |

  `InnoDB` [事务](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_transaction)在放弃之前等待[行锁](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_row_lock) 的时间长度（以秒为单位）。默认值为 50 秒。尝试访问被另一个 `InnoDB`事务锁定的行的事务在发出以下错误之前最多等待这么多秒以对该行进行写访问：

  ```terminal
  ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
  ```

  当发生锁等待超时时，当前语句被 [回滚](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_rollback)（而不是整个事务）。

#### InnoDB缓存池

* innodb_buffer_pool_size
  是用于设置InnoDB缓存池（InnoDBBufferPool）的大小，会缓冲索引页、数据页、undo页、插入缓冲、自适应哈希索引、innodb存储的锁信息、数字字典信息等。InnoDB缓存池的大小对InnoDB的整体性能影响较大，默认值是128M。

   通过查询show status like 'Innodb_buffer_pool_%'，保证Innodb Buffer Pool的Read命中率越高越好：(Innodb_buffer_pool_read_requests – Innodb_buffer_pool_reads) /Innodb_buffer_pool_read_requests * 100%

* innodb_buffer_pool_instance

  允许多个缓冲池实例，每页根据哈希平均分配到不同缓冲池实例中，减少数据库内部资源竞争，可以提升InnoDB的并发性能。默认值是1，表示InnoDB缓存池被划分为一个区域。一般配置数值<=服务器CPU的个数。

* innodb_additional_mem_pool_size
  指定InnoDB用于来存储数据字典和其他内部数据的缓存大小，默认值是2M.InnoDB的表个数越多，就应该适当的增加该参数的大小，当过小的时候，MySQL会记录Warning信息到数据库的错误日志中，这时就需要该调整这个参数大小。对于大数据设置16M足够用。

* innodb_log_buffer_size =xxxxx
  日志缓冲区大小,一般不用设置太大，能存下1秒钟操作的数据日志就行了，mysql默认1秒写一轮询写一次日志到磁盘。



#### InnoDB缓存池内部结构

* [`innodb_old_blocks_pct`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_old_blocks_pct)

  Specifies the approximate percentage of the `InnoDB` [buffer pool](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_buffer_pool) used for the old block [sublist](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_sublist). The range of values is 5 to 95. The default value is 37 (that is, 3/8 of the pool). Often used in combination with [`innodb_old_blocks_time`](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_old_blocks_time).

#### InnoDB缓冲池预热

* innodb_buffer_pool_dump_at_shutdown
  默认是关的，如果开启参数，停止MySQL服务时，InnoDB缓存中的热数据将会保存到硬盘中。

* innodb_buffer_pool_dump_now
  默认关闭，如果开启该参数，停止MySQL服务时，以手动方式将InnoDB缓存池中的热数据保存到本地硬盘。

* innodb_buffer_pool_load_now
  默认关闭，如果开启该参数，启动MySQL服务时，以手动方式将本地硬盘的数据加载到InnoDB缓存池中。

* innodb_buffer_pool_filename
  如果开启InnoDB预热功能，停止MySQL服务时，MySQL将InnoDB缓存池中的热数据保存在磁盘里ib_buffer_pool文件中，位于数据库根目录下，默认文件名是这个参数的值。

> 只有在正常关闭MySQL服务，或者pkill mysql时，会把热数据dump到内存。**机器宕机或者pkill -9 mysql，是不会dump**。

作者：easfire
链接：https://zhuanlan.zhihu.com/p/453978775

### 阿里巴巴规范

[原文连接](https://www.jianshu.com/p/4dc9c82f13d5)

### sql语句优化

* [limit 10000，10 分页优化](https://zhuanlan.zhihu.com/p/419597601)
  * 当id是有序的时候，用子查询id来优化
  * [基于主键索引的子查询方式优化](https://blog.csdn.net/it_lihongmin/article/details/115435262)
    select * from t where id > (select id from t limit 100000,1) limit 20;
* 出现filesort
  如果mysql在排序的时候没有使用到索引那么就会输出using filesort，即使使用文件排序。文件排序是通过相应的排序算法，将取得的数据在内存中进行排序；mysql需要将数据在内存中进行排序。
  * 如何优化？
    * 尽量使用单路排序
    * 加大max_length_for_file_sort_data(让mysql根据空间大小去选择单路排序)
    * 去掉一些不必要的返回字段或者使列长度尽量小
    * 增大buffer_size长度
      [原文链接](https://blog.csdn.net/albertsh/article/details/90296520)

### 分库分表

* 分库
  垂直分库：按照业务模块进行切分，将不同模块的表切分到不同的数据库中。

  如电商系统有电商数据库，按照业务模块可以分为用户库、商品库、订单库，这些都可以当做独立数据库，不需要放到一起。好处是既能独立变更，又能隔绝相互影响。

* 分表
  垂直分表：也就是“大表拆小表”，基于列字段进行的。一般是因为表设计不合理，需要进行拆分。

  如一张表存放学生、老师、课程、成绩信息，最好拆分为学生表、课程表、成绩表。

  水平分表：针对数据量巨大的单张表（比如订单表），按照某种规则（RANGE,HASH取模等），切分到多张表里面去。但是这些表还是在同一个库中，所以库级别的数据库操作还是有IO瓶颈。不建议采用。

* 水平分库分表
  将单张表的数据切分到多个服务器上去，每个服务器具有相应的库与表，只是表中数据集合不同。水平分库分表能够有效的缓解单机和单库的性能瓶颈和压力，突破IO、连接数、硬件资源等的瓶颈。

  ![](C:\Users\xyk\Desktop\knowledgeTree\数据库\mysql\images\分库分表.jpg)

## Explain

### type

中文：连接类型/访问类型

mysql5.7中join type一共有14种，常用的是all, index, range, ref, eq_ref, const从左到右效率依次增强

| 类型            | 含义                                                         |
| --------------- | ------------------------------------------------------------ |
| system          | 表只有一行                                                   |
| const           | 表最多只有一行匹配，通用用于主键或者唯一索引比较时           |
| eq_ref          | 每次与之前的表合并行都只在该表读取一行，这是除了system，const之外最好的一种，特点是使用=，而且索引的所有部分都参与join且索引是主键或非空唯一键的索引 |
| ref             | 如果每次只匹配少数行，那就是比较好的一种，使用=或<=>，可以是左覆盖索引或非主键或非唯一键 |
| fulltext        | 全文搜索                                                     |
| ref_or_null     | 与ref类似，但包括NULL                                        |
| index_merge     | 表示出现了索引合并优化(包括交集，并集以及交集之间的并集)，但不包括跨表和全文索引。这个比较复杂，目前的理解是合并单表的范围索引扫描（如果成本估算比普通的range要更优的话 |
| unique_subquery | 在in子查询中，就是value in (select...)把形如“select unique_key_column”的子查询替换。PS：所以不一定in子句中使用子查询就是低效的！ |
| index_subquery  | 同上，但把形如”select non_unique_key_column“的子查询替换     |
| range           | 常数值的范围                                                 |
| index           | a.当查询是索引覆盖的，即所有数据均可从索引树获取的时候（Extra中有Using Index）；<br/>b.以索引顺序从索引中查找数据行的全表扫描（无 Using Index）；<br/>c.如果Extra中Using Index与Using Where同时出现的话，则是利用索引查找键值的意思；<br/>d.如单独出现，则是用读索引来代替读行，但不用于查找 |
| All             | 全表扫描                                                     |

[链接](https://blog.csdn.net/dennis211/article/details/78170079)

[链接1](https://blog.51cto.com/lijianjun/1881208)

### extra

using where 代表**MYSQL服务器层将在存储引擎层返回行以后再应用WHERE过滤条件**；

## 网络通信方式

mysql网络通信方式是**半双工通信**，这样使得连接和断开速度非常的快，但是也有缺点就是一旦一方发送数据另一方只有接收完数据之后才能进行响应，这也就无法完成流量控制

* 涉及的参数
  max_allowed_pocket

## Mysql问题

* count(1)、count(*)和count(字段)有什么区别吗？

  count(1)和count(*)的速度是相同的，会统计所有的NULL，速度上比count(字段要快)

  **count(*)和count(1)在innodb引擎中做的优化是：查数量的时候选择使用最小的非聚簇索引，因为体积比聚簇索引小的多（当然只能在没有where条件的情况下使用）**

  count(字段)如果字段是主键的话，速度是最快的，但是不会统计NULL值

* 记录的数据是一条的大小是1Kb，问2层树和三层树分别能够存储多少数据？

  2层是1140*16

  3层是1140\*1140\*16=2000万

  如果是2层的B+树，即存在一个根节点和若干个叶子节点，那么这棵B+树的存放总记录数为：根节点指针数单个叶子节点记录行数。因为单个页的大小为16kb，而一行数据的大小为1kb，也就是说一页可以存放16行数据。然后因为非叶子节点的结构是：“页指针+键值”，我们假设主键ID为bigint类型，长度为8字节（byte），而指针大小在InnoDB源码中设置为6字节（byte），这样一共14字节（byte），因为一个页可以存放16k个byte，所以一个页可以存放的指针个数为16384/14=1170个。因此一个两层的B+树可以存放的数据行的个数为：1170*16=18720（行）

  也就是说第一层的页，即根页（page:3）可以存放1170个指针，然后第二层的每个页（page:4,5,6,7）也可以存放1170个指针。这样一共可以存放1170*1170个指针，所以一共可以存放1170*1170*16=21902400行记录。也就是说一个三层的B+树就可以存放千万级别的数据了。而每经过一个节点都需要IO一次，把这个页数据从磁盘读取到缓存，也就是说读取一个数据只需要三次IO。
  [原文链接](https://blog.csdn.net/qq_35590091/article/details/107361172)

  

  

## 习题

[链接](https://zhuanlan.zhihu.com/p/370224104)

## 链接

[Mysql结构](https://zhuanlan.zhihu.com/p/531731716?utm_source=wechat_session&utm_medium=social&utm_oi=1102878652735262720)

## BOOK

### 高性能Mysql第三版

#### 第五章 创建高性能的索引

* 在查询时候使用and和or的效果甚至不如全表扫描，小数量级别的数据库可能不明显但是数据量变大之后的操作，特别是并发操作性能非常糟糕
  * 当出现服务器对多个索引做相交操作时候（通常有多个And条件），通常意味着需要一个包含所有相关列的多列索引，而不是多个独立的单列索引
  * 当服务器需要对多个索引做联合操作时（通常有多个OR条件），通常消耗大量的CPU和内存资源在算法的缓存，排序和合并操作上。特别是当其中有些索引的选择性不高，需要合并扫描返回大量的数据的时候
  * 总之一句话：不如改成union查询
* 5.3.4 选择合适的索引序列
  * 选择建议
    将选择性最高的列放到索引的最前列

# 消息队列

## 模型讲解

[消息队列的作用](https://juejin.cn/post/6850418106372882446)

# nginx

# haproxy

专业做负载均衡的中间件

[比较全面的回答](https://segmentfault.com/a/1190000039713086)

# k8s

## 结构

### Master Node

链接：https://zhuanlan.zhihu.com/p/292081941

* **API Server**。**K8S的请求入口服务**。API Server负责接收K8S所有请求（来自UI界面或者CLI[命令行工具](https://www.zhihu.com/search?q=命令行工具&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"article"%2C"sourceId"%3A"292081941"})），然后，API Server根据用户的具体请求，去通知其他组件干活。

* **Scheduler**。**K8S所有Worker Node的调度器**。当用户要部署服务时，Scheduler会选择最合适的Worker Node（服务器）来部署。

* **Controller Manager**。**K8S所有Worker Node的监控器**。Controller Manager有很多具体的Controller，在文章[Components of Kubernetes Architecture](https://link.zhihu.com/?target=https%3A//medium.com/%40kumargaurav1247/components-of-kubernetes-architecture-6feea4d5c712)中提到的有Node  Controller、Service Controller、Volume Controller等。Controller负责监控和调整在Worker Node上部署的服务的状态，比如用户要求A服务部署2个副本，那么当其中一个服务挂了的时候，Controller会马上调整，让Scheduler再选择一个Worker Node重新部署服务。

* **etcd**。**K8S的存储服务**。[etcd](https://www.zhihu.com/search?q=etcd&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"article"%2C"sourceId"%3A"292081941"})存储了K8S的关键配置和用户配置，K8S中仅API Server才具备读写权限，其他组件必须通过API Server的接口才能读写数据（见[Kubernetes Works Like an Operating System](https://link.zhihu.com/?target=https%3A//thenewstack.io/how-does-kubernetes-work/)）。

### Slave Node

- **Kubelet**。**Worker Node的监视器，以及与Master Node的通讯器**。Kubelet是Master Node安插在Worker Node上的“眼线”，它会定期向Worker Node汇报自己Node上运行的服务的状态，并接受来自Master Node的指示采取调整措施。
- **Kube-Proxy**。**K8S的网络代理**。私以为称呼为Network-Proxy可能更适合？Kube-Proxy负责Node在K8S的网络通讯、以及对外部网络流量的负载均衡。
- **Container Runtime**。**Worker Node的运行环境**。即安装了容器化所需的软件环境确保容器化程序能够跑起来，比如Docker Engine。大白话就是帮忙装好了Docker运行环境。
- **Logging Layer**。**K8S的监控状态收集器**。私以为称呼为Monitor可能更合适？Logging Layer负责采集Node上所有服务的CPU、内存、磁盘、网络等监控项信息。
- **Add-Ons**。**K8S管理运维Worker Node的插件组件**。有些文章认为Worker Node只有三大组件，不包含Add-On，但笔者认为K8S系统提供了Add-On机制，让用户可以扩展更多定制化功能，是很不错的亮点。

总结来看，**K8S的Master Node具备：请求入口管理（API Server），Worker Node调度（Scheduler），监控和自动调节（Controller Manager），以及存储功能（etcd）；而K8S的Worker Node具备：状态和监控收集（Kubelet），网络和负载均衡（Kube-Proxy）、保障容器化运行环境（Container Runtime）、以及定制化功能（Add-Ons）。**

## k8s网络通信

https://www.zhihu.com/zvideo/1325164415990734848

# MongoDB

## 特点

MongoDB是一个基于分布式文件存储的数据库。由C++语言编写。旨在为WEB应用提供可扩展的高性能数据存储解决方案。
MongoDB是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。
它支持的数据结构非常松散，是类似json的bson格式，因此可以存储比较复杂的数据类型。Mongo最大的特点是它支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。

## 为什么使用MongoDB

（1）MongoDB提出的是文档、集合的概念，使用BSON（类JSON）作为其数据模型结构，其结构是面向对象的而不是二维表，存储一个用户在MongoDB中是这样子的。

```text
{
username :'123',
password:'123',
}
```

使用这样的数据模型，使得MongoDB能在生产环境中提供高读写的能力，吞吐量较于mysql等SQL数据库大大增强。
（2）易伸缩，自动故障转移。易伸缩指的是提供了分片能力，能对数据集进行分片，数据的存储压力分摊给多台服务器。自动故障转移是副本集的概念，MongoDB能检测主节点是否存活，当失活时能自动提升从节点为主节点，达到故障转移。
（3）数据模型因为是面向对象的，所以可以表示丰富的、有层级的数据结构，比如博客系统中能把“评论”直接怼到“文章“的文档中，而不必像myqsl一样创建三张表来描述这样的关系。

[详细链接](https://zhuanlan.zhihu.com/p/87722764)

# Redis

## 高危操作启示

* 线上线下均不允许使用keys 正则进行检索（测试服除外），要使用scan来代替
  因为keys *在进行模糊匹配的时候会引发redis锁，造成redis锁住，一方面redis可能因此宕机，另一方面如果redis正在运行则所有的流量都会转到RDS数据库中，使数据库挂掉

## 常见数据类型和分类

### string

* **set**

  set key value

* **get**

  get key

* keys *

  获取全部的key

* exists key

  是否存在某个key

* append key str

  给key对应的字符串后面加上一个str

* strlen key

  获取key对应的字符的长度

* getrange key 0 3

  获得key对应的字符中[0,3]对应的字符串

* incr view

  view对应的数字加1

* decr view

  view对应的数字减1

* incrby view 10

  可以设置步长，指定增量

* decrby view 10

  可以设定步长，指定减量

* setex key 30 "hello"

  设置key的值为hello，30秒后过期

  如果 key 已经存在， SETEX 命令将会替换旧的值。

* ttl key

  查看key的剩余生存时间

* setnx mykey "redis"

  如果mykey不存在则创建mykey，可以配合看门狗实现分布式锁，看门狗的作用，当自己的时间快不够用的时候，看门狗重新设置过期时间

* getset  db mongodb

  如果存在值，获取原来的值返回，并设置新的价值

### list

list的所有的命令都是以l开头的，用的好的话可以当队列（一端进一端出）、栈、阻塞队列来使用

* lpush	list  value
  将一个或者多个值从列表的左边放进去
* lpop list
  从左边弹出一个
* rpush  list value
  将一个或者多个值从列表的右边放进去
* rpop list
  从右边弹出一个
* llen list
  返回list的长度
* lrem list 1 one(移除列表中的指定个数的字符)
  从list中移除一个one字符
* ltrim list 1 2
  截取列表中的[1,2]
* lrange list 0 -1
  可以显示列表中的所有元素
* lpoplpush srclist deslist
  从srclist的左边弹出一个到deslist的左边

### hash

| 1    | [HDEL key field1 field2](https://www.runoob.com/redis/hashes-hdel.html) 删除一个或多个哈希表字段 |
| ---- | ------------------------------------------------------------ |
| 2    | [HEXISTS key field](https://www.runoob.com/redis/hashes-hexists.html) 查看哈希表 key 中，指定的字段是否存在。 |
| 3    | [HGET key field](https://www.runoob.com/redis/hashes-hget.html) 获取存储在哈希表中指定字段的值。 |
| 4    | [HGETALL key](https://www.runoob.com/redis/hashes-hgetall.html) 获取在哈希表中指定 key 的所有字段和值 |
| 5    | [HINCRBY key field increment](https://www.runoob.com/redis/hashes-hincrby.html) 为哈希表 key 中的指定字段的整数值加上增量 increment 。 |
| 6    | [HINCRBYFLOAT key field increment](https://www.runoob.com/redis/hashes-hincrbyfloat.html) 为哈希表 key 中的指定字段的浮点数值加上增量 increment 。 |
| 7    | [HKEYS key](https://www.runoob.com/redis/hashes-hkeys.html) 获取所有哈希表中的字段 |
| 8    | [HLEN key](https://www.runoob.com/redis/hashes-hlen.html) 获取哈希表中字段的数量 |
| 9    | [**HMGET key field1 field2**](https://www.runoob.com/redis/hashes-hmget.html) 获取所有给定字段的值 |
| 10   | [**HMSET key field1 value1 field2 value2** ](https://www.runoob.com/redis/hashes-hmset.html) 同时将多个 field-value (域-值)对设置到哈希表 key 中。<br />HMSET myhash field1 "Hello" field2 "World" |
| 11   | [HSET key field value](https://www.runoob.com/redis/hashes-hset.html) 将哈希表 key 中的字段 field 的值设为 value 。 |
| 12   | [HSETNX key field value](https://www.runoob.com/redis/hashes-hsetnx.html) 只有在字段 field 不存在时，设置哈希表字段的值。 |
| 13   | [HVALS key](https://www.runoob.com/redis/hashes-hvals.html) 获取哈希表中所有值。 |
| 14   | [HSCAN key cursor [MATCH pattern\] [COUNT count]](https://www.runoob.com/redis/hashes-hscan.html) 迭代哈希表中的键值对。 |

### set

| 1    | [SADD key member1 member2](https://www.runoob.com/redis/sets-sadd.html) 向集合添加一个或多个成员 |
| ---- | ------------------------------------------------------------ |
| 2    | [SCARD key](https://www.runoob.com/redis/sets-scard.html) 获取集合的成员数 |
| 3    | [SDIFF key1 key2](https://www.runoob.com/redis/sets-sdiff.html) 返回第一个集合与其他集合之间的差异。 |
| 4    | [SDIFFSTORE destination key1 key2](https://www.runoob.com/redis/sets-sdiffstore.html) 返回给定所有集合的差集并存储在 destination 中 |
| 5    | [SINTER key1 key2](https://www.runoob.com/redis/sets-sinter.html) 返回给定所有集合的交集 |
| 6    | [SINTERSTORE destination key1 key2](https://www.runoob.com/redis/sets-sinterstore.html) 返回给定所有集合的交集并存储在 destination 中 |
| 7    | [SISMEMBER key member](https://www.runoob.com/redis/sets-sismember.html) 判断 member 元素是否是集合 key 的成员 |
| 8    | [**SMEMBERS** key](https://www.runoob.com/redis/sets-smembers.html) 返回集合中的所有成员 |
| 9    | [SMOVE source destination member](https://www.runoob.com/redis/sets-smove.html) 将 member 元素从 source 集合移动到 destination 集合 |
| 10   | [SPOP key](https://www.runoob.com/redis/sets-spop.html) 移除并返回集合中的一个随机元素 |
| 11   | [SRANDMEMBER key count](https://www.runoob.com/redis/sets-srandmember.html) 返回集合中一个或多个随机数 |
| 12   | [SREM key member1 member2](https://www.runoob.com/redis/sets-srem.html) 移除集合中一个或多个成员 |
| 13   | [SUNION key1 key2](https://www.runoob.com/redis/sets-sunion.html) 返回所有给定集合的并集 |
| 14   | [SUNIONSTORE destination key1 key2](https://www.runoob.com/redis/sets-sunionstore.html) 所有给定集合的并集存储在 destination 集合中 |
| 15   | [SSCAN key cursor [MATCH pattern\] [COUNT count]](https://www.runoob.com/redis/sets-sscan.html) |

### zset

### bitmap

### hyperloglog

### 地理坐标

* hget key
  获得key对应的value

* hgetall key

  获得key下面所有field对应的所有值

* Evalsha 脚本命令
  https://www.runoob.com/redis/scripting-evalsha.html

### 控制信息

* info
  看到所有库的key数量，并显示当前redis的状态

## 为什么那么快

* 纯内存操作
* 核心是非阻塞IO的多路复用机制
* 单线程避免了多线程频繁切换上下文带来的性能问题

## 数据结构

### Redis底层的基本数据结构

* 哈希表O(1)
  采用了渐进式rehash，防止大规模迁移的时候造成Redis线程阻塞无法服务其他请求
* 跳表O(logN)
  在链表的基础上增加了多级索引，通过索引位置的几个跳转，可以实现数据的快速定位
  * 跳表结构
    * 每个节点肯定都有第一层指针（每个节点都在第1层链表里）
    * 如果一个节点由第i（i>=1）指针（即节点已经再第1层到第i层链表中）， 那么他有第（i+1)层指针的概率为P
    * 节点最大的层数不允许超过一个最大值，记为MaxLevel。Redis中skiplist的两个参数为p=1/4 MaxLevel=32
* 双向链表O(N)
* 压缩列表O(N)
  压缩列表是为了解决内存开发而设计的，由连续内存块组成的顺序型数据结构，类似与数组
* 整数数组O(N)

### Redis九大数据结构

* String:字符串

* List:列表

* Hash:哈希表

* set:无序组合

* Sorted Set:有序组合

* bitmap:布隆过滤器

  * 实现原理：它本质上是一个位图，把元素通过多次的hash计算出来的值当作索引，如果索引对应的位图的二进制位为0，说明该元素不存在，如果都为1，该元素可能存在

    ![布隆过滤器](https://img-blog.csdnimg.cn/e7035fbb3ed84222aa16e10bf7cbd616.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA5LiN5piv56ug6bG855qE56ug6bG85ZOl,size_12,color_FFFFFF,t_70,g_se,x_16#pic_center)

* GeoHash:坐标（基于Sorted Set）

* HyperLogLog:统计不重复数据，用于大数据基数统计

* Streams:内存版本的kafaka

### 数据类型和基本数据结构的关系

![](C:\Users\xyk\Desktop\knowledgeTree\redis\images\redis数据结构图.jpg)

## 主从同步

* 如何开启主从复制
  在Redis中，用户通过执行`slaveof`命令或者设置配置文件`slaveof`选项的方式，让一个服务器(从服务器)去复制(`replicate`)另一个服务器(主服务器)，这个复制过程就叫做**主从复制**

  最常见的命令就是：127.0.0.1:6382> slaveof 127.0.0.1 6380

* 主从同步原理

  * 增量同步
    Redis同步的是**指令流**，主节点会将那些对自己的状态产生修改性影响的指令记录在本地内存buffer中，然后异步将buffer中的指令同步到从节点，从节点一边执行同步的指令流来达到和主节点一样的状态，一边向主节点反馈自己同步到了哪里（buffer）。

    因为内存的 buffer 是有限的，所以 Redis 主库不能将所有的指令都记录在内存 buffer 中。Redis 的复制内存 buffer 是一个定长的环形数组，如果数组内容满了，就会从头开始覆盖前面的内容。

    如果因为网络状况不好，从节点在短时间内无法和主节点进行同步，那么当网络状况恢复时，Redis 的主节点中那些没有同步的指令在 buffer 中有可能已经被后续的指令覆盖掉了，从节点将无法直接通过指令流来进行同步，这个时候就需要用到更加复杂的同步机制——快照同步。

  * 快照同步
    快照同步是一个非常耗费资源的操作，它首先需要在主库上进行一次 bgsave 将当前内存的数据全部快照到磁盘文件中，然后再将快照文件的内容全部传送到从节点。从节点将快照文件接受完毕后，立即执行一次全量加载，加载之前先要将当前内存的数据清空。加载完毕后通知主节点继续进行增量同步。

    在整个快照同步进行的过程中，主节点的复制 buffer 还在不停的往前移动，如果快照同步的时间过长或者复制 buffer 太小，都会导致同步期间的增量指令在复制 buffer 中被覆盖，这样就会导致快照同步完成后无法进行增量复制，然后会再次发起快照同步，如此极有可能会陷入快照同步的死循环。

    所以**务必配置一个合适的复制 buffer 大小参数，避免快照复制的死循环**。
    当从节点刚刚加入到集群时，它必须先要进行一次快照同步，同步完成后再继续进行增量同步。

  * 无盘复制主节点在进行快照同步时，会进行很重的文件 IO 操作，特别是对于非 SSD 磁盘存储时，快照会对系统的负载产生较大影响。特别是当系统正在进行 AOF 的 fsync 操作时如果发生快照，fsync 将会被推迟执行，这就会严重影响主节点的服务效率。
    所以从 Redis 2.8.18 版开始支持无盘复制。所谓无盘复制是指主服务器直接通过 SOCKET 将快照内容发送到从节点，生成快照是一个遍历的过程，主节点会一边遍历内存，一遍将序列化的内容发送到从节点，从节点还是跟之前一样，先将接收到的内容存储到磁盘文件中，再进行一次性加载。

  * Wait指令（同步复制）
    Redis 的复制是异步进行的，wait 指令可以让异步复制变身同步复制，确保系统的强一致性(不严格)。wait 指令是Redis3.0 版本以后才出现的。
    wait 提供两个参数，第一个参数是从库的数量 N，第二个参数是时间 t，以毫秒为单位。它表示等待 wait 指令之前的所有写操作同步到 N 个从库(也就是确保 N 个从库的同步没有滞后)，最多等待时间 t。如果时间 t=0，表示无限等待直到 N 个从库同步完成达成一致。
    假设此时出现了网络分区，wait 指令第二个参数时间 t=0，主从同步无法继续进行，wait 指令会永远阻塞，Redis 服务器将丧失可用性。

* 为什么使用主从复制
  主从复制可以实现读写分离和数据备份

## 持久化

### RDB-全量持久化

* 过程
  redis database 将某一个时刻的内存快照，以二进制的方式写入磁盘实际操作过程是，fork一个子进程，读取完数据后用二进制压缩

* 命令

  * save
    save命令来触发，使得redis处于阻塞状态，会阻塞其他客户端发来的命令，生产环境一定要禁用

  * bgsave
    bgsave这个命令会fork一个进程来执行这个操作，但是在fork的时候会阻塞进程，直到阻塞完成。

    父进程继续处理client请求，子进程负责将内存内容写入到临时文件。由于os的写时复制机制（copy on write)父子进程会共享相同的物理页面，当父进程处理写请求时os会为父进程要修改的页面创建副本，而不是写共享的页面。所以子进程的地址空间内的数据是fork时刻整个数据库的一个快照。

### AOF-增量持久化

Append Only File 保存的是操作日志

## 高可用

### 哨兵模式

* 哨兵模式的功能
  * 监控
    sentinel会不断的检查主从服务器是否正常工作

  * 提醒
    当监控到哪个服务器出现问题的时候，通过API通知管理员

  * 故障转移

    当一个主服务器不能正常工作时候，Sentinel会开始一次自动故障转移操作，他会将失效主服务器的其中一个从服务器升级为新的主服务器，并让失效主服务器的其他从服务器改为复制新的主服务器，当客户端视图连接失效的主服务器时，集群也会向客户端返回新主服务器的地址，使得集群可以使用新主服务器代替失效服务器

  * 配置中心

* 哨兵模式的缺点
  哨兵的配置略微复杂，并且性能和高可用性等各方面表现一般，特别是在主从切换的瞬间存在访问瞬断的情况，而且哨兵模式只有一个主节点对外提供服务，没法支持很高的并发，且单个主节点内存也不宜设置得过大，否则会导致持久化文件过大，影响数据恢复或主从同步的效率

### sharding 分片模式

redis cluster出现之前业界主要采用的方法。

* 实现方法
  主要思想是采用哈希算法将redis数据的key进行散列，通过hash函数，特定的key会映射到特定的redis节点上。

* 优点

  优势在于非常简单，服务端的redis实例彼此独立，相互无关联，每个redis实例像单服务器一样运行，非常容易线性拓展，线性灵活性很强。

* 缺点
  拓展非常困难，拓展一个需要全部从新hash

### cluster 集群模式

* redis集群原理Redis Cluster 将所有数据划分为 16384 个 slots(槽位)，每个节点负责其中一部分槽位。槽位的信息存储于每个节点中。
  **当 Redis Cluster 的客户端来连接集群时，它也会得到一份集群的槽位配置信息并将其缓存在客户端本地**。这样当客户端要查找某个 key 时，可以直接定位到目标节点。同时因为槽位的信息可能会存在客户端与服务器不一致的情况，还需要纠正机制来实现槽位信息的校验调整。[原文](https://blog.csdn.net/weixin_44795847/article/details/123114968)

* 优点
  * 无中心架构、每个节点保存数据和整个集群的状态，每个节点都和其他所有节点连接。官方要求：至少6个节点才可以保证高可用，即3主3从；拓展性强，更好做高可用
  * 各个节点会互相通信，采用gossip协议交换节点元数据信息
  * 数据分散存储在各个节点上
  * 客户端不需要连接集群所有节点，连接集群中任何一个可用节点即可
  * 高性能，客户端直连redis服务，免去了proxy代理的损耗
* 缺点
  * 运维相对麻烦，需要手动导入导出槽
  * 不支持批量操作
  * 分布式逻辑和存储模块耦合

### cluster和哨兵节点的不同

* 哨兵模式是每个节点持有全量的数据，且数据保持一致，目的为系统Redis高可用
* 集群模式，每个节点主数据不同，是数据的子集，利用多台服务器构建集群提高超大规模数据处理能力，突破单台redis的存储极限，同时提供高可用支持（slaver提供）

## 分布式锁

* 实现步骤

  * 指定一个 key 作为锁标记，存入 Redis 中，指定一个 唯一的用户标识 作为 value。

    ```redis
    setnx(String key, String value)
    ```

  * 当 key 不存在时才能设置值，确保同一时间只有一个客户端进程获得锁，满足 互斥性 特性。

  * 设置一个过期时间，防止因系统异常导致没能删除这个 key，满足 防死锁 特性。

    ```
    expire(String key, int seconds)
    ```

  * 当处理完业务之后需要清除这个 key 来释放锁，清除 key 时需要校验 value 值，需要满足 只有加锁的人才能释放锁 。

    ```
    del(String key)
    ```

    

##  缓存问题

### 三大缓存问题

* 缓存雪崩
  
  * 问题
    大量的key同时失效，请求都落到了数据库上
  
  * 解决方案
  
    缓存时间设置为随机，防止大量数据同时过期
  
* 缓存击穿
  
  * 问题
    缓存中没有数据库中有的数据，请求都落到了数据库上
  * 解决方案
    * 热点数据设置永不过期
    * 加锁同时只有一个线程去数据库查询其他线程用cas自旋锁来自旋
  
* 缓存穿透
  
  * 问题
    缓存和数据库中都没有的数据，请求都落到了数据库上
  * 解决方案
    * 接口层增加检校，用户鉴权、id基础检校等不合法的直接拦截
    * 缓存或者是数据库中都取不到的数据，可以将key-value对写成key-null
    * 采用布隆过滤器，一定不存在的数据一定会被这个bitmap拦截掉

### 缓存过期策略

* 定时过期

  每个设置过期时间的key都需要创建一个定时器，到过期时间就会立即清除。该策略可以立即清理过期的数据，堆内存很友好，但是会占用大量的cpu资源去处理过期数据，从而影响缓存的响应时间和吞吐量。

* 惰性过期
  只有当访问一个key的时候，才会判断该key是否已经过期，过期则清除。该策略可以最大化地节省cpu资源，但是很消耗内存、许多的过期数据都还存在内存中。

* 定期过期
  每隔一定的时间，会扫描一定数量的数据库的expires中zi'dian一定数量的key（随机的），并清除其中已经过期的key。该策略是定时过期和惰性过期的折衷方案。
  
* 在用的方案：定期删除和惰性删除相结合

### 内存淘汰策略



### Redis的使用场景

* 会话缓存，保存用户的session等信息

* 全页缓存、缓存整个网页页面

* 消息队列
  Redis 中list的数据结构实现是双向链表，所以可以非常便捷的应用于消息队列（生产者 / 消费者模型）。消息的生产者只需要通过lpush将消息放入 list，消费者便可以通过rpop取出该消息，并且可以保证消息的有序性。如果需要实现带有优先级的消息队列也可以选择sorted set。而pub/sub功能也可以用作发布者 / 订阅者模型的消息。无论使用何种方式，由于 Redis 拥有持久化功能，也不需要担心由于服务器故障导致消息丢失的情况。

  List提供了两个阻塞的弹出操作：blpop/brpop，可以设置超时时间

  blpop：blpop key1 timeout 移除并获取列表的第一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
  brpop：brpop key1 timeout 移除并获取列表的最后一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
  上面的操作。其实就是java的阻塞队列。学习的东西越多。学习成本越低

  队列：先进先除：rpush blpop，左头右尾，右边进入队列，左边出队列
  栈：先进后出：rpush brpop

* 排行榜

* 计数器、文章的阅读量、微博的点赞数等，允许有一定的延迟，先写入Redis再定时同步到数据库

* 限流-每个用户对应一个key进行计数，计数超过一定的数量则不允许进入

## 面试题

* redis内存满了如何处理
  * 增加内存
  * 搭建redis集群
  * 内存淘汰策略
  
* redis热key造成服务雪崩的问题

* Redis大Key问题
  
  * 定义：
    string长度大于10K，list长度大于10240认为是big bigkeys
  
  * 场景
  
    * 热门话题下的评论、答案排序场景
    * 大V的粉丝列表
    * 使用不恰当，或者对业务预估不准去、不及时进行处理垃圾数据等
  
  * 影响
  
    * 内存不均
      集群模式在slot分片均匀情况下，会出现数据和查询倾斜情况，部分有大key的Redis节点占用内存多，QPS高。
    * 阻塞请求
      redis为单线程，单value较大的情况下，读写需要较长的处理时间，会阻塞后续的请求处理；大key相关的删除或者自动过期时候，会出现qps突降或者突升的情况，极端情况下，会造成主从复制异常，Redis服务阻塞无法响应请求
    * 阻塞网络
      单value较大时候会占用服务器网卡较多带宽，可能会影响该服务器上的其他Redis实例或者应用
  
    [详情网址](https://blog.csdn.net/xushiyu1996818/article/details/121326089)
  
    
  
    
  
    

# Java

## 基础知识

### 关键字

1. public：可以被所有其他类所访问
2. private：只能被自己访问和修改
3. protected：自身、子类及同一个包中类可以访问
4. default：同一包中的类可以访问，声明时没有加修饰符，认为是friendly。

### 接口和抽象类

1. 接口可以多继承，抽象类不行
2. 接口定义方法，就不能实现，默认是public abstract, 而抽象类可以实现部分方法
3. 接口中基本数据类型为public static final 并且需要给出初始值，而抽象类不是

### 数组

* 初始化方法
  * 数据类型 [ ] 数组名称 = new 数据类型[数组长度]
  * 数据类型 [ ] 数组名称 = new 数据类型[ ]{元素1，元素2，元素3…}
  * 数据类型 [ ] 数组名称 = {元素1，元素2，元素3…}

### 对象

#### 对象的创建方式

* new

* Class.newInstance()
  Person p = Person.Class.newInstance()

* Constructor.newInstance()
  Constructor<Person> constructor = Person.Class.getConstructor();

  Person p = constructor.newInstance();

* clone()

  Person p = new Person();

  Person p1 = p.clone();

* 反序列化
  Person p1 = new Person();

  byte[] bytes = SerializationUtils.serialize(p1);

  Person p2 = (person) SerializationUtils.deserialize(bytes);

#### 对象的实体结构

![](C:\Users\xyk\Desktop\knowledgeTree\java\images\对象内存结构图.png)

* 对象头

  * markword结构
    markword中的内容可能是：

    HashCode、GC分代年龄、锁状态标志、线程持有的锁、偏向线程ID、偏向时间戳等

    ![](C:\Users\xyk\Desktop\knowledgeTree\java\images\markword的结构.jpg)

  * Klass指针
    对象指向类元数据的指针，虚拟机通过这个指针确定该对象是哪个类的实例

  * 数组长度
    有/无，只有数组对象有数组长度，如果对象为数组，对象头中必须记录数组长度

* 实例数据
  是对象的有效数据，包括对象中的各种类型的属性、方法和字段的内容

* 填充数据

### 异常

throw和throws：

throw是一个动作，作用于方法内，用于主动抛出异常

throws作用于方法声明上，声明该方法有可能会抛出某些异常

## 泛型

### 泛型分类

* 方法泛型
* 接口泛型
* 类泛型

### 经典示例

```java
public class ArrayList<T> {
    private T[] array;
    private int size;
    public void add(T e) {...}
    public void remove(int index) {...}
    public T get(int index) {...}
}
```

## 反射

### 概念

* 程序可以访问、检测和修改它本身状态或行为的能力，即自描述和自控制。
* 可以在运行时加载、探知和使用编译期间完全未知的类。
* 给java插上动态语言特性的翅膀，弥补强类型语言的不足。

### 作用

* 在运行中分析类的能力
* 在运行中查看和操作对象
  * 基于反射自由创建对象
  * 反射构建出无法访问的成员变量
  * 调用不可访问的方法
* 实现通用的数组操作代码
* 类似函数指针的功能

### 反射中的关键类

* Class类

  JVM为每个对象都保留其类型标识信息(Runtime Type Identification)三种获取方式

  * Class c1 = s1.getClass()
  * Class c2 = Class.forName("java.lang.Strinig")
  * Class c3 = String.Class

* Field类

  * getFields()
    获取本类以及父类所有的public字段

  * getDeclareFields()

    获得本类所有声明的字段

* Constructor构造方法

  * getConstructors()

* Method成员方法

  * getMethod()
  * getDeclareMethods()

* 获取父类Class
  getSuperClass()

* 获取修饰符

  getModifiers()

* 利用反射生成对象和调用方法的示例

  ```java
  public class A {
        public void hello() {
               System.out.println("hello from A");
        }
  }
  public class NO {
  	public static void main(String []s) throws InstantiationException, 
  		IllegalAccessException, ClassNotFoundException, NoSuchMethodException, 
  		SecurityException, IllegalArgumentException, InvocationTargetException {
  		
  		              Object obj6 = Class.forName("reflection.A").newInstance();
  		              Method m = Class.forName("reflection.A").getMethod("hello");
  		              m.invoke(obj6);
  		}
  	}
  ```

  

### 创建对象的几种方法

* 直接使用new

* 使用克隆

  ```java
  package reflection;
  public class B implements Cloneable {
  		     public void hello(){
  		            System.out.println("hello from B");
  		     }
  		     protected Object clone() throws CloneNotSupportedException{
  		            return super.clone();
  		     }
  }
  package reflection;
  		public class cloneB {
  		public static void main(String [] s) throws CloneNotSupportedException {
  		              B obj2 = new B();
  		              obj2.hello();
  		              B obj3 = (B) obj2.clone();
  		              obj3.hello();
  		       }
  }
  ```

* 序列化和反序列化

* 反射
  

## Collection

### List

* 分类
  * ArrayList
  * Vector（线程安全）
  * LinkedList

* 常用方法
  * 在末尾添加一个元素：`boolean add(E e)`
  * 在指定索引添加一个元素：`boolean add(int index, E e)`
  * 删除指定索引的元素：`E remove(int index)`
  * 删除某个元素：`boolean remove(Object e)`
  * 获取指定索引的元素：`E get(int index)`
  * 获取链表大小（包含元素的个数）：`int size()`
* * 

### Queue

* 分类
  * LinkedList
  * PriorityQueue
  * Deque
    * Stack

* 常用方法

  * `int size()`：获取队列长度；
  * `boolean add(E)`/`boolean offer(E)`：添加元素到队尾；
  * `E remove()`/`E poll()`：获取队首元素并从队列中删除；
  * `E element()`/`E peek()`：获取队首元素但并不从队列中删除。

* Deque拓展
  虽然也有Queue中的相应的方法，但是不建议使用

  * 特有方法

    | 添加元素到队尾     | add(E e) / offer(E e)  | addLast(E e) / offerLast(E e)   |
    | ------------------ | ---------------------- | ------------------------------- |
    | 取队首元素并删除   | E remove() / E poll()  | E removeFirst() / E pollFirst() |
    | 取队首元素但不删除 | E element() / E peek() | E getFirst() / E peekFirst()    |
    | 添加元素到队首     | 无                     | addFirst(E e) / offerFirst(E e) |
    | 取队尾元素并删除   | 无                     | E removeLast() / E pollLast()   |
    | 取队尾元素但不删除 | 无                     | E getLast() / E peekLast()      |

### Set

​	无序、唯一

* 分类
  * HashSet
  * LinkedHashSet
    底层数据结构是链表和哈希表。(FIFO插入有序,唯一)
    1.由链表保证元素有序
    2.由哈希表保证元素唯一
  * TreeSet
    底层数据结构是红黑树。(唯一，有序)
* 常用方法
  * 将元素添加进`Set`：`boolean add(E e)`
  * 将元素从`Set`删除：`boolean remove(Object e)`
  * 判断是否包含元素：`boolean contains(Object e)`

## Map

### HashMap

无序、不可重复的

key和value都可以为空

* 常用方法
  * containsKey(key)
  * get(key)
  * put(key, value)
* 遍历方法
  * for(String key: map.keySet()){map.get(key);}
  * for(Map.Entry<String, Integer> entry:map.entrySet()){entry.getKey();}

* 扩容为什么都是以二倍的方式进行的？

  * 在寻找插入数组的时候可以用&位运算代替%运算
  * hash值不用再运算一遍

* HasMap中红黑树和链表是可以同时存在的嘛？

  * 可以

* 桶中的链表转化为红黑树的条件？

  * 当链表长度大于8的时候，总元素个数大于64的时候会进行扩容

    当红黑树的长度小于6的时候会从红黑树退化为链表

### HashTable（线程安全）

线程安全，使用了Sychronized来进行同步

### TreeMap

* 特点
  * 不允许出现重复的key
  * 可以插入null键，null值
  * 可以对元素进行排序
  * 无序集合（插入和遍历的顺序不一样）

### linkedHashMap

支持有序的插入和遍历

## 多线程

### 线程三大基本概念

（并发的三大特性）

* 原子性
  在java中，对基本数据类型的变量的读取和赋值操作是原子性操作，这些操作时不可被中断，要么执行，要么不执行
* 可见性
  在一个共享变量被volatile修饰时候，他会保证修改的值会立刻被更新到主存，当以后其他线程需要读取时候，他会去内存中读取新值
* 有序性
  java内存模型中，允许编译器和处理器对指令进行重排序，但是重排序过程不会影响到单线程程序的运行，却会影响到多线程并发执行的正确性

### 线程的状态

* 新建
* 就绪
* 运行
* 阻塞
  * 等待阻塞
    运行线程执行wait方法，需要依靠其他线程调用notify或者notifyAll方法才能唤醒。
  * 同步阻塞
    运行的线程获取对象的同步锁时候，若该同步锁被其他线程占用，则JVM会把该线程放入锁池中。
  * 其他阻塞（sleep等）
    线程运行sleep或者join方法，或者发出I/O请求时，JVM会把该线程设置为阻塞状态，sleep是thread类的方法。
* 死亡

### 线程的分类

* daemon线程-后台守护线程（如果主线程结束，则内部线程结束）

  使用Thread one.setDaemon(true)可以设置为后台线程

* 非daemon线程-非后台守护线程

### 线程的基本方法

* 对象的基本方法
  * wait()
    释放锁
  * notify()
  * notifyAll()
* 线程的基本方法
  * sleep()
    不释放锁
  * join()
    等待其他线程终止
  * yield()
    让出当前的cpu

### 多线程的分类

* 实现Runnable接口

  * 继承Thread类

    ```java
    public class main extends Thread{
    
    ​    public void run(){
    ​        System.*out*.println("start");
    ​    }
    
    ​    public static void main(String[] args){
    ​        main m = new main();
    ​        System.*out*.println("hello");
    ​        m.start();
    ​    }
    }
    ```

  * 实现Runnable接口

* 实现Callable接口
  Callable比如Runnable实现并没有什么大的差别，一个能带返回值，一个无法带

  ```java
  class Task implements Callable<String> {
      public String call() throws Exception {
          return longTimeCalculation(); 
      }
  }
  
  ExecutorService executor = Executors.newFixedThreadPool(4); 
  // 定义任务:
  Callable<String> task = new Task();
  // 提交任务并获得Future:
  Future<String> future = executor.submit(task);
  // 从Future获取异步执行返回的结果:
  String result = future.get(); // 可能阻塞
  ```

* 使用线程池
  * 线程池的工作流程
    * 判断**线程池里的核心线程**（核心线程数量）是否都在执行任务，如果不是（核心线程空闲或者还有核心线程没有创建）则创建一个新的工作线程来执行任务。如果核心线程都在执行任务，则进入下个流程
    * 线程池判断工作队列是否已经满了，如果工作队列没有满，则将提交的任务存储在这个工作队列里。如果工作队列满了，则进入下一个流程
    * 判断**线程池里的线程**（是最大线程数）是否都处于工作状态，如果没有，则创建一个新的工作线程来执行任务。如果已经满了，则交给饱和策略来处理这个任务
    * 流程图
      ![](C:\Users\xyk\Desktop\knowledgeTree\java\images\线程池的工作流程图.JPG)

### 线程池

* 线程池函数

  ```
  public ThreadPoolExecutor(int corePoolSize,
  							int maxmumPoolSize,
  							long keepAliveTime,
  							TimeUnit unit,
  							BlockingQueue<Runnable> workQueue,
  							RejectedExecutionHandler handler)
  ```

  * corePoolSize：线程池核心线程数量
  * maximumPoolSize:线程池最大线程数量
  * keepAliverTime：当活跃线程数大于核心线程数时，空闲线程最大存活时间
  * unit：存活时间的单位
  * workQueue：存放任务的队列
    * ArrayBlockingQueue基于数组的有界阻塞队列，按FIFO排序。
    * LinkedBlockingQueue基于链表的无界阻塞队列，按照FIFO
    * SynchronousQuene一个不缓存任务的阻塞队列，生产者放入一个任务必须等到消费者取出这个任务
    * PriorityBlockingQueue具有优先级的无界阻塞队列
  * handler：超出线程范围和队列容量的任务的处理程序
    * AbortPolicy:不处理丢弃掉
    * CallerRunsPolicy: 由调用线程处理该任务（谁调用，谁处理）
    * DiscardPolicy: 忽视，什么都没有发生
    * DiscardOldestPolicy: 丢弃线程队列的旧的任务，将新的任务添加
      当任务被拒绝添加时，会抛弃任务队列中最旧的任务也就是最先加入队列的，再把这个新任务添加进去。在rejectedExecution先从任务队列种弹出最先加入的任务，空出一个位置，然后再次执行execute方法把任务加入队列。
  
* 如何设置线程池参数

  * 需要根据几个值来决定
    * tasks: 每秒的任务数
    * taskcost:每个任务花费的时间
    * responsetime:系统允许容忍的最大响应时间，假设为1S
  
* 使用方法

  ```java
  ExecutorService pool  = new Executors.newCachedThreadPool();
  
  MyTask one = new MyTask();
  
  pool.execute(one);
  
  pool.shutdown();
  ```

  

### 线程同步

* 互斥同步

  * 临界区Synchronized、ReentrantLock

  * 信号量semaphoreSemaphore 有两个构造函数，参数为许可的个数 permits 和是否公平竞争 fair。通过 acquire 方法能够获得的许可个数为 permits，如果超过了这个个数，就需要等待。当一个线程 release 释放了一个许可后，fair 决定了正在等待的线程该由谁获取许可，如果是公平竞争则等待时间最长的线程获取，如果是非公平竞争则随机选择一个线程获取许可。不传 fair 的构造函数默认采用非公开竞争。

    Semaphore(int permits)

    Semaphore(int permits, boolean fair)

    Semaphore.acquire();//获取一个执行许可

    Semaphore.release();//释放一个执行许可

  * 互斥量mutex

* 非阻塞同步
  CMS自旋锁

* 无同步方案
  ThreadLocal变量，每个线程有一份各自独占的变量

### 线程通信

* volatile
* wait/notify()
* join()
* 管道流

## 锁

### 锁的分类

* 按照乐观悲观来分

  * 乐观锁

    ​	乐观锁是一种乐观思想，即认为读多写少，遇到并发写的可能性低，每次去拿数据的时候都认为别人不 

    会修改，所以不会上锁，但是在更新的时候会判断一下在此期间别人有没有去更新这个数据，采取在写 

    时先读出当前版本号，然后加锁操作（比较跟上一次的版本号，如果一样则更新），如果失败则要重复 

    读-比较-写的操作。 

    ​	java中的乐观锁基本都是通过CAS操作实现的，CAS是一种更新的原子操作，比较当前值跟传入值是否一 样，一样则更新，否则失败。 

  * 悲观锁

    ​	悲观锁是就是悲观思想，即认为写多，遇到并发写的可能性高，每次去拿数据的时候都认为别人会修 

    改，所以每次在读写数据的时候都会上锁，这样别人想读写这个数据就会block直到拿到锁。

    ​	java中的悲 观锁就是Synchronized,AQS框架下的锁则是先尝试cas乐观锁去获取锁，获取不到，才会转换为悲观 锁，如RetreenLock。 

* 按照公平性

  * 公平锁
    自旋锁

  * 非公平锁

    ReentrantLock、Sychronized

### java中的锁

* 偏向锁

  1. java6中的一项多线程优化。他会偏向于第一个访问锁的线程，如果在运行过程中，同步锁只有一个线程访问，不存在多线程争用的情况，则线程不需要触发同步的，这种情况下，就会给线程加一个偏向锁。如果在运行过程中，遇到了其他线程抢占锁，则持有偏向锁的线程会被挂起，JVM会消除它身上的偏向锁，将锁恢复到标准的轻量级锁。

  2. 生成时刻：一个线程获取锁时候会由无锁升级为偏向锁

  3. 相当于给对象贴了一个标签（将自己的线程id存入到对象中），下次我再进来时，发现标签是我的，我就可以继续使用了

* 轻量级锁

* 重量级锁

### Synchronized

* 实现原理

  (1) java里面的每个对象JVM底层都会为它创建一个监视器monitor，这个监视器就类似于一个锁，哪个线程持有这个monitor的操作权，就相当于获得了锁

  (2) 用synchronized 修饰的代码或者方法，底层会生成两条指令分别是monitorenter,monitorexit

  (3) 进入synchronized的代码块之前会执行monitor enter指令，去申请**monitor监视器的操作权**，如果申请成功了，就相当于获取到了锁。如果已经有别的线程申请成功monitor了，这个时候它就得等着，等别的线程执行完synchronized里面的代码之后就会执行monitorexit指令释放monitor监视器，这样其它在等待的线程就可以再次申请获取monitor监视器了。

* 作用范围

  * 修饰普通方法，锁住的是对象的实例（this）

  * 修饰静态方法，锁住的是Class实例。又因为Class的相关数据存储在永久代（java8中是metaspace），永久代是全局共享的，因此静态方法锁住的相当于一个类的全局锁，会锁住所有调用该方法的线程

  * 修饰代码块，作用于当前对象实例，需要指定加锁对象

    ```java
    //Synchronized可以锁住普通方法，也可以锁住一个类，那么它锁的粒度能否更小呢？是的，它还能锁住一段简易的代码块。那么Synchronized如何定义一段代码块呢？其实定义一下作用的对象，然后将代码用括号{ }包裹起来就可以了：
    public class fancySyncTest {
        public synchronized void method1(){
            synchronized (this) {
                // 逻辑代码
            }
        }
    }
    //代码块锁住的对象就是后面括号里的东西。比如这里的synchronized (this)，意味着只有当前对象才可以访问这段代码块，你也可以定义为其它对象。
    ```

    

### ReentrantLock

* 定义
  可重入锁就是可以重新反复进入的锁，仅限于当前线程

* 常用方法

  * lock()

    1. 锁空闲：直接获取锁并返回，同时设置锁持有者数量为：1

    2. 当前线程持有锁：直接获取锁并返回，同时锁持有者数量递增1
    3. 其他线程持有锁：当前线程会休眠等待，直至获取锁为止

  * unlock()
    减1，到0的时候释放锁

  * tryLock()
    尝试获取锁，获取成功返回true，获取失败返回false;

### AQS抽象队列同步器

AbstractQueuedSynchronizer类如其名，抽象的队列式的同步器，AQS定义了一套多线程访问共享资源的同步框架，许多同步类实现都依赖于它

java.util.concurrent：提供大部分关于并发的接口和类
如BlockingQueue、Callable、ConcurrentHashMap、ExecutorService、 Semaphore等。

* atomic包

  * java.util.concurrent.atomic：提供所有原子操作的类， 如AtomicInteger, AtomicLong等；

    ```
    AtomicInteger cnt;
    
    cnt.getAndIncrement();
    ```

* locks包

  * java.util.concurrent.locks:提供锁相关的类, 如Lock, ReentrantLock, ReadWriteLock, Condition等；

* concurrent包中的类
  * CopyOnWriteArrayList
  * CopyOnWriteArraySet
  * ConcurrentHashMap
  * ArrayBlockingQueue
    生产者与消费者队列使用put()和take()

## 网络编程

### 同步和阻塞

* 阻塞与非阻塞、同步与异步的区别
  * 同步和异步的概念与**消息的通知机制**有关。
    对于消息处理者而言，在同步的情况下，由处理消息者自己去等待消息是否被触发；在异步的情况下，由触发机制来通知处理消息者，然后进行消息的处理。同步和异步仅仅是关于所关注的消息如何通知的机制，而不是处理消息的机制。
  * 阻塞和非阻塞与**消息的处理机制**有关。
    阻塞模式是指在指定套接字上调用函数执行操作时，在没有完成操作之前，函数不会立即返回。非阻塞模式是指在指定套接字上调用函数执行操作时，无论操作是否完成，函数都会立即返回。

## IO 

### 按照阻塞分类

* BIO-阻塞同步

* NIO-非阻塞同步
  
  核心组件
  * Channels
    * 既可以从通道中读取数据，又可以写数据到通道。但是流的读写通常是单向的
    * 通道可以异步读写
    * 通道中的数据总是要先读到一个Buffer，或者总是要从一个Buffer中写入
  * Buffers
    * 作用：用于NIO Channel交互。我们从Channel中读取数据到buffers里，从buffer把数据写入到Channels
    * Buffer本质上是一块内存
    * Buffer中有三个属性特别重要：capacity容量、position位置、limit限制
    * 常用方法
      * clear()
      * flip()
      * rewind()
        模式切换，将Buffer从写模式切换到读模式
      * Buffer position(int newPosition)
  * Selectors
  
* AIO-彻底的异步通信

* 经典例子
  假设有这么一个场景，有一排水壶（客户）在烧水。

  - AIO的做法是，每个水壶上装一个开关，当水开了以后会提醒对应的线程去处理。
  - NIO的做法是，叫一个线程不停的循环观察每一个水壶，根据每个水壶当前的状态去处理。
  - BIO的做法是，叫一个线程停留在一个水壶那，直到这个水壶烧开，才去处理下一个水壶。
    可以看出AIO是最聪明省力，NIO相对省力，叫一个人就能看所有的壶，BIO最愚蠢，劳动力低

## JVM

### 结构

![](C:\Users\xyk\Desktop\knowledgeTree\java\images\jvm内存模型.JPG)

* java虚拟机栈
  局部变量表、操作数栈、动态连接、方法出口等
* 本地方法栈
  类似于java虚拟机栈是为本地方法提供服务
* java堆
  所有的实例对象以及数据
* 方法区
  常量、常量池、static、Class（版本-字段-方法-接口）



### 调优

#### 性能定义

- 吞吐量 - 指不考虑 GC 引起的停顿时间或内存消耗，垃圾收集器能支撑应用达到的最高性能指标。
- 延迟 - 其度量标准是缩短由于垃圾啊收集引起的停顿时间或者完全消除因垃圾收集所引起的停顿，避免应用运行时发生抖动。
- 内存占用 - 垃圾收集器流畅运行所需要的内存数量。

#### GC 优化的两个目标

1. 将进入老年代的对象数量降到最低
2. 减少 Full GC 的执行时间
   Full GC的执行时间比Minor GC要长很多，因此，如果在Full GC上花费过多的时间（超过1s），将可能出现超时错误

#### 调优命令

##### jmap

jmap 即 JVM Memory Map。

jmap 用于生成 heap dump 文件。

如果不使用这个命令，还可以使用 `-XX:+HeapDumpOnOutOfMemoryError` 参数来让虚拟机出现 OOM 的时候，自动生成 dump 文件。

jmap 不仅能生成 dump 文件，还可以查询 finalize 执行队列、Java 堆和永久代的详细信息，如当前使用率、当前使用的是哪种收集器等。

* 命令格式：

  ```
  jmap [option] LVMID
  ```

  * option 参数：
    * dump - 生成堆转储快照
    * finalizerinfo - 显示在 F-Queue 队列等待 Finalizer 线程执行 finalizer 方法的对象
    * heap - 显示 Java 堆详细信息
    * histo - 显示堆中对象的统计信息
    * permstat - to print permanent generation statistics
    * F - 当-dump 没有响应时，强制生成 dump 快照

* 使用示例

  * 示例：jmap -dump PID生成堆快照

  * 示例： jmap -heap PID查看指定进程的堆信息

##### jstack

用于生成java虚拟机当前时刻的线程快照

线程快照是当前 java 虚拟机内每一条线程正在执行的方法堆栈的集合，生成线程快照的主要目的是定位线程出现长时间停顿的原因，如线程间死锁、死循环、请求外部资源导致的长时间等待等。

线程出现停顿的时候通过 jstack 来查看各个线程的调用堆栈，就可以知道没有响应的线程到底在后台做什么事情，或者等待什么资源。如果 java 程序崩溃生成 core 文件，jstack 工具可以用来获得 core 文件的 java stack 和 native stack 的信息，从而可以轻松地知道 java 程序是如何崩溃和在程序何处发生问题。另外，jstack 工具还可以附属到正在运行的 java 程序中，看到当时运行的 java 程序的 java stack 和 native stack 的信息, 如果现在运行的 java 程序呈现 hung 的状态，jstack 是非常有用的。

* 命令格式：

  ```
  jstack [option] LVMID
  ```
  * option 参数：
    * `-F` - 当正常输出请求不被响应时，强制输出线程堆栈
    * `-l` - 除堆栈外，显示关于锁的附加信息
    * `-m` - 如果调用到本地方法的话，可以显示 C/C++的堆栈

* 示例：

  * jstack [pid] 
        查找进程死锁
        "Thread-1" 线程名  
        prio=5 优先级=5 
        tid=0x000000001fa9e000 线程id 
        nid=0x2d64 线程对应的本地线程标识nid 
        java.lang.Thread.State: BLOCKED 线程状态

##### jstat

jstat，用于监视虚拟机运行时状态信息的命令，它可以显示出虚拟机进程中的类装载、内存、垃圾收集‘JIT编译等运行数据。

* 命令格式：

  ```
  jstat [option] LVMID [interval] [count]
  ```
  * 参数：
    * [option] - 操作参数
    * LVMID - 本地虚拟机进程 ID
    * [interval] - 连续输出的时间间隔
    * [count] - 连续输出的次数
  * jstat -gc pid 最常用，可以评估程序内存使用及GC压力整体情况
        S0C：第一个幸存区的大小，单位KB 
        S1C：第二个幸存区的大小 
        S0U：第一个幸存区的使用大小
        S1U：第二个幸存区的使用大小 
        EC：伊甸园区的大小 
        EU：伊甸园区的使用大小 
        OC：老年代大小 
        OU：老年代使用大小 
        MC：方法区大小(元空间) 
        MU：方法区使用大小 
        CCSC:压缩类空间大小 
        CCSU:压缩类空间使用大小 
        YGC：年轻代垃圾回收次数 
        YGCT：年轻代垃圾回收消耗时间，单位s 
        FGC：老年代垃圾回收次数  
        FGCT：老年代垃圾回收消耗时间，单位s 
        GCT：垃圾回收消耗总时间，单位s 

### JVM可能出现的问题

* 频繁的进行full gc但是并没有出现oom？
  * 每次gc之后剩余的空间不大，说明一部分万股对象一直没法被回收，导致可用内存变小
  * 新生代的设置过小，频繁的从survivor0移动到survivor1，会导致快速达到一定年龄进入到老年代
* 新生代设置过小产生什么后果？
  * 一个是新生代GC次数非常频繁，增大系统消耗
  * 导致大对象直接进入老年代，占据了老年代的剩余空间，诱发了Full GC
* 新生代比例设置过大
  * 导致老年代过小，从而诱发full gc
  * 新生代的gc时间大幅度增加
* Survivor设置过小
  * 导致对象从eden直接到达老年代，降低了在新生代的存活时间
* Survivor设置过大
  * 导致eden过小，增加了GC频率

## GC

### 常用的gc算法分类

* 复制算法
  * 复制算法使用在新生区（因为新生区对象存活度比较低，不用来回复制）
  * 优点：速度快，没有内存碎片
  * 缺点：无法完全利用空间
* 标记清除算法
  * 优点：能够完全利用空间，不需要额外空间
  * 缺点：浪费时间（两次扫描），会产生内存碎片
  * 优化：标记清除压缩算法。向一段移动，目的是为了防止内存碎片

### jvm中的gc算法

也叫“分代收集算法”，综合使用了复制算法和标记清除算法

* 新生代

  * Serial
  * ParNew
  * parallel Scavenge

* 老年代

  * Serial Old

  * CMS

    步骤

    * 标记清除
    * 并发标记
    * 重新标记
    * 标记清除

* G1
  老年代和新生代都可以收集

  1. 基于标记清除算法，不产生内存碎片
  2. 精确控制停顿时间，在不牺牲吞吐量的前提下，实现低停顿回收垃圾

### minor gc/full gc的触发条件

* minor gc/full gc的触发条件、OOM的触发条件，降低GC的调优的策略。
    分析：列举一些我期望的回答：eden满了minor gc，升到老年代的对象大于老年代剩余空间full gc，或者小于时被HandlePromotionFailure参数强制full gc；gc与非gc时间耗时超过了GCTimeRatio的限制引发OOM，调优诸如通过NewRatio控制新生代老年代比例，通过MaxTenuringThreshold控制进入老年前生存的次数。

## JIT

### JIT即时编译器

特点：在运行期间对”热点代码“进行二次编译

在HotSpot虚拟机中，Java是通过解释器实现代码的运行的，但当某些代码执行较为频繁时，JVM就会认为这些代码为”热点代码“，而为了提高热点代码的执行效率，JVM会将这些热点代码编译为与本地平台相关的机器码，并进行各种层次的优化，而此时的操作就是通过即时编译器完成的



![](C:\Users\xyk\Desktop\knowledgeTree\java\images\运行原理.JPG)

[视频讲解](https://www.bilibili.com/video/BV1SZ4y1Z7Zp?spm_id_from=333.337.search-card.all.click)

## Spring

### SpringBoot面试题

#### SpringBoot的自动配置原理

* @SpringBootApplication注解是由
  	@SpringBootConfiguration
      **@EnableAutoConfiguration**
    	@ComponentScan三个注解组成，三个注解共同完成自动装配
* @SpringBootConfiguration注解标记启动类为配置类
* @ComponentScan注解实现启动时扫描启动类所在的包以及子包下所有标记为bean的类由IOC容器注册为bean
* EnableAutoConfiguration通过@Import注解导入AutoConfigrationImportSelector类，然后通过AutoConfigurationImportSelector类，然后通过类中selectImports方法去读取需要被自动装配的组件依赖下的spring.factories文件配置的组件的类全名，并按照一定的规则过滤掉不符合要求的组件的类全名，将剩余读取到的各个组件的类全名集合返回给IOC容器并将这些组件注册为bean

#### SpringBean的生命周期

![](C:\Users\xyk\Desktop\knowledgeTree\java\images\Bean的生命周期图解.jpg)

四个阶段：实例化-属性赋值-初始化-销毁

* 实例化Bean对象
  Spring中的对象可以分为两类

  * 用户自定义的对象person等
  * 容器需要进行使用的比如BeanPostProcessor、BeanFactoryPostProcessor、BeanNameAware

* 设置对象属性（依赖注入）
  Spring通过BeanDefinition找到对象依赖的其他对象，并将这些对象赋予当前对象

* 检查Aware相关接口并设置相关依赖

  * Spring会检测对象是否实现了xxxAware接口，如果实现了就会调用对应的方法用来设置这个Bean的属性值
  * 如果对象中需要引用容器内部的对象，那么需要调用aware接口的子类方法来进行统一的设置

* BeanPostProcessor的前置处理
  通过实现postProcessorBeforeInitialzation(Obect bean, String beanName)

  当前正在初始化的bean对象会被传递进来，我们就可以对这个bean作任何处理

* 判断当前bean对象是否设置了InitializingBean接口，然后调用afterPropertiesSet来进行属性的设置等工作

* 检查是否配置有自定义的init-method方法
  如果当前bean对象定义了初始化方法，那么在此处调用初始化方法

* BeanPostProcessor的后置处理

  通过实现postProcessorAfterInitialzation(Obect bean, String beanName)

  当前正在初始化的bean对象会被传递进来，我们就可以对这个bean作任何处理

  对生成的bean对象进行后置的处理工作（在这里可以完成AOP的相关功能）

* 注册必要的Destruction相关回调接口
  为了方便对象的销毁，在此处调用注销的回调接口，方便对象进行销毁操作

### Filter-过滤器

* 实现过程

  * 实现FilterRegistrationBean<Filter>接口

  * 实现Filter接口

    ```java
    @Order(20)
    @Component
    public class ApiFilterRegistrationBean extends FilterRegistrationBean<Filter> {
    
    	@PostConstruct
    	public void init() {
    		setFilter(new ApiFilter());
    		setUrlPatterns(List.of("/api/*"));
    	}
    
    	class ApiFilter implements Filter {
    		@Override
    		public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
    				throws IOException, ServletException {
    			HttpServletResponse resp = (HttpServletResponse) response;
    			resp.setHeader("X-Api-Version", "1.0");
    			chain.doFilter(request, response);
    		}
    	}
    }
    ```



### Interceptor-拦截器

* 实现过程

  * 编写拦截器实现类，实现HandlerInterceptro接口

    ```java
    public class MyInterceptor implements HandlerInterceptor {
        @Override
        public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
            // 统一拦截，判断是否有登录（输入有账号密码，userName就会存储到session）
            Object value = request.getSession().getAttribute("Lotus");
            if (value != null) {
                return true;
            }else {
                request.getRequestDispatcher("/WEB-INF/views/user/login.jsp").forward(request, response);
                return false;
            }
        }
    
        @Override
        public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView modelAndView) throws Exception {
    
        }
    
        @Override
        public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) throws Exception {
    
        }
    }
    
    
    ```

  * 编写拦截器配置类，实现WebMvcConfigurer接口

    ```java
    @Configuration //一定要加上这个注解，成为Springboot的配置类，不然不会生效
    public class WebMvcConfiguration implements WebMvcConfigurer {
     
        @Override   //拦截器配置 
        public void addInterceptors(InterceptorRegistry registry) {
            registry.addInterceptor(new MyInterceptor()) //拦截器注册对象
            .addPathPatterns("/**") //指定要拦截的请求
            .excludePathPatterns("/user/login"); //排除请求
    
        }
    }
    
    ```

### Spring AOP

* JDK 动态代理和 CGLIB 动态代理对比
  JDK 动态代理只能代理实现了接口的类，而 CGLIB 可以代理未实现任何接口的类。 另外CGLIB动态代理是通过生成一个被代理类的子类来拦截被代理类的方法调用，因此不能代理声明为final类型的类和final或static方法。
  就二者的效率来说，大部分情况都是JDK动态代理更优秀，随着 JDK 版本的升级，这个优势更加明显。
  CGLIB底层采用ASM字节码生成框架，使用字节码技术生成代理类，而JDK采用的是Java反射来生成的类
  Spring在选择用JDK还是CGLib的依据

* JDK和CGLIB的选择依据

  当Bean实现接口时，Spring就会用JDK的动态代理
  当Bean没有实现接口时，Spring使用CGLib来实现
  [原文链接](https://blog.csdn.net/MrLiar17/article/details/88869326)

  


# C/C++

## new和malloc

new不能完全取代malloc，因为C++有很多地方要调用c语言，而C语言都是用malloc和free来管理内存的

* malloc()
  https://zhuanlan.zhihu.com/p/452686042
* malloc_tirm()

## delete

delete[] 和free p（数组）如何知道数组的长度的？在申请这些内存的时候会在返回的指针的前面一点的位置预留一小段内存来存放数组的长度信息

## 编译

* c语言编译后的可执行文件的存储区域划分
  * 栈区（向下增长）
    由编译器自动分配和释放，存放函数参数，局部变量等
  * 堆区（向上增长）
    由程序员分配和释放
  * 数据区
    * rodata只读数据段（常量段），如字符串常量、全局const变量
    * .bss未初始化以及初始化为0的全局变量和静态局部变量
    * .data已经初始化读写数据段，如初始化为非0的全局变量和静态局部变量
  * 代码区.（text）
    存放的是即将要执行的代码

# golang

## 运算符优先级

| 运算符           | 优先级 |
| :--------------- | :----- |
| ^ !              | 7      |
| * / % << >> & &^ | 6      |
| + - \| ^         | 5      |
| == != < <= >= >  | 4      |
| <-               | 3      |
| &&               | 2      |
| \|\|             | 1      |

单目>运算>逻辑and>逻辑or

## 进制

八进制、十进制、十六进制

```
// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制
 
	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b)  // 77
 
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF
 
	// 二进制不能直接去表示

```

## 格式化输出

%T表示输出变量的类型

%v表示按照默认格式输出

%+v输出结构体的时候会增加字段名字

%b    表示为二进制 

%c    该值对应的unicode码值 

%d    表示为十进制 

%o    表示为八进制 

%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 

%x    表示为十六进制，使用a-f 

%X    表示为十六进制，使用A-F 

%U    表示为Unicode格式：U+1234，等价于"U+%04X"

## xorm

* 软删除

  ```Go
  type User struct {
      Id int64
      Name string
      DeletedAt time.Time `xorm:"deleted"`
  }
  ```

  在Delete()时，deleted标记的字段将会被自动更新为当前时间而不是去删除该条记录，如下所示：

  ```Go
  var user User
  engine.Id(1).Get(&user)
  // SELECT * FROM user WHERE id = ?
  engine.Id(1).Delete(&user)
  // UPDATE user SET ..., deleted_at = ? WHERE id = ?
  engine.Id(1).Get(&user)
  // 再次调用Get，此时将返回false, nil，即记录不存在
  engine.Id(1).Delete(&user)
  // 再次调用删除会返回0, nil，即记录不存在
  ```

* 硬删除

  ```Go
  engine.Id(1).Unscoped().Get(&user)//硬删除，普通的调用delete是软删除并非真正的删除
  ```

## GC

* [go ballast](https://cloud.tencent.com/developer/article/1903097?from=article.detail.1900650)超大数组实现达到特定内存值触法GC
  目前go 1.19已经通过SetMemoryLimit设置内存上限来解决了

## Doc

(1) https://www.bookstack.cn/read/qcrao-Go-Questions/README.md

# Git

## [更新git软件版本的方法](https://blog.csdn.net/Ezreal_King/article/details/79999131)

sudo add-apt-repository ppa:git-core/ppa

sudo apt-get update

sudo apt-get install git

## [游戏地址](https://learngitbranching.js.org/?locale=zh_CN)

通过实际操作理解，是最有效的实践方式

## 变量配置

Git具有三个配置文件：

1. 版本库级配置文件（.git/config）
配置仅仅针对当前项目有效；
若使用 git config 时用 --local选项，读写的就是这个文件。
2. 全局配置文件（~/.gitconfig ）
用户目录下的配置文件只适用于该用户；
若使用 git config 时用 --global选项，读写的就是这个文件。
3. 系统级配置文件（/etc/gitconfig ）
系统中对所有用户都普遍适用的配置；
若使用 git config 时用 --system选项，读写的就是这个文件。
注意----优先级

版本库级配置文件 > 全局配置文件 > 系统级配置文件
因此，每一个级别的配置都会覆盖上层的相同配置.
原文链接：https://blog.csdn.net/CXHPLY/article/details/50419951

* git config配置命令
  ```
  git config --global user.name="xuyongkang"
  git config --global user.email="xuyongkang"
  ```

* 查看git config的命令
  git config --list

## checkout

签出（真实作用是切换指向树的指针）

* 命令:切换到某个分支
  git checkout -b newBranch
  拷贝当前分支到一个新的分支并切过去
  
* 命令：将当前指针指向某个commit（将HEAD指针指向那个commitid）
  
  git checkout commitid
  例如：git checkout c6

## branch

分支管理

* 命令

  git branch

  列出本地所有分支
  git branch -d myBranch
  删除myBranch分支

  git branch -D myBranch
  强制删除myBranch分支
  
  git branch -f master HEAD~2
  强制将master分枝指向当前HEAD的[后退两个节点](https://learngitbranching.js.org/?locale=zh_CN)

## commit

* git commit --amend

  1. 用于将当前当前的更改追加到上一个commitid上，并且可以更改commit说明

  2. 用在rebase -i之后，可以交互式更改之前commit的commit说明

## stash

将当前的更改暂时存起来，这个非常的好用，场景就是当前自己正在做某一件事情，突然被要求去做另一件事情，但是当前的事情还没有做完，无法提交代码，因此这个命令可以将当前的更改保存起来，之后使用git stash pop弹出当前的更改就可以了。

当然，git stash 可以命名，同时git stash pop的时候可以弹出特定名称的更改



## remote

远程分支管理

* 命令

  git remote
  查看当前远程仓库

  git remote -v

  查看当前远程仓库的详细信息

  git remote add origin git://github.com/bakkdoor/grit.git
  （git   remote   add   远程仓库名   远程仓库url）

## fetch

拉取数据，更新（所有）本地的远程分支，但是不会更改本地的远程分支对应的分支

git fetch 之后 origin/master origin/dev 等等分支都会被更新到最行的状态，但是不会移动本地的master和dev分支

## merge

将两个分支，合并提交为一个新分支，并且新提交有2个parent

## pull **新奇**

git pull 就是git fetch && git merge的缩写两者的执行效果是一样的

**git pull --rebase 就是git fetch && git rebase的简写**

## rebase

会取消分支中每个提交，并把他们临时存放，然后把当前分治更新到最新的origin分支，最后再把所有的提交应用到分支上

* merge & rebase

  merge命令会保留所有的commit的历史时间。形成了以merge时间为基准的网状历史结构。每个分支上都会继续保留各自的代码记录，主分支上只保留merge的历史记录。

  简单来讲就是merge会保留次分支的历史记录，而rebase并不会
  
* 常用命令和作用

  * git rebase -i commit_id
    交互式的处理commit_id之前的分支
  * git rebase master
    将master上的新的commit加入到自己的分支上（非merge）
  

## reset

从当前commit想前会退

* 常用命令
  git reset HEAD^1
  撤销掉当前的提交**git reset对本地起作用但是对于远程的别人的分枝并不起作用**

* 参数

  * soft比mix模式更进一步，还将add记录了下来（也就是已经stage了之后直接commit就算完成了）
  * mix   使用混合模式，使用reset复位到了某个commit，之前的commit都被删除了，但是以往commit的文件都被保存到原来的位置并没有动，也就是可以再次提交
  * hard 如果使用hard，所有的commit也会消失，而且所有commit期间的文件也都消失了，所以之前历史的commit增加的文件无法再次提交了。
    **最常用以及默认的都是mix，这个模式会将你本地的commit撤回到add之前的状态，这样你就可以分拣文件再add**

  所以一个小总结：soft保存着你add的历史（也就是已经帮你stage过了），-mix还保存着你的文件，-hard连文件都给你删除了。

  注意：不管你什么时候reset，没有add和commit的文件会在切换分支和resetcommit的过程中一直保持文件空悬着

## revert

此命令专业用来去除掉代码不需要的commits

* 命令1：清除某一个commit
  git revert commit_id

* 命令2: 清除commitA到commitB的提交
  git revert -n commitA...commitB

  然后
  git commit -m "注释一下刚才revert"

* 常用命令

  * 第一步
    git revert -n 60d0ce86628734d789e6dcbcca050821a0cee469^...b5403832bdfe15f7511d70559295147b0263cf95

    revert掉60d0ce86628到b5403832bdf之间的所有commits（包括二者）

  * 第二步
    git commit -m "注释一下刚才的revert"

## [submodule](https://blog.csdn.net/qq_43382853/article/details/117401168)

如果当前模块中引入git文件，那么在上传或者下载的时候，并不会把代码真正的上传，可以使用这个命令来注解子模块的下载地址

```
zp@ZPdeMacBook-Pro game-baloot % git add .
warning: adding embedded git repository: src/github.com/pmezard/go-difflib
hint: You've added another git repository inside your current repository.
hint: Clones of the outer repository will not contain the contents of
hint: the embedded repository and will not know how to obtain it.
hint: If you meant to add a submodule, use:
hint: 
hint:   git submodule add <url> src/github.com/pmezard/go-difflib
hint: 
hint: If you added this path by mistake, you can remove it from the
hint: index with:
hint: 
hint:   git rm --cached src/github.com/pmezard/go-difflib
hint: 
hint: See "git help submodule" for more information.
warning: adding embedded git repository: src/gopkg.in/yaml.v3

```



## cherry-pick

拣取别的分支的commits到自己的分支上

* 常用命令:拣取一个commit

  on master branch: git cherry-pick commit_id
  
* 常用命令：拣取两个commitid之间的所有的commit
  git cherry-pick commitidA...commitidB
  
  注意取的commit是(commitidA, commitidB]

## HEAD^与HEAD~

HEAD~表示某个提交分支往上的第几个提交

HEAD^表示第几个父提交

也可以~^综合使用，从二可以在提交上游走

[典型示例](https://blog.csdn.net/fly_zxy/article/details/82593842)

## origin master与origin/master

* origin master 代表远程分支如

  git push origin master表示推送到远程分支

* origin/master代表本地分支
  git reset origin/master代表reset本地的master分支

## 强制拉新

```
git fetch --all
git reset --hard origin/master
git pull
```

## tag

* 加本地tag

​		git tag tagname -m "xxxxcommit"

* 删除本地tag

  git tag -d tagname

* 将tag推至远程服务器上

​		git push origin tagname

* 删除远程服务器的tag
  git push origin  :refs/tags/tag-name



# 操作系统

## 协程

子程序就是协程的一种特例，也被称为用户级线程

协程的切换在用户态完成，切换的代价比线程从用户态到内核态的代价小很多，相应的信息再协程栈中保存

协程对计算密集型的任务也没有太大的好处，计算密集型的任务本身不需要大量的线程切换，因此协程的作用也十分有限，反而还增加了协程切换的开销。

[图解说明](https://zhuanlan.zhihu.com/p/215231969)

# Linux

## ssh

1. 实现免密登陆
1.1 rsa 文件生成
如果你~/.ssh目录下面有 id_rsa 和 id_rsa.pub公钥、私钥两个文件，可以跳过此步骤

使用ssh-keygen命令来生成rsa秘钥文件到~/.ssh目录下

$ ssh-keygen -t rsa
你可以到~/.ssh目录下查看生成的id_rsa id_rsa.pub两个文件。

1.2 将公钥 rsa.pub 上传到服务器
使用ssh-copy-id命令将生成的公钥上传到服务器

$ ssh-copy-id -i ~/.ssh/id_rsa.pub username@server -p 22
1
ssh-copy-id命令需要提供你服务器的登陆方式和用户密码。
————————————————
版权声明：本文为CSDN博主「苏铎」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/u014568993/article/details/84308268

## vi

* 运行模式

  * i插入
  * o插入
  * a追加

* 查找

  * :/word向后查找
  * :?word向前查找

* 复制、粘贴和删除

  * **x**

    向后删除一个字符

  * **X**

    向前删除一个字符

  * **nx**
    连续向后删除n个字符

  * **dd**

    删除游标所在的那一整列

  * **ndd**

    n为数字。删除光标所在的向下n列

  * d1G
    删除光标所在到第一行的所有数据

  * dG

    删除光标所在到最后一行到所有数据

  * d$
    删除光标所在到改行的最后一个字符

  * d0

    删除光标所在到该行首字符的所有数据

  * 复制

    * yy
      复制游标所在那一行

    * nyy

      复制光标所在的向下n列，例如20yy表示复制20列

    * yG

      复制光标所在列到最后一列的所有数据

* 替换

  * **:n1,n2s/word1/word2/g**
    在第n1与n2行之间寻找word1这个字符串，并将该字符串取代为word2。
  * :**1,$s/word1/word2/g**
    从第一行到最后一行寻找word1字符串，并将该字符串取代为word2
  * **:1,$s/word1/word2/gc**
    从第一行到最后一行寻找word1字符串，并将该字符串取代为word2，且在取代前显示提示字符给用户确认（confirm）是否需要取代。（**非常好用**）

* 移动光标

  * **ctrl+f**
    向下翻一页
  * **ctrl+b**
    向上翻一页

  
  

  * gg
    移动到第一行
  * G
    移动到最后一行

  * H 
    光标移动到这个屏幕的最上方的那一行的第一个字符

  * M

    光标移动到这个屏幕的中央那一行的第一个字符

  * L

    光标移动到这个屏幕最下方的第一个字符

  

  

  * nG
    移动到第n行
  * n<space>行移动
    表示表表会向右/左移动n个字符
  * n<enter>列移动
    光标从当前位置向下移动n行

* 翻页

  * ctrl+f
    屏幕向下翻页
  * ctrl+b
    屏幕向上范爷



## awk

* 使用awk来检索大于某个大小的文件
  ls -Ral | awk '{if($5 > 100000){print $0 $5 $9}}'
* ls -alh | awk '{if($7 == 13 && $6 == "Aug" && $9 != "."){print $9}}' | xargs cat | grep "playerid eq"

## xargs

* ls -alh | awk '{if($7 == 13 && $6 == "Aug" && $9 != "."){print $9}}' | xargs cat | grep "playerid eq"
  筛选所有的8月13号更新的log中的带有“playerid eq”字符的行

## grep

Grep PLAYLOG stderr.log | grep "Room(225)" 二次过滤

## tail

tail -f stderr.log 动态更新显示stderr.log

## shell

* declare

  ```
  declare [-aixr] variable
  选项与参数：
  -a  ：将后面名为 variable 的变量定义成为数组 (array) 类型
  -i  ：将后面名为 variable 的变量定义成为整数数字 (integer) 类型
  -x  ：用法与 export 一样，就是将后面的 variable 变成环境变量；
  -r  ：将变量配置成为 readonly 类型，该变量不可被更改内容，也不能 unset
  ```

* 数组
  var[index] = content
  其中var是数组名

* 文件系统及程序的限制关系：ulimit

  我们的 bash 是可以『限制用户的某些系统资源』的，包括可以开启的文件数量， 可以使用的 CPU 时间，可以使用的内存总量等等

  示例：ulimit -f 10240

* **变量内容的删除、取代与替换**

  | 变量配置方式                                         | 说明                                                         |
  | ---------------------------------------------------- | :----------------------------------------------------------- |
  | ${变量#关键词} ${变量##关键词}                       | 若变量内容从头开始的数据符合『关键词』，则将符合的最短数据删除 若变量内容从头开始的数据符合『关键词』，则将符合的最长数据删除 |
  | ${变量%关键词} ${变量%%关键词}                       | 若变量内容从尾向前的数据符合『关键词』，则将符合的最短数据删除 若变量内容从尾向前的数据符合『关键词』，则将符合的最长数据删除 |
  | ${变量/旧字符串/新字符串} ${变量//旧字符串/新字符串} | 若变量内容符合『旧字符串』则『第一个旧字符串会被新字符串取代』 若变量内容符合『旧字符串』则『全部的旧字符串会被新字符串取代』 |

## 资源

### 文件结构

![目录树相关性示意图](http://cn.linux.vbird.org/linux_basic/0130designlinux_files/dirtree.gif)

### 文件资源

* ulimit -n 显示当前用户可以打开的最大的文件数量

## 系统配置

* 添加环境变量路径

  ```shell
  Vi ~/.bash_profile
  export PATH="$HOME/bin:$PATH"
  source ~/.bash_profile
  ```

## ssh-scp

* scp复制文件到远程机器上
  scp -P port file_name user@ip:/dir_name



## curl

* -O  将请求下载为文件

  curl -L -O https://github.com/buger/goreplay/releases/download/v1.3.1/gor_1.3.1_x64.tar.gz

* -d  发送POST请求

## supervisor

* 配置文件的地址
  程序运行在远程的linux机器上，使用supervisor进行管理管理的配置文件在对应的远程机器上，如test_xj上的/etc/supervisor/conf.d/文件夹下面
* 更新配置文件的步骤（增加或者删除conf文件）

​		编辑/etc/supervisor/conf.d/xxx.conf并保存修改

​		supervisorctl reread （读取新添加的配置让supervisor服务知道）

​		supervisorctl update （新添加的配置正式在supervisor中生效）

​		（一般不执行）supervisorctl restart xxx
​		链接：https://www.jianshu.com/p/1cdda9089506

* supervisorctl是supervisor的客户端
  * Supervisorctl status 显示所有正在运行的进程的状态

## 网络IO

### 网络通信编码

* 流程图

![](C:\Users\xyk\Desktop\knowledgeTree\计算机网络\images\socket通信流程.png)

* 示例代码

```c
//创建socket
int s = socket(AF_INET, SOCK_STREAM, 0);   
//绑定
bind(s, ...)
//监听
listen(s, ...)
//接受客户端连接
int c = accept(s, ...)
//接收客户端数据
recv(c, ...);
//将数据打印出来
printf(...)
```

* [listen函数讲解](https://blog.csdn.net/m0_46655373/article/details/122166674)

  ```c
  #include <sys/types.h>          /* See NOTES */
  #include <sys/socket.h>
  
  int listen(int sockfd, int backlog);
  1.sockfd:一个已绑定未被连接的套接字描述符
  2.backlog: 规定了内核应该为相应套接字排队的最大连接个数。用SOMAXCONN则为系统给出的最大值(linux系统中默认为128，一般都调到1024)
  
  ```

  

### IO多路复用机制

#### select

#### poll

#### epoll

* 原理

  [一片文章弄清除epoll的本质](https://zhuanlan.zhihu.com/p/64746509)

  **再来看看epoll**：
  epoll不会让每个 socke t的等待队列都添加进程A引用，而是在等待队列，添加 eventPoll对象的引用。
  当socket就绪时，中断程序会操作eventPoll，在eventPoll中的就绪列表(rdlist)，添加scoket引用。
  这样的话，进程A只需要不断循环遍历rdlist，从而获取就绪的socket。
  从代码来看每次执行到epoll_wait，其实都是去遍历 rdlist。

  如果rdlist为空，那么就阻塞进程。
  当有socket处于就绪状态，也是发中断信号，再调用对应的中断程序。
  此时中断程序，会把socket加到rdlist，然后唤醒进程。进程再去遍历rdlist，获取到就绪socket。

  **总之**： poll是翻译轮询的意思，我们可以看到poll和epoll都有轮询的过程。
  **不同点在于**：
  poll轮询的是所有的socket。
  而epoll只轮询就绪的socket。

  **自己看的**：用eventpoll代替了阻塞的线程，接管所有的通知，然后注册到rdlist，然后epoll_wait去循环遍历rdlist

* 水平触发和边缘触发

  * 水平触发
    只要文件描述符关联的读内核缓冲区非空，有数据可读，就一直发出刻度信号进行通知
    只要文件描述符关联的内核缓冲区不满，有空间可以写入，就一直发出可写信号进行通知

    如果系统中有大量的不需要读写的就绪文件描述符，而他们每次都会返回，这样会大大降低处理程序检索自己关心的就绪文件描述符的效率。

  * 边缘触发
    当文件描述符关联的读内核缓冲区由空转化为非空的时候，发出可读信号进行通知
    当文件描述符关联的写内核缓冲区由满转化为不满的时候，则发出可写信号进行通知

    如果这次没有吧数据全部读写完（如读写缓冲区太小）它不会再次通知你，直到该文件描述符上出现第二次可读写事件才会通知你，这种模式比水平触发效率高，系统不会充斥大量不关心的就绪文件描述符。

* 示例代码

  ```c
  int s = socket(AF_INET, SOCK_STREAM, 0);  
  bind(s, ...)
  listen(s, ...)
  
  int fds[] =  存放需要监听的socket
  
  while(1){
      int n = select(..., fds, ...)
      for(int i=0; i < fds.count; i++){
          if(FD_ISSET(fds[i], ...)){
              //fds[i]的数据处理
          }
      }
  }
  ```



## 调试工具

* 启动调试

  ```linux
  g++ -g test.cpp -o test
  gdb ./test
  ```

* gdb的调试命令

  ```
  start                   #开始调试,停在第一行代码处,(gdb)start
  l                         #list的缩写查看源代码,(gdb) l [number/function]
  b <lines>           #b: Breakpoint的简写，设置断点。(gdb) b 10
  b <func>            #b: Breakpoint的简写，设置断点。(gdb) b main
  b filename:[line/function]  #b:在文件filename的某行或某个函数处设置断点
  i breakpoints  #i:info 的简写。(gdb)i breakpoints
  d [bpNO]        #d: Delete breakpoint的简写，删除指定编号的某个断点，或删除所有断点。断点编号从1开始递增。 (gdb)d 1
  s                     #s: step执行一行源程序代码，如果此行代码中有函数调用，则进入该函数；(gdb) s
  n                     #n: next执行一行源程序代码，此行代码中的函数调用也一并执行。(gdb) n
  r                      #Run的简写，运行被调试的程序。如果此前没有下过断点，则执行完整个程序；如果有断点，则程序暂停在第一个可用断点处。(gdb) r
  c                       #Continue的简写，继续执行被调试程序，直至下一个断点或程序结束。(gdb) c
  finish                #函数结束
  p [var]              #Print的简写，显示指定变量（临时变量或全局变量 例如 int a）的值。(gdb) p a
  display [var]                #display，设置想要跟踪的变量(例如 int a)。(gdb) display a
  undisplay [varnum]     #undisplay取消对变量的跟踪，被跟踪变量用整型数标识。(gdb) undisplay 1
  set args               #可指定运行时参数。(gdb)set args 10 20
  show args           #查看运行时参数。
  q                          #Quit的简写，退出GDB调试环境。(gdb) q 
  help [cmd]           #GDB帮助命令，提供对GDB名种命令的解释说明。如果指定了“命令名称”参数，则显示该命令的详细说明；如果没有指定参数，则分类显示所有GDB命令，供用户进一步浏览和查询。(gdb)help
  回车                    #重复前面的命令，(gdb)回车
  
  ```


## 运维

* 运维工具
  * 远程运维工具 fabric
    https://fabric-chs.readthedocs.io/zh_CN/chs/tutorial.html

* 

  
  
  
  
  
  
  
  


# 问题

* 线程池中有空闲线程，正在工作线程个数也没有超过设定的工作线程的个数，问再来一个任务的时候是使用空闲线程还是新启动一个线程？
* treeMap的特点
  
  * 无序，唯一，可排序，底层使用红黑树
* treeMap的应用场景
  * 文件目录
  * 组织机构上下级
* 红黑树的结构特点
  
* 每次插入和删除最多只需要三次旋转就能达到平衡
  
* java中类的加载过程
  
* 类的加载是将类的.class文件中的二进制数据读入到内存中，将其放到运行时数据区的方法去中，然后再堆区创建一个java.lang.Class对象，用来封装类在方法区内的数据结构。类的加载的最终产品是位于堆区的Class对象，Class对象封装了类在方法区内的数据结构，并且向java程序员提供了访问方法去内的数据结构的接口。
  
* 多线程中notify()和notify()的区别
  notify()方法不能唤醒某个具体的线程，所以只有一个线程在等待的时候它才有用武之地。而notifyAll()唤醒所有线程并允许他们争夺锁，确保了至少有一个线程能继续运行。

* 为什么wait和notify方法要在同步块中调用？

  当一个线程需要调用对象的wait()方法的时候，这个线程必须拥有该对象的锁，接着它就会释放这个对象锁并进入等待状态直到其他线程调用这个对象上的notify()方法。同样的，当一个线程需要调用对象的notify()方法时，它会释放这个对象的锁，以便其他在等待的线程就可以得到这个对象锁。由于所有的这些方法都需要线程持有对象的锁，这样就只能通过同步来实现，所以他们只能在同步方法或者同步块中被调用。如果你不这么做，代码会抛出IllegalMonitorStateException异常。