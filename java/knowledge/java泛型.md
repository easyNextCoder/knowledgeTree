## java泛型
* 方法泛型
* 接口泛型
* 类泛型

## java多线程

##### 多线程的创建（最终都是重写Run（）函数）

* 实现Runnable接口

	    class MyTask implements Runnable{
	        MyTask(){}
	        public void run(){}
	    }
* 继承Thread类

		   class MyTask extends Thread{
		        MyTask(){}
		        public void run(){}
		   }
* lambda表达式

		new Thread(()->{
		    public void run(){
		        
		    } 
		}).start();
* 使用线程池

首先是线程池的使用：
    首先创建一个线程池
    ExecutorService pool = Executors.newCachedThreadPool();
    创建一个新的线程
    MyTask t1 = new MyTask(5);
    线程池开始执行任务
    pool.execute(t1);
    线程池关闭
    pool.shutdown();
    
    
##### 多线程的控制

.start()开始执行
.sleep()暂时停止线程的运行
try{Thread.sleep(1000);}catch{ InterruptedException e}


##### 多线程的同步

* 线程不同步带来的问题？
    比如开5000个线程同时让其对一个变量++，则最后得到的结果并不一定是5000，而是小于5000；原因是有可能两个线程同时去加，而后一个线程的结果覆盖了前一个线程的结果（变量缓存导致的结果），造成了结果小于5000

这个synchronized关键字，能够对对象进行同步，类似于lock与mutex配合使用的效果

* synchronized 用来与对象的互斥锁联系
    * 用法：
        * 对代码片段：
            * synchronized(对象){....}
        * 对某个方法：
            * synchronized放在方法声明中，
            * public synchronized void push(char c){...}; 
            * 相当于对synchronized(this),表示整个方法为同步方法
    * synchronized是独占锁
* .wait();
* .notify()//.notifyAll();

##### 死锁

* a运行的前提是与b同步，b运行的前提是与a同步

		class Worker
		{
		       int id;
		       public Worker(int id){ this.id=id; }
		       synchronized void doTaskWithCooperator(Worker other){
		              try{ Thread.sleep(500); } catch(Exception e){}
		              synchronized(other){
		              //synchronized(other)->执行这一句时，获取other对象，同时执行println
		              //如果别的线程已经占用则无法锁定，如果自己锁定则别人无法锁定
		                     System.out.println("doing" + id);
		              }
		       }
		}
		
		Worker w1 = new Worker(1);
		Worker w2 = new Worker(2);
		Thread td1 = new Thread(()->{
		        w1.doTaskWithCooperator(w2);
		        });
		Thread td2 = new Thread(()->{
		        w2.doTaskWithCooperator(w1);
		        });



##### 并发API
* 线程锁
    java.util.concurrent.lock
* 显式锁
* 原子变量
    * java.util.concurrent.atomic
    AtomicInteger cnt;
    cnt.getAndIncrement();
* 线程池
* 集合与线程
    * JDK1.5之前集合对线程安全的情况是：
    * ArrayList/HashMap不是线程安全的
    * vector及Hashtable是线程安全的
    * 产生一个相对线程安全的集合对象,但是即使这样之后这样也不方便
        * Collections.synchronizedArrayList(list)
    * JDK1.5之后**java.util.concurrent**包中增加了一些方便的类
        * CopyOnWriteArrayList
        * CopyOnWriteArraySet（以上适合于很少写入而读取频繁的对象）
        * ConcurrentHashMap
            * putIfAbsent(), remove(), replace()
        * ArrayBlockingQueue
            * 生产者与消费者，使用put()及take()
* Timer类
    * java.util.Timer
    每隔一定时间执行某段代码
###### 显式锁
###### 集合与线程


    
    
    //线程的创建
    package multiThread;
    import java.util.concurrent.*;
    
    public class myThreadPoolDemo {
           public static void main(String [] s) {
                  ExecutorService myPool = Executors.newCachedThreadPool();
                  oakTask task1 = new oakTask(1);
                  oakTask task2 = new oakTask(2);
                  oakTask task3 = new oakTask(3);
                  obkTask task4 = new obkTask(4)；
     
                  myPool.execute(task1);
                  myPool.execute(task2);
                  myPool.execute(task3);
                  myPool.execute(task4);
                  myPool.execute(new Thread(()-> {
                               for(int i = 0; i< 100; i++) {
                                      System.out.print("n");
                                  try {
                                             Thread.sleep(500);
                                      } catch (InterruptedException e) {
                                             // TODO Auto-generated catch block
                                             e.printStackTrace();
                                      }
       }
                  }));
                  myPool.shutdown();
           }
    }
    class obkTask extends Thread{
           int n = 0;
           obkTask(int n){
                  this.n = n;
           }
           public void run() {
                  for(int i = 0; i<100; i++) {
                         System.out.print(n);
                  }
           }
    }
    
    class oakTask implements Runnable{
           int n = 0;
           oakTask(int n){
                  this.n = n;
           }
           @Override
           public void run() {
                  for(int i = 0; i<100; i++) {
                         System.out.print(n);
                  }
           }
    }
    
    
    



##### java中线程的分类
* daemon线程->后台守护线程（如果主线程结束，则内部线程结束）
* 非daemon线程->非后台守护线程
使用.setDaemon(true)可以设置为后台线程

## java反射
* 概念：
    * 程序可以访问、检测和修改它本身状态或行为的能力，即自描述和自控制。
    * 可以在运行时加载、探知和使用编译期间完全未知的类。
    * 给java插上动态语言特性的翅膀，弥补强类型语言的不足。
* 作用：
    * 在运行中分析类的能力
    * 在运行中查看和操作对象
        * 基于反射自由创建对象
        * 反射构建出无法访问的成员变量
        * 调用不可访问的方法
    * 实现通用的数组操作代码
    * 类似函数指针的功能
* 创建对象的几种方法：
    * 直接使用new
    * 使用克隆
    * 序列化和反序列化
    * 反射



    		package reflection;
    		public class B implements Cloneable {
    		
    		       public void hello()
    		       {
    		              System.out.println("hello from B");
    		       }
    		       protected Object clone() throws CloneNotSupportedException
    		       {
    		              return super.clone();
    		       }
    		}


           ---
			package reflection;
			public class cloneB {
			
			       public static void main(String [] s) throws CloneNotSupportedException {
			              B obj2 = new B();
			              obj2.hello();
			              B obj3 = (B) obj2.clone();
			              obj3.hello();
			       }
			}


##### java反射中的关键类

######  关键类1：Class 类型标识

   * JVM为每个对象都保留其类型标识信息(Runtime Type Identification)三种获取方式：
        String s1 = "abc";
        Class c1 = s1.getClass();
        Class c2 = Class.forName("java.lang.String);
        Class c3 = String.class;
###### Field类
    
   * getFields()；
    获取本类及父类所有的public字段
   * getDeclareFields()
    获得本类所有声明的字段
    
###### Constructor构造方法
    
   * getConstructors()
   youd
    
###### Method成员方法
   
   * getMethods()
   * getDeclareMethods()

###### 父类(Class)
    
   * getSuperClass()

###### 修饰符 
    
   * getModifiers()

##### java反射的示例代码：


class对象NO使用反射生成A对象
```
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
		
		```
		
		public class A {
		       public void hello() {
		              System.out.println("hello from A");
		       }
		}
```

## java流

##### 函数式编程（强调做什么，而不是怎么样去做）
lambda实现了一个函数式编程，是一种全新的思考问题的方式




## String类
* 方法
    * .concat(String s)->String   ;连接两个String类


## HashSet容器
 * .add() 操作；map中是put操作

## Collections.sort()
示例代码：



	class Solution {
	    //这一题看的那个超级简单的题解，并且学习了java中的对象的排序算法
	    public String minNumber(int[] nums) {
	        List<String> list = new ArrayList<String>();
	        for(int i = 0; i<nums.length; i++){
	            list.add(Integer.valueOf(nums[i]).toString());
	        }  
	        Collections.sort(list, new Comparator<String>(){
	            public int compare(String a, String b) {
	                int lengthMax = a.length() + b.length();
	                String aPb = a.concat(b);//学习了，学习了
	                String bPa = b.concat(a);
	                for(int i = 0; i<lengthMax; i++) {
	                    if(aPb.charAt(i) == bPa.charAt(i)){
	                        continue; 
	                    }else{
	                        return aPb.charAt(i) - bPa.charAt(i);
	                    }
	                }
	                return -1;
	            }
	        });
	        StringBuffer buffer = new StringBuffer();
	        
	        for(int i = 0; i<list.size(); i++)
	            buffer.append(list.get(i));
	        return buffer.toString();
	    }
	}



## 关于重构和MVC设计模式
https://www.cnblogs.com/chanshuyi/p/some_experienc_in_system_refactor.html

## java重构
https://www.cnblogs.com/chanshuyi/p/some_experienc_in_system_refactor.html

[TOC]
