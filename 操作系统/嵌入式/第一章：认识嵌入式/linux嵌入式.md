认识嵌入式系统

### 嵌入式的基本组成

* bootloader
* 内核
* 驱动程序
* 根文件系统

### 嵌入式微处理器
* arm公司CPU指令集，CPU架构，CPU系列
  

![armcpu](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\arm-cpu.JPG)

## 体验linux系统

### linux文件系统

* 文件系统的类型
  * ext2(linux标准文件系统，专为linux开发)
  * swap
  * vfat（linux文件系统）
  * NFS（网络文件系统）

linux的文件系统采用阶层式的树状目录机构，在该机构中的最上层是根目录：/，然后在根目录下再建立其他的目录。

linux中无论操作系统管理几个磁盘分区，这样的目录树只有一个。从结构上讲，各个磁盘分区上的树形目录不一定是并列的。

/

​	/bin

​	/boot

​	/home

​	/opt

​	/usr

​	/etc

​	/var

​	/...

* linux文件类型

  * 普通文件：通常是流式文件

    * ls 显示出来是-
    * 分类
      * 纯文本文件
      * 二进制文件
      * 数据文件

  * 目录文件：用于表示和管理系统的全部文件

    * ls显示信息是：d

  * 设备文件：包括块设备文件和字符设备文件
    通常是放在/dev目录下

    * 分类

      * 块设备文件：
        * 用于存储数据提供随机存储的设备，常用的就是硬盘U盘等
        * ls显示的是b
      * 字符设备文件：
        * 是一些串行接口设备，比如：键盘串口等

      

  * 管道FIFO文件：提供进程间通信的一种方式

    * 提供进程间通信的一种方式
    * 文件属性的第一个字符是：p

  * 链接文件：用于不同目录下文件的共享

    * ls显示出来是：l
    * 这个文件指向另一个文件

  * 套接字文件：该文件类型与网络通信有关

    * 作用：该文件类型与网络通信有关
    * ls显示的是s

  ### 分区

  磁盘分区是使用分区编辑器在磁盘上划分几个逻辑部分，盘片一旦划分成数个分区，不同类的目录与文件可以存储进不同的分区。

  在传统的磁盘管理中，将一个硬盘划分为两大类分区：**主分区和拓展分区**。主分区是能够安装操作系统，能够进行计算机启动的分区，这样的分区可以直接格式化，然后安装系统，直接存放文件。

  拓展分区是除主分区以外的分区，不能够直接使用，必须将拓展分区分为若干个逻辑分区才可以使用，拓展分区可以没有，至多一个。逻辑分区在数量上没有什么限制。

  ### 挂载

  linux下的分区需要挂载到目录后才能使用，挂载的意义就是把磁盘分区的内容放到某个目录下。这个**把分区和目录对应的过程叫做挂载(Mount)**而这个挂载在文件树中的位置就是挂载点。

  当要使用某个设备时，例如要读取硬盘中的一个格式化好的分区、光盘或软件等设备时，必须先把这些设备对应到某个目录上，而这个目录就称为"挂载点（mount point)",这样才可以读取这些设备，将物理分区细节屏蔽掉，用户只有统一的逻辑概念，所有的东西都是文件。

  

  ## shell使用技巧

  略

  ### 环境变量

  * 环境变量的划分（按照环境变量的声明周期来划分）
  * 一类是永久的，需要修改配置文件，变量永久生效
      * 修改用户主目录下的.profile或.bashrc文件
    * 修改系统目录下的profile文件（**这个要非常慎重，尤其是通过root用户修改的环境变量，如果修改错了会发生非常严重的错误**）
    * 另一类是临时的，使用export声明即可，变量在关闭shell的时候就失效了
    * 命令演示：
        * export PATH=$PATH:PAHT 1:PATH 2
      * export MYNAME="my name is jack"
  * 常用的环境变量

  ​	PATH  	UID  	MAIL 	HOME 	SHELL	HISTSIZE	USER	TERM	HOSTNAME	LOGNAME	PWD

  * 常用的环境变量的命令
  * env命令：显示所有的环境变量
    * set命令：set命令显示本地定义的shell变量
    * 使用unset命令来清除环境变量；-f仅删除函数，-v仅删除变量
  
  ### 管理linux文件
  
  
  
  ### 内容管理
  
  * cat
    * cat file1>>file2     将file1追加到file2上
  * more
    * 这个命令与cat类似，不过它可以以分页的方式来显示
  * less
    * 与more命令非常相似，可以随意浏览文件
  * diff
    * 比较文件的差异
    * 显示的内容
      * c：是change的意思
      * d：是delete的意思
      * a：是add的意思
    * 命令使用示例：
      * diff log1 log2 -y -w 50(这个命令是将两个文件并列成两排并显示两个文件之间的差别)
    * 比较两个文件并生成补丁文件
      * diff -ruN  log1  log2 > log.patch
        * -u  表示统一格式
        * -r  表示比较目录
        * -N 表示将不存在的文件当作空文件来进行处理
  * patch
    * 用途：让用户利用设置修补文件的方式，修改，更新原始文件
    * 格式：patch [options]  [originalfile]  [patchfile]
    * 使用示例：patch -p0  cp.c   cp.patch   使用cp.patch对cp.c打补丁
  * grep
    * egrep执行效果与grep带参数-E相似
    * 参数
      * -n  表示显示行数
      * -c  表示count以下的意思
      * -r 递归搜索目录下的子目录
    * grep正则表达式元字符集
      * .  匹配任意单个字符
      * [] 匹配指定范围内的任意单个字符
      * [^]匹配指定范围外的任意单个字符
      * 标准的字符类名称如下：
        * [:alnum:] 字母数字字符
        * [:alpha:]字母字符
        * [:digit:]数字：0 1 2 3 4 5 6 7 8 9
        * [:lower:] 匹配所有的小写字母
        * [:upper:]匹配所有的大写字母
        * [:black:]空字符：空格符，制表符
        * [:space:]空格字符：制表符，换行符，垂直制表符，换页符，回车符和空格键符
      * 位置锚定
        * ^ 行首锚定：用于模式的最左侧
        * $行尾锚定，用于模式的最右侧
        * ^PATTREN$:用PATTERN来匹配整行
        * ^$用来匹配空白行
        * ^[[:space:]]*$:空行或包含空白字符的行
        * \word  锚定单词word为开始的行
      * 匹配前面的字符m到n次
        * c\\{m, n\\}
    * 命令演示
      * ls -l | grep '^d'    返回所有的目录文件
  
  ### 用户与权限管理
  
  * linux中用户一共有三种：超级用户，系统用户，普通用户
  * linux中文件有三种可操作的权限设置：ugo 文件的拥有者，相同用户组成员，其他成员
  * 用户与用户组管理
    * 用户账号的添加、删除与修改
      * 用户账号文件全部被存放在-passwd种
      * 用户的密码信息-/etc/shadow（只有root用户拥有权限，其他用户没有任何权限）
      * 用户组配置文件：在group文件
      * useradd 选项  用户名
        * -r  指定系统用户
        * -d  filePath     指定用户的home文件夹
      * passwd  
        * 用户设置和更改用户名密码，使用useradd创建一个用户后，必须设置一个密码才算完成
        * 命令示例：
          * passwd  testuser(回车之后，就可以设置或者更改testuser的密码了)
      * groupadd
        * 使用示例：
          * usermod -a -G testgroup  testuser(将testuser添加到testgroup种)
    * 用户口令的管理
    * 用户组的管理

### 查看目录和文件的属性

-rwxrwxrwx    文件所有者  文件所属组

* chmod命令
  * 用途：修改文件或目录的权限
  * 格式：chmod  [option] mod  file
  * 简单使用：
    * 数字更改法：
      * chmod xxx filename(xxx表示数字)
    * 文字更改法：
      * chmod  a-x hello  对hello文件的所有使用者，去除可执行属性
      * chmod ug+w , o-x log1  同时修改不同用户权限，文件所有者与文件所属组增加写权限，其他用户删除可执行权限
      * chmod a-x  log1  删除所有用户的可执权限
      * chmod u=x log1  使用=设置权限，将文件log1的所属用户的权限全部取消，并重设为只拥有可执行权限
* chown
  * 用途：修改文件和目录的所有者和所属组
  * 格式：chown  [option]  [owner] file...
  * 使用简单示例：
    * chown [-R]  所有者  文件或目录       -R表示递归修改文件目录
    * chown [-R] 所有者:所属组   文件或目录        （所有者：所属组     的格式是将文件的所有者设为所有者，同时将所属组设置为所属组）

### 归档及压缩命令

* 各种打包和压缩类型的分类
  * .Z		compress程序压缩的文件
  * .gz       gzip程序压缩的文件
  * .bz2     bzip2程序压缩的文件
  * .tar      程序打包的数据，并没有经过压缩
  * .tar.gz          tar程序打包的文件，其中并且经过gzip的压缩
  * .tar.bz2        tar程序打包的文件，其中并且经过bzip2的压缩
* gzip命令的使用示例：（gzip   -> gnu  zip）
  * gzip  filename        对文件进行压缩
  * gzip  -d  filename    对文件进行解压缩
* tar
  * 命令作用：归档及压缩命令
  * 用途：制作归档文件，释放归档文件
  * 格式：
    * tar  选项...     归档文件名    源文件或目录
    * tar  选项...     归档文件名     【-C 目标目录】
  * 常见命令选项
    * -c：创建.tar格式的包文件
    * -x：解开.tar格式的包文件
    * -z：调用gzip程序进行压缩或解压
    * -j:  调用bzip2程序进行压缩或解压
    * -v：输出详细信息
    * -f：表示使用归档文件
    * -p:打包时保留原始文件及目录的权限
    * -t:列表查看包内的文件
    * -C:解包时指定释放的目标文件夹

## 管理磁盘和文件系统

磁盘分区按照功能性的不同可以分为主分区(Primary)，拓展分区(Extended)及逻辑分区(Logical)三种

* 硬盘最多可以分割成4个主分区或三个主分区+1个拓展分区
* 拓展分区又可以分为数个（没有限制，但总容量不得超过拓展分区大小)逻辑分区



* fdisk
  * 用来管理和增加磁盘和对磁盘进行分区
* mkfs
  * 用途：make Filesystem，创建文件系统（格式化）
  * 格式：mkfs -t 文件系统类型分区设备
* mount
* umount



### 网络管理

* 查看网络连接情况
* netstat
  * -a 列出所有端口
  * -at列出所有的tcp连接
  * -apu列出所有的UDP连接
  * -i显示接口信息
  * -r显示路由信息
  * -n显示所有已经建立连接的
* ping
  * 测试网络的连通性
  * 格式：ping  【选项】 目标主机
* traceroute
  * 测试从当前主机到目的主机之间经过的网络结点
  * 格式：traceroute  目标主机地址
* ifconfig
  * 查看所有网络接口信息  ->ifconfig
  * 查看指定网络接口信息-> ifconfig  网络接口名
  * 禁用或者激活网卡
    * ifconfig up
    * ifconfig down  例子：ifconfig  eth0 down
  * 修改网卡的硬件地址
    * 首先先禁用掉网卡
      * ifconfig eth0 down
    * 然后修改硬件的硬件地址
      * ifconfig eth0 hw ether 00:01:02:03:04:05
  * 设置虚拟网络接口（同一块网卡要设置多个IP地址)
    * 格式：ifconfig 接口名：序号  IP地址
* route命令：设置路由记录
  * 要上网需要设置默认的网关，两个电脑之间通过网关才能通信
  * 删除路由表中的默认网关记录
    * 格式：route del default gw  IP地址
  * 向路由表中添加默认网关记录
    * 格式：route add default gw IP地址
  * 添加到指定网段的路由记录
    * 格式: route add -net 网段地址
  * 临时配置-使用命令调整网络参数
    * 简单、快速，可直接修改运行中的网络参数，一般只适合在调试网络的过程中使用，系统重启以后，所做的修改将会失效。
  * 固定设置-通过配置文件修改网络参数
    * 修改各项网络参数的配置文件，适合对服务器设置固定参数时使用，需要重载网络服务或者重启以后才会生效
* 网络接口配置文件
  * /etc/sysconfig/network-scripts/ 目录下的
    * ifcfg-eth0:第1块以太网卡的配置文件
    * ifcfg-eth1:第2块以太网卡的配置文件
    * ...
* 主机名称配置文件
  * /etc/sysconfig/network文件
    * 用途保存全局网络设置，主要包括主机名信息
* 域名解析配置文件
  * //etc/resolv.conf文件
    * 用途：保存本机需要使用的DNS服务器的IP地址

### 进程管理

引导流程总览：

![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\linux开机启动的步骤.JPG)

rc.d  -> rc意思是run-command的意思，d是文件夹的意思；意思就是运行需要开机自启动的命令 

* ps
  * 当前正在运行的进程
  * STAT：显示当前的进程的状态，linux中一共显示5中状态
    * 1.R 显示正在运行
    * 2.S中断，进程处于休眠中，或者受阻的状态
    * 3.D不可中断，是指受到信号不被唤醒，或者不可运行
    * 4.Z僵死进程
    * 5T停止
  * 命令使用
    * -aux 更加详细的显示进程的状态
* top
  * 实时的查看动态进程的排名信息
* Ctril+z
  * 中断正在运行的程序
* kill      killall
  * killall  进程名
  * kill   PID号

### shell

* 学习路线：shell脚本语言也是一种**语言**,所以可以按照语言来学习

* 分类：
  * 变量
  * 运算符
  * 流程控制
  * 函数
  * 内置函数

* 变量

  * 普通变量

    my_var = 'new.ruoduan.cn'

    my_var=5

* 数组

  arr_name=(1 2 3)

  arr_name=(string1 string2 string3)

  echo ${arr_name[0]}

* 运算符
  **表达式两旁加上``才能表示内部的表达式是按照命令来执行**

  * 算术运算符
  * 关系运算符
  * 布尔运算符
  * 字符串运算符
  * 文件测试运算符

* 流程控制

  * for var in ...
  * if  then  elif then elif fi
  * while(($count<=6)) do command done

* 函数

  ```shell
  function demoFun()
  {
  	echo "this is my first shell function."
  	return let "1+1"
  	#或者return `expr 1+1`
  }
  
  demoFun
  echo "demoFun result is:$?"
  
  ```

* shell 参数

  * 在.sh文件中

    ```
    $0   代表执行的sh command命令
    $1   代表输入sh  command的第一个参数
    $2
    $3...
    ```

* shell中的输入输出重定向

  * \>表示重定向

  * \>\>表示以追加的方式进行重定向

  * 如果希望将 stdout 和 stderr 合并后重定向到 file，可以这样写：

    ```
    $command >file 2>&1
    或者
    $command >>file 2>&1
    ```

  * 如果希望对 stdin 和 stdout 都重定向，可以这样写：

    ```
    $command <file1 >file2
    ```

  * 在命令行中通过命令计算Here Document的行数

    ```
    $wc -l <<H
    >hello
    >world
    >H(这个H代表终止符)
    $2
    ```

* shell文件包含

  ```
  . filename   # 注意点号(.)和文件名中间有一空格
  
  或
  
  source filename
  ```

### vim

10yy拷贝从当前光标开始往下的10行

dd删除当前行



## 嵌入式开发的常用工具

### makefile

* makefile里主要包含五种类型的语句

  显式规则、隐式规则、变量定义、文件指示、注释

  * 显式规则
    
    显式规则说明了如何生成一个或多个目标文件。这是由makefile的书写者明显指出，要生成的文件，文件的依赖文件，生成的命令
    
    * 通配符
      * 【*】匹配0个或者任意个字符
      * 【?】匹配任意一个字符
      * 【[]】指定匹配的字符放在"[]"中
      * 【%】匹配任意个字符，使用在targets、prerequisites中
    
  * 隐式规则
    
    * 由于make有自动推导的功能，所以隐式的规则可以让程序员比较简略的书写makefile，这是由make所支持，例如makefile发现.o文件，程序就会自动去寻找.c文件，并编译成.o文件
    
  * 变量定义
    
    * 在makefile中可以定义一系列的变量，变量一般都是字符串，当makefile被执行时，其中的变量都会被拓展到响应的位置上
    
  * 文件指示
    
    * 包括在一个makefile中引用另一个makefile,根据某些情况指定执行makefile中的有效部分、定义一个多行的命令。
    
  * 注释：
    
    * makefile注释使用#,若makefile需要用到#,则需要进行转义

* 命令格式：<target>:<depend>

​					command

command之前是tab键而不是空格符号



* makefile中，如果有部分源文件更新，make程序会只能判断出那些需要重新执行

* 命令示例：

  ```makefile
  #sample makefile script  #表示注释
  CC=gcc
  SRCS=fun1.c fun2.c main.c
  EXEC=test
  all:
  	$(CC) $(SRCS) -o $(EXEC)
  clean:
  	rm -rf $(EXEC)
  ```

  ```makefile
  #sample makefile script  #这个文件makefile能自动检查更新
  CC=gcc
  OBJS=fun1.o  fun2.o  main.o
  EXEC=test
  
  all:$(OBJS)
  	$(CC) $(OBJS) -o $(EXEC)
  fun1.o:fun1.c
  	$(CC) -c fun1.c
  fun2.o:fun2.c
  	$(CC) -c fun2.c
  main.o:main.c
  	$(CC) -c main.c
  clean:
  	rm -rf $(EXEC)
  ```

  ```makefile
  #
  # Makefile for nps ingate exec, by denghuibin
  #
  # *** TODO: (1) change 'test' to whatever you want for output file-name.
  EXECUTABLE = he
  
  
  CXX = g++
  CXXFLAGS  =  -W -g -O2 -DDEBUG
  COMPILE = $(CXX) $(CXXFLAGS)
  
  # *** TODO: (2) add paths of include files.
  INC  := -I./include -I./include/net-snmp 
  
  # *** TODO: (3) add paths of source files.
  SRCS := $(wildcard *.cpp) $(wildcard ./src/*.cpp)
  OBJS := $(patsubst %.cpp,%.o, $(SRCS))
  #DEPS := $(patsubst %.o,%.d, $(OBJS))
  
  # *** TODO: (4) add library to project.
  LIBS := -lrt ./lib/libnetsnmp.a
  
  #
  # Output file
  #
  all: $(EXECUTABLE)
  $(EXECUTABLE) : $(OBJS)
  	$(CXX) -o $(EXECUTABLE) $(OBJS) $(LIBS)
  
  #
  # specify the dependency
  #
  %.d : %.c
  	$(CXX) $(INC) -MM $< > $@
  	$(CXX) $(INC) -MM $< -MQ $@  >> $@
  #
  # specify that all .o files depend on .cpp files, and indicate how to realize it
  #
  %.o : %.cpp
  	$(COMPILE) -c $(INC) -o $@ $<
  
  .PHONY: clean
  clean:
  	-rm $(EXECUTABLE)
  	-rm $(OBJS)
  	rm -f *.bak *.o
  #	-rm $(DEPS)
  #depend : $(DEPS)
  #	@echo "Dependencies are now up-to-date."
  
  #-include $(DEPS)
  
  ```

  http://www.ruanyifeng.com/blog/2015/02/make.html

  https://www.icourse163.org/learn/JSIT-1001754045?tid=1450235453#/learn/content?type=detail&id=1214366465&cid=1218053059

## 嵌入式linux c编程

### 串口

* 串行接口简称串口，也称串行通信接口（通常指COM接口），是采用串行通信方式的拓展接口

  * RS-232-C:也称标准串口，是目前最常用的一种串行通信接口

  * RS-422：以差动方式发送和接收，不需要地线，差动工作是同等条件下，传输距离远的根本原因，这也是和RS-232之间的根本区别。RS-422和RS-485的传输距离都能超过1KM

  * RS-485采用平衡发送和差分接收因此具有抑制共模干扰的能力，RS-485采用半双工的工作方式，任何时候只能有一点处于发送状态，因此发送的信号需要由使能信号加以控制。RS-485用于多点互联时非常的方便，可以省掉很多的信号线

  * USB：通用串行总线，用于规范电脑与外部设备的连接和通讯，是应用在PC领域的接口技术。

  * RJ-45：RJ-45接口，适用于有双绞线构建的网络，以太网集线器都会提供这种端口

    ![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\串口.JPG)



一般情况下，使用两个脚RX,TX交叉相连其他引脚对应着连接起来就好了。



* 串口文件

  * 在linux中，针对所有的周边设备都提供了设备文件供用户访问，所以如果要访问串口，只要打开相关的设备文件即可。
    * 在Linux下串口文件是位于/dev下的
      * COM1串口1为/dev/ttyS0
      * COM2串口2为/dev/ttyS1

* 在使用串口之前必须设置相关配置，包括：波特率、数据位、检校位、停止位等

  * 涉及的头文件：<termios.h>

  * 一个简单的读写程序：

    ```c
    #include <unistd.h>
    #include <strings.h>
    #include <sys/types.h>
    #include <sys/stat.h>
    #include <stdlib.h>
    #include <stdio.h>  
    #include <fcntl.h>
    #include <termios.h>
    int main(void)
    {
    int fd;
    struct termios new_cfg,old_cfg;
    int speed;
    char buff[24];
    fd=open("/dev/ttyUSB0", O_RDWR|O_NOCTTY|O_NDELAY);
    if(fd < 0)
    {
    	perror("open serial port");
    	return(-1);
    }
    
    if (fcntl(fd, F_SETFL, 0) < 0)
    {
    	perror("fcntl F_SETFL\n");
    }
    if (tcgetattr(fd, &old_cfg) != 0)
    {
    	perror("tcgetattr");
    	return -1;
    }
    new_cfg = old_cfg;
    cfmakeraw(&new_cfg); 
    new_cfg.c_cflag &= ~CSIZE;
    speed = B115200;
    cfsetispeed(&new_cfg, speed);
    cfsetospeed(&new_cfg, speed);
    new_cfg.c_cflag |= CS8;
    new_cfg.c_cflag &= ~PARENB;
    new_cfg.c_iflag &= ~INPCK;
    new_cfg.c_cflag &= ~CSTOPB;
    new_cfg.c_cc[VTIME] = 0;
    new_cfg.c_cc[VMIN] = 0;
    tcflush(fd, TCIFLUSH);
    if ((tcsetattr(fd, TCSANOW, &new_cfg)) != 0)
    {
    	perror("tcsetattr");
    	return -1;
    }
    
    do
    {
    	memset(buff, 0, 24);
    	if (read(fd,buff,24) > 0)
    	{
    		printf("The received words are : %s \n", buff);
    	}
    } while(strncmp(buff, "quit", 4));
    close(fd);
    return 0;
    }
    ```

    

  

# 命令积累

* fdisk -l

  显示这台linux上的硬盘分区情况

* 输出重定向
    * 命令>>文件1        2>>文件2
    * 将错误输出到文件2中
    * 命令  &>>文件
    * 无论命令成功与否都把它保存到文件中
    * 命令 >>文件       2>&1
    
* 环境变量相关的命令
    * source  ：读入环境配置文件的命令
    * 使用格式：source 配置文件名
    * 环境变量更改后，在用户下次登陆时生效。如果想立刻生效，则可以执行下面的语句
    * source .bash_profile  或者 source  ~/.bashrc
    
  ### 文件管理命令
  
  pwd:Print Working Directory
  
  cd
  
  mkdir  -p  test2/test22     -p命令表示，如果test2不存在就重新创建
  
  ​	mkdir -m 777 test3     带权限的创建命令
  
  cp
  
  rm
  
  touch命令也可以改变文件的权限和时间戳
  
  * ln
    * 硬链接具有相同的节点号，不可以链接目录，修改任意一个都会发生改变
    * -s 软连接的确是一个新的文件 ，只是以路径的形式存在，删除源文件，只是删除了数据不会删除链接，软链接可以跨文件系统，软链接可以对目录进行链接（**相当于windows中的快捷方式，是指向文件或者目录的间接指针**）
  
  
  
  * 命令文件查找-which
  * 文件名查找-whereis locate find
    * whereis 找到与指定名字匹配的是二进制文件、源文件、和帮助手册文件所在的路径（也是通过数据库来查找）
    * locate查找的是所有与名字匹配的文件
    * find使用find进行查找时不要设置过大的搜索范围，否则会给电脑带来很大的压力
      * 常用的命令选项
        * -name:按文件名称查找
        * -size:按文件大小查找  +-  b(bit)c(char)w(word)k(KB)M(MB)G(GB)
        * -time:按照修改时间搜索
        * -perm:按文件权限搜索
          * -perm 权限模式：查找文件权限刚好等于“权限模式的文件”
          * -perm -权限模式：查找文件权限全部包含“权限模式”的文件
          * -perm +权限模式：查找文件权限包含“权限模式”的任意一个权限的文件
        * -type:按文件类型查找
          * d:查找目录
          * f：查找普通文件
          * l：查找软链接文件
          * c：字符设备
          * b：块设备
          * p：管道文件
          * -：普通文件
    * (**一般情况下：在进行查找的时候，首先使用which   locate然后再使用find**, locate是利用数据库进行查找，所以速度更快；find命令虽然更加强大但是是从硬盘上直接查找的，速度比较慢；但是locate查询的数据库是一天才更新一次的，所以如果有个文件是刚刚创建的，使用这个命令可能会找不到)
    * updatedb  这是一个单独命令，用来更新数据库
  

## BootLoader

常用的有uBoot, 还有三星的vivi（提供内核的配置界面）

## Linux内核

* linux内核的组成

  linux内核一共分为5个部分

  * 进程调度
  * 内存管理
  * 虚拟文件系统
  * 网络接口
  * 进程间通信

* 进程调度

  控制进程对CPU的访问。当需要选择下一个进程运行时，由调度程序选择最值得运行的进程。Linux使用了比较简单的基于优先级的进程调度算法选择新的进程。

  ![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\内核的运行状态.JPG)

* 内存管理

  Linux采用了分页的内存管理机制。

  Linux允许多个进程安全地共享主内存区域、支持虚拟内存

  Linux内存管理机制可以分为3个层次，从上而下依次为**物理内存的管理、页表的管理、虚拟内存的管理**

  ![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\内存管理.JPG)

* 虚拟文件系统
  虚拟文件系统隐藏了各种硬件的具体细节，为所有的设备提供了统一的接口

  虚拟文件系统可以分为**逻辑文件系统**和**设备驱动程序**。

  **逻辑文件系统**指Linux所支持的文件系统

  **设备驱动程序**指为每一种硬件控制器所编写的设备驱动程序模块

  ![image-20200531134506763](C:\Users\xyk\AppData\Roaming\Typora\typora-user-images\image-20200531134506763.png)

  * 字符设备

    是指在I/O传输过程中以字符为单位进行传输的设备，例如键盘、打印机等。在UNIX系统中，字符设备以特别文件方式在文件目录树中占据位置并拥有相应的结点。

    字符设备可以使用与普通文件相同的文件操作命令对字符设备文件进行操作，例如打开、关闭、读、鞋凳。

    当一台字符型设备在硬件上与主机相连之后，必须为这台设备创建字符特别文件。操作系统的mknod命令被用来建立设备特别文件。例如为一台终端创建名为/dev/tty03的命令如下（设主设备号为2，次设备号为13，字符型类型标记c）

  * 单个文件系统

    对于普通文件，就要通过文件系统；**文件系统把对文件的操作转换成对块设备的操作**。

    为什么要分块设备与字符设备？以前用字符设备是通过open, read, write等函数来操作底层的，然而在块设备中不能这么做，NAND flash的最小操作单位是块，一个块内有多个扇区。

    为了追求高效，需要把对块设备的操作以一定算法进行读写分类，排序，然后再进行实际的NAND FLASH读写操作。

    向NAND FLASH层提供硬件设备的相关存储器以及某些相关操作函数。

* 网络接口

  网络接口提供了对各种网络标准的存取和各种网络硬件的支持。

  网络接口可以分为**网络协议**和**网络驱动程序**。

  网络协议部分负责实现每一种可能的网络传输协议。

  网络设备驱动程序负责与硬件设备通讯，每一种可能的硬件设备都有相应的设备驱动程序。

  ![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\网络接口.JPG)

* 进程间通信

  进程调度子系统处于中心位置，所有其他的子系统都依赖它，因为每个子系统都需要挂起或恢复进程。

  * Linux中的几种通信机制
    * 管道
    * 命名管道fifo
    * 信号
    * 信号量
    * 消息队列
    * 共享内存

  ![image-20200531143750048](C:\Users\xyk\AppData\Roaming\Typora\typora-user-images\image-20200531143750048.png)

* 系统各个模块之间的关系
  ![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\系统模块之间的关系.JPG)

## Linux文件系统结构

### 文件系统：

文件系统是一套实现了**数据的存储、分级组织、访问和获取**等操作的抽象数据类型，一种存储和组织计算机文件和数据的方法，它使得对其访问和查找变得容易。

Linux最早的文件系统是**Minix**。

专门为Linux设计的文件系统-拓展文件系统第二版(EXT2最成功的文件系统)。

### Linux的虚拟文件系统

* 内核在它的底层文件系统接口上建立了一个抽象层，使得linux能够支持各种文件系统

  Linux支持**ext, ext2**,xia,minix, umsdos, msdos,msdes, fat32, ntfs,jfss,yaffs,romfs,cramfs,nfs以及ufs等多种文件系统

* VFS提供了一个通用文件系统模型，该模型囊括了我们所能想到的文件系统的常用功能和行为
* 定义了所有文件系统都支持的基本抽象接口和数据结构

### Linux文件系统结构

* 用户层

* 内核层

* 驱动层

* 硬件层

  

![](C:\Users\xyk\Desktop\knowledgeTree\操作系统\嵌入式\第一章：认识嵌入式\Linux文件系统结构.JPG)

