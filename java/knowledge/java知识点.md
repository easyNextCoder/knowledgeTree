## java中的内部类

### 访问属性

* private变量使用private进行修饰的时候，这个变量只可以在自己的类中被访问，在其他类中不可以被访问
* public对于类成员来说：意思是谁都可以访问，注意，构造函数本身也是public的对于类来说：可以有public的类，但是这个public的类必须与源代码文件重名；
* 默认（friend）对于类成员来说：在同一个包中的类可以访问的到，不在同一个包中的无法访问的到对于类来说：如果类的前面是默认的什么都没有的，那个这样的类只能在这个包中来使用，出了这个包就无法使用

父类成员访问属性在父类中的含义在子类中的含义public对所有人开放对所有人开放protected只有包内其它类，自己和子类可以访问只有包内其它类、自己和子类可以访问只有包内其它类可以访问如果子类与父类在同一个包内：只有包内其它类可以访问否则：相当于private，不能访问只有自己可以访问不能访问 2.static    使用static来修饰的成员称为 类的成员不属于任何一个对象；类是对象，函数是实体；static函数不能与任何实体对象建立联系，因此也就无法调用对象中的非static变量和函数

### 对象容器

1. 继承自Collection的有
	* ArrayList
		* a.size(); ArrayList的长度统计方法是
	* Stack
		* stack.push();
		* stack.pop();
		* stack.peek();->返回顶部的元素
			
				stack中的遍历方法：
				for(Integer i:stack){
				    print(i);
				}
	* HashSet
		* for(:){}进行循环
	* LinkedList(Queue)
		* queue.offer(tmp.left);从队列尾部加入
		* queue.poll()->object;从队列头部取出
		* linkedlist.add(index, element);将元素插入到特定的位置			

				//Queue<Integer> myQueue = new LinkedList<Integer>();
				//LinkedList也实现了队列的接口
2. 继承自map的容器
	1. HashMap
		1. .get(object key)->value
		2. .containsValue(object value)->boolean;
		3. .containsKey(object key)->boolean;
		4. .put(key,value)->返回之前对应的key或者null
		5. .getKey();
		6. .getValue();


				HashMap中如何进行遍历：

				Map<String, String> map = new HashMap<String,String>();
				for(Entry<String, String>entry:map.entrySet()){
				    System.out.println("键："+entry.getKey()+"值："+entry.getValue());
				}
				
				Iterator<Map.Entry<String, String>> iterator = map.entrySet().iterator();
				
				while(iterator.hasNext()){
				    Map.Entry<String,String>entry = iterator.next;
				    System.out.println("键："+entry.getKey()+"值："+entry.getValue());
				}

	1. LinkedHashMap（能够实现LRU缓存）
		1. .put();
		2. .get();
		3. Entry遍历，key,value

3. 注意：Collections.sort( ,comparator);

	java 7中list没有sort只有java 8中的list是有sort的，但是都是要自定义

		comparatorComparator c = new Comparator<Integer>(){    @override     public int compare(Integer o1, Integer o2){        if((int)o1<(int)o2)            return true;        else            return false;    }}

3. 注意：对象数组当数组的元素的类型是类的时候，数组的每个元素其实只是对象的管理者而不是对象本身。因此，仅仅创建数组并没有创建其中的每个对象。

    List<String> ret = new List<String>();        List是一个借口，所以不能实例化        List<String> ret = new ArrayList<String>();正确2.集合容器

## 基本对象的封装对象

* java中的基本类型占的字节数
	* boolean和byte是占一个字节的
	* char类型占两个字节
* java中的四类八种基本数据类型
	* 整数类型 byte short int long
	* 浮点型   float double
	* 逻辑型   boolean
	* 字符型   char

* Integer 类
 
	Integer类中有一个函数叫做parseInt(String);
	
	Integer.valueOf(int);返回一个Integer类型的对象
	
	.parseInt(string s);->int

其实当我们在为Integer赋值的时候，java编译器会将其翻译成调用valueOf()方法。比如Integer i=127翻译为Integer i=Integer.valueOf(127)
然后我们来看看valueOf()函数的源码：
		
		public static Integer valueOf(int i)
		    {
		        //high为127
		        if(i >= -128 && i <= IntegerCache.high)
		            return IntegerCache.***[i + 128];
		        else
		            return new Integer(i);
		    }

可以看出，对于-128到127之间的数，Java会对其进行缓存。而超出这个范围则新建一个对象。

> https://www.nowcoder.com/test/question/done?tid=32274197&qid=112819#summary

* String 类

	**String类中没有remove函数**
	.substring(1,5);返回一个子字符串[1,5)区间
	.subSequence();返回一个char序列
	String.valueOf(float)；将一个float类型的数据转换为字符串类型
	.split(String regix)->String []
	.replace('o','n')->String  {用法s = s.replace('0', 'n');//将字符串中的所有0替换成n}
	.equals(string s)->boolean


## 面向对象
* 抽象类（java中的抽象类和C++中的抽象类很相似）

	1、抽象类使用abstract修饰；2、抽象类不能实例化，即不能使用new关键字来实例化对象；3、含有抽象方法（使用abstract关键字修饰的方法）的类是抽象类，必须使用abstract关键字修饰；4、抽象类可以含有抽象方法，也可以不包含抽象方法，抽象类中可以有具体的方法；5、如果一个子类实现了父类（抽象类）的所有抽象方法，那么该子类可以不必是抽象类，否则就是抽象类；6、抽象类中的抽象方法只有方法体，没有具体实现；

* 接口

	1、接口使用interface修饰；2、接口不能被实例化；3、一个类只能继承一个类，但是可以实现多个接口；4、接口中方法均为抽象方法；5、接口中不能包含实例域或静态方法（静态方法必须实现，接口中方法是抽象方法，不能实现）

> https://baijiahao.baidu.com/s?id=1620965468190584914&wfr=spider&for=pc

### 继承与多态

1. 关于多态
	* 静态绑定：这种在编译之前就已经确定了
	* 动态绑定：在运行的过程中才确定

2. 关于向上造型（cast）    
	1. 对于基本变量来说：是类型转换    int a = (int)10.2;    这是将float类型直接改造成int类型    而对于对象类型的变量来说：把一个对象当做另一个类型的对象    item a = (item)CD是用item变量来管理一个子对象。
	
	* **总结正则表达式已经掌握；复习java语法；复习网络知识**

## 流

* StringBuffer 相关用法和属性
	
	StringBuffer buffer;
	buffer.append(boolean/char/char[]/CharSequence/int/float/double/object);
	
	buffer.delete(int start, int end)
	buffer.deleteCharAt(int index);
	
	buffer.insert(int offset, char/char[]/String);
	
	buffer.substring(int start);
	buffer.substring(int start, int end);
	
	buffer.toString();
	
	buffer.reverse();


## 刷leetcode中使用到的相关库的算法：stringbuffer  -> stringmath.abs();
	* 关于java中的深拷贝和浅拷贝问题（就简单的数组而言）http://www.pianshen.com/article/3292318749/


##  JDK和JRE的区别

* JDK(Java Development Kit)的组成：
    * 编译器等开发工具
    * JRE(Java Runtime Enviroment)
        * JVM
        * 运行类库
JVM是java字节码执行的引擎，还能优化字节码，使之转化成效率更高的机器指令。JVM中类的装载是由类加载器完成的。
JRE：包含了JVM和一些类库，如果不对代码进行编译只是运行，那么只需要jre就行
JDK：包含了JRE和一些开发工具，如javac命令可以对java代码进行编译

## ArrayList和LinkedList的区别

ArrayList底层实现是数组，实现了随机的访问，但是插入和删除费时。
LinkedList底层是用链表实现的，增删元素快，但是查询速度比较慢


## HashMap的内部实现，如何防止冲突？
* HashMap的底层实现：
	底层使用哈希映射的方法，链地址法，也就是数组+链表+红黑树的方式，来解决冲突问题；当链表的长度超过8的时候，就会将链表转化成红黑树，实现快速的增删。

* HashMap的性能：
	HashMap中的链表出现越少，性能才会越好。
* HashMap何时进行resize()
	当发生哈希冲突并且size大于阈值的时候，需要进行数组扩容，扩容时，需要新建一个长度为之前数组2倍的新的数组，然后将当前的Entry数组中的元素全部传输过去，扩容后的新数组长度为之前的2倍，所以扩容相对来说是个耗资源的操作。
* JDK1.8中HashMap如何防止rehash
	因此，我们在扩充HashMap的时候，不需要像JDK1.7的实现那样重新计算hash，只需要看看原来的hash值新增的那个bit是1还是0就好了，是0的话索引没变，是1的话索引变成“原索引+oldCap”。（oldCap就是原数组的长度，因为最终的索引的计算是h&（length-1）而length-1这个数除了最高一位其余位都是1）

https://blog.csdn.net/woshimaxiao1/article/details/83661464
https://www.cnblogs.com/look-look/p/11715439.html

## LinkedHashMap

* LinkedHashMap是HashMap的子类，但是内部还有一个双向链表维护键值对的顺序，每个减值对即位于哈希表中，也位于双向链表中。LinkedHashMap支持两种顺序插入顺序，访问顺序。
	* 插入顺序：先添加在最前面，后添加的在后面。修改操作不影响顺序
	* 访问顺序：所谓访问指的是get/put操作，对一个键执行get/put操作后，其对应的键值对会移动到链表末尾，所以最末尾的是最近访问的，最开始的是最久没有被访问的，这就是访问顺序。

* 使用LinkedHashMap可以用来实现LRU(Least Recently Used 最近最少使用）
	
	* 算法思想是：算法根据数据的历史记录来进行淘汰数据，其核心思想是:**如果数据最近被访问过，那么将来被访问的几率也更高**

	* 算法过程是：

		1. 新数据插入到链表头部；
		2. 每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
		3. 当链表满的时候，将链表尾部的数据丢弃。
	* 代码示例：

			package LRUCache;
			
			import java.util.LinkedHashMap;
			import java.util.Map;
			
			class LRUCaches<K, V> extends LinkedHashMap<K, V>{
				private int maxEntries;
				
				public LRUCaches(int maxEntries) {
					super(16, 0.75f, true);
					this.maxEntries = maxEntries;
				}
				
				@Override
				protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
					return size() > maxEntries;
				}
			}
			
			public class LRUCacheTest {
				public static void main(String []s) {
					LRUCaches<String, Object>cache = new LRUCaches<>(3);
					cache.put("a","abstract");
			        cache.put("b","basic");
			        cache.put("c","call");
			        cache.get("a");
			        cache.put("d","滴滴滴");
			        System.out.println(cache);
				}
			}


## TreeMap

	TreeMap可以实现存储元素的自动排序。在TreeMap中，键值对之间按键有序，TreeMap的实现基础是平衡二叉树。

	但是为了实现其正确的排序必须生成传入比较器。

	
## Spring

	spring能够快速的实现控制翻转IOC（控制反转）,和AOP编程（面向切面的编程）

> https://www.cnblogs.com/wmyskxz/p/8820371.html


## JDBC和Hibernate

* JDBC
	
	Java数据库连接，（Java Database Connectivity，简称JDBC）是Java语言中用来规范客户端程序如何来访问数据库的应用程序接口，提供了诸如查询和更新数据库中数据的方法。
	
> https://www.runoob.com/w3cnote/jdbc-use-guide.html

* JDBC使用示例：(一共是5个步骤)

		import java.sql.Connection;
		 2 import java.sql.DriverManager;
		 3 import java.sql.ResultSet;
		 4 import java.sql.SQLException;
		 5 import java.sql.Statement;
		 6 
		 7 public class connectionMysql {
		 8 
		 9     public static void main(String[] args) {
		10         
		11         String driver="com.mysql.jdbc.Driver";//驱动路径
		12         String url="jdbc:mysql://localhost:3306/eshop";//数据库地址
		13         String user="root";//访问数据库的用户名
		14         String password="123456";//用户密码        
		15         try {
		16             //1、加载驱动
		17             Class.forName(driver);
		18             //2、链接数据库
		19             Connection con = DriverManager.getConnection(url, user, password);
		20             if(!con.isClosed()){//判断数据库是否链接成功
		21                 System.out.println("已成功链接数据库！");
		22                 //3、创建Statement对象
		23                 Statement st = con.createStatement();
		24                 //4、执行sql语句
		25                 String sql="select *from user";//查询user表的所有信息
		26                 ResultSet rs = st.executeQuery(sql);//查询之后返回结果集
		27                 //5、打印出结果
		28                 while(rs.next()){
		29                        System.out.println(rs.getString("Id")+"\t"+rs.getString("name")+"\t"+rs.getString("password"));
		　　　　　　　　　　　}
		               }
		31                 rs.close();//关闭资源
		32                 con.close();//关闭数据库
		33             }
		34                 
		35             } catch (Exception e) {
		36                 // TODO Auto-generated catch block
		37                 e.printStackTrace();
		38             }
		39     }
		40 }

* Hibernate

	Hibernate是一个开放源代码的对象关系映射框架，它对JDBC进行了非常轻量级的对象封装，它将POJO与数据库表建立映射关系，是一个全自动的orm框架，**hibernate可以自动生成SQL语句，自动执行，使得Java程序员可以随心所欲的使用对象编程思维来操纵数据库**。 Hibernate可以应用在任何使用JDBC的场合，既可以在Java的客户端程序使用，也可以在Servlet/JSP的Web应用中使用，最具革命意义的是，Hibernate可以在应用EJB的JaveEE架构中取代CMP，完成数据持久化的重任.(来自百度百科)

> https://www.cnblogs.com/mq0036/p/8522150.html

	**自己看的理解的就是：能将繁杂的sql语句编程直接对对象的操作**

* 用法
	* session.get()
	* session.load()

	load方法来得到一个对象时，此时hibernate会使用延迟加载的机制来加载这个对象，即：当 我们使用session.load()方法来加载一个对象时，此时并不会发出sql语句，当前得到的这个对象其实是一个***对象，这个***对象只保存了实 体对象的id值，只有当我们要使用这个对象，得到其它属性时，这个时候才会发出sql语句，从数据库中去查询我们的对象。
	相对于load的延迟加载方式，get就直接的多，当我们使用session.get()方法来得到一个对象时，不管我们使不使用这个对象，此时都会发出sql语句去从数据库中查询出来。

## JVM垃圾回收(GC)原理

### 基本的垃圾回收算法
1. 引用计数
	* 比较古老的回收算法。原理是此对象有一个引用即增加一个计数，删除一个引用则减少一个计数。垃圾回收时，只用收集计数为0的对象。此算法最致命的是无法处理循环引用问题。

2. 分代(Generational Collection)
	* 基于对对象生命周期分析后得到的垃圾回收算法。把对象分为年轻代，年老代，持久代，对不同生命周期对象使用不同的算法（上述方法的一种）进行回收。现在的垃圾回收器都是使用此算法的。
		* Young年轻代：
		
			年轻代分三个区。一个Eden区，两个 Survivor区。大部分对象在Eden区中生成。当Eden区满时，还存活的对象将被复制到Survivor区（两个中的一个），当这个 Survivor区满时，此区的存活对象将被复制到另外一个Survivor区，当这个Survivor去也满了的时候，从第一个Survivor区复制过来的并且此时还存活的对象，将被复制“年老区(Tenured)”。需要注意，Survivor的两个区是对称的，没先后关系，所以同一个区中可能同时存在从Eden复制过来对象，和从前一个Survivor复制过来的对象，而复制到年老区的只有从第一个Survivor去过来的对象。而且，Survivor区总有一个是空的。 

		* Tenured年老代:

			年老代存放从年轻代存活的对象。一般来说年代存放的都是生命期教长的对象。

		* Perm持久代

			用于存放静态文件，如java类，方法等。持久代对垃圾回收没有显著影响，但是有些应用可能动态生成或者调用一些class，例如Hibernate等，在这种时候需要设置一个比较大的持久空间来存放这些运行过程中新增的类。持久代大小通过-XX:MaxPermSize=<N>进行设置。 

## 二、垃圾回收器 


* 目前的收集器主要有三种：串行收集器、并行收集器、并发收集器。

4. 小结 
* 串行处理器： 
--适用情况：数据量比较小（100M左右）；单处理器下并且对响应时间无要求的应用。 
--缺点：只能用于小型应用 
* 并行处理器： 
--适用情况：“对吞吐量有高要求”，多CPU、对应用响应时间无要求的中、大型应用。举例：后台处理、科学计算。 
--缺点：应用响应时间可能较长 
* 并发处理器： 
--适用情况：“对响应时间有高要求”，多CPU、对应用响应时间有较高要求的中、大型应用。举例：Web服务器/应用服务器、电信交换、集成开发环境。 


## java中的六原则一法则

* 接口隔离原则
* 合成聚合复用原则
* 迪米特法则-最小知识原则

## 常见配置举例 



> https://www.iteye.com/blog/chenchendefeng-455883


## java中的关键字

final可以修饰类，方法，变量
final 用于声明属性，方法和类，分别表示属性不可变，方法不可覆盖，类不可继承。
finally是异常处理语句结构的一部分，表示总是执行。
finalize是Object类的一个方法，在垃圾收集器执行的时候会调用被回收对象的此方法，可以覆盖此方法提供垃圾收集时的其他资源
回收，例如关闭文件等。

## LRU->redis

## Spring 快速网页应用


创建对象和对象的管理都是由Spring来完成，主要用来控制组件，降低企业的开发难度。

* 主要能够实现的功能：
	* IoC
		* Inversion of Control
		* 控制反转：对象创建对象管理由程序员管理变为由Spring管理
	* DI
		* Dependency Injection
		* 依赖注入：对象和对象之间的依赖关系的管理
	* AOP
		* Aspect Oriented Programming
		* 面向切面的编程
	

* Spring的组成
	* 核心模块（Core Container）
		* Bean（Bean工厂：创建对象的工厂）
		* Core
		* Context（上下文：Spring的容器）
		* SpEL（Spring的表达式语言）
	* 数据访问模块(Data Access/Integration)
		* 支持JDBC封装
		* ORM
		* OXM（对象和XML之间转换的一个支持）
		* JMS(生产者和消费者，消息功能的实现)
		* Transactions(事务管理)
	* Web(面向Web应用程序，比如Web中的一个上传功能)
		* WebSocket
		* Servlet(监听器，实现MVC的功能->实现Spring MVC)
		* Web
		* Portlet
	* AOP
	* Aspects(面向切面编程的一个重要组成部分)
	* Instrumentation
	* Messaging
	* Test(使用Spring提供的单元测试模块中方便的进行集成测试)


* Spring的使用
	* 使用注解符来使用Spring
		* @Component 是注册为组件 ->  XML <bean
			* @primary ->作为首选bean
			* @Qualifier("normal")->使用限定符
		* @ComponentScan是组件扫描类，将所有的组件扫描过来用
		* @autoWired 控制类之间的归属关系->XML <properity(可以放在构造函数之前或者set函数之前实现依赖注入，也可以放在内部的对象之前通过反射的方法实现依赖注入效率较低。)
	* 使用注解类的时候初始化Spring容器
	
		//初始化Spring容器
        ApplicationContext context = new AnnotationConfigApplicationContext(ApplicationSpring.class);
		运行完这段代码之后，@Component注解的类都被新生成了。

	* 使用XML文件来使用Spring

		<bean id = "service" class = "hello.MessageService"></bean>
		<bean id = "printer" class = "hello.MessagePrinter">
			<property name = "service" ref = "service"></property>
		</bean>

	* 关于自动装配
		* 组件扫描
			* @Component:表示这个类需在应用程序中被创建
			* @ComponentScan:自动发现应用程序中创建的类
		* 自动装配
			* @Autowired：自动满足bean之间的依赖
		* 定义配置类
			* @Configuration：表示当前类是一个配置类（能将componentScan与main函数之间解耦）
	* 使用单元测试
		* 引入Spring单元测试模块
			* maven:junit,spring-test
			* @RunWith(SpringJUnit4ClassRunner.class)
		* 加载配置类
			* @ContextConfiguration(classes=AppConfig.class)
	
	* @Autowire的四种不同的使用方法
		* 用在构造函数上
			* 多个依赖的情况
		* 用在成员变量上
		* 用在setter方法上
		* 用在任意方法上
	* @Component(required = false)//表明这个组件并不是必须要的，有可以没有也可以

	* 处理自动伞装配的歧义性 
		* 首选bean
			* 在声明类的时候使用@Primary
			* 只能定义一个@Primary
		* 使用限定符
			* 在声明的时候和装配的时候分别使用@Qualifier
		* 使用限定符和bean id
			* 在声明的时候指定bean的id（默认的id是首写字母小写的类名)
			* 在装配的时候使用@Qualifier

> https://www.cnblogs.com/wmyskxz/p/8820371.html

## IntelliJ IEDA

* 快捷键
	* alt+insert插入各种对应的construct和insert方法
	* ctrl+o ->插入无参的构造函数