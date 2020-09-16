week2访问属性
	* private
变量使用private进行修饰的时候，这个变量只可以在自己的类中被访问，在其他类中不可以被访问
	* public
对于类成员来说：意思是谁都可以访问，注意，构造函数本身也是public的
对于类来说：可以有public的类，但是这个public的类必须与源代码文件重名；
	* 默认（friend）（访问权限是包级别访问权限）
对于类成员来说：在同一个包中的类可以访问的到，不在同一个包中的无法访问的到
对于类来说：如果类的前面是默认的什么都没有的，那个这样的类只能在这个包中来使用，出了这个包就无法使用





父类成员访问属性		在父类中的含义	在子类中的含义
public				对所有人开放		对所有人开放
protected			只有包内其它类，自己和子类可以访问		只有包内其它类、自己和子类可以访问

					只有包内其它类可以访问					如果子类与父类在同一个包内：只有包内其它类可以访问否则：相当于private，不能访问

					只有自己可以访问						不能访问 


2.static
    使用static来修饰的成员称为 类的成员不属于任何一个对象；类是对象，函数是实体；static函数不能与任何实体对象建立联系，因此也就无法调用对象中的非static变量和函数

week3对象容器
1.顺序容器(List)
	* 
ArrayList
	* 
对象数组
当数组的元素的类型是类的时候，数组的每个元素其实只是对象的管理者而不是对象本身。因此，仅仅创建数组并没有创建其中的每个对象。


        List<String> ret = new List<String>();
        List是一个借口，所以不能实例化
        List<String> ret = new ArrayList<String>();正确
2.集合容器
	* 
HashSet
	* 
HashMap


week4继承与多态
1.关于多态
	* 
静态绑定：这种在编译之前就已经确定了
	* 
动态绑定：在运行的过程中才确定


2.关于向上造型（cast）
    对于基本变量来说：是类型转换
    int a = (int)10.2;
    这是将float类型直接改造成int类型
    而对于对象类型的变量来说：把一个对象当做另一个类型的对象
    item a = (item)CD是用item变量来管理一个子对象。


	* 总结正则表达式已经掌握；复习java语法；复习网络知识

刷leetcode中使用到的相关库的算法：
stringbuffer  -> string
math.abs();

java 7中list没有sort只有collection.sort();
java 8中的list是有sort的，但是都是要自定义comparator

Comparator c = new Comparator<Integer>(){
    @override
     public int compare(Integer o1, Integer o2){
        if((int)o1<(int)o2)
            return true;
        else
            return false;
    }
}


	* 
关于java中的深拷贝和浅拷贝问题（就简单的数组而言）
http://www.pianshen.com/article/3292318749/

java中的Stack数据结构
Stack<Integer> stack = new Stack<>();
stack中的常用操作：
stack.push();
stack.pop();
stack.peek();->返回顶部的元素
stack中的遍历方法：
for(Integer i:stack){
    print(i);
}
java中的queue队列
//Stack<Integer> stack = new Stack<Integer>();
//Queue<Integer> myQueue = new LinkedList<Integer>();
//LinkedList也实现了队列的接口
queue.offer(tmp.left);从队列尾部加入
queue.poll()->object;从队列头部取出

Integer 类
Integer类中有一个函数叫做parseInt(String);
Integer.valueOf(int);返回一个Integer类型的对象
.parseInt(string s);->int

String 类
String类中没有remove函数
.substring(1,5);返回一个子字符串[1,5)区间
.subSequence();返回一个char序列
String.valueOf(float)；将一个float类型的数据转换为字符串类型
.split(String regix)->String []
.replace('o','n')->String  {用法s = s.replace('0', 'n');//将字符串中的所有0替换成n}
.equals(string s)->boolean

ArrayList
ArrayList的长度统计方法是：a.size();


StringBuffer 相关用法和属性
StringBuffer buffer;
buffer.append(boolean/char/char[]/CharSequence/int/float/double/object);

buffer.delete(int start, int end)
buffer.deleteCharAt(int index);

buffer.insert(int offset, char/char[]/String);

buffer.substring(int start);
buffer.substring(int start, int end);

buffer.toString();

buffer.reverse();

HashSet
.add ->增加元素
for(:){}进行循环

HashMap
.get(object key)->value
.containsValue(object value)->boolean;
.containsKey(object key)->boolean;
.put(key,value)->返回之前对应的key或者null
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

LinkedList
linkedlist.add(index, element);将元素插入到特定的位置


