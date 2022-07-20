# java包
## java.math

	* BigInteger 大数操作


		包中的常用方法和变量：
		初始化：BigInteger a = new BigInteger("100");
		方法：a.mod();
		          a.add(b);
		          a.intValue();  ->输出最终的int型数据。
		import java.math.*;
		class Solution {
		    public int fib(int n) {
		        if(n == 0)return 0;
		        if(n == 1)return 1;
		        BigInteger  [] arr = new  BigInteger[n+1];
		        arr[0] = arr[0].valueOf(0);
		        arr[1] = arr[1].valueOf(1);
		        for(int i = 2; i<=n; i++)
		            arr[i] = arr[i-1].add(arr[i-2]);
		        BigInteger toMod = new BigInteger("1000000007");
		        
		        return  arr[n].mod(toMod).intValue();
		    }
		}

## java.io.Serializable
* Seriallizable反射接口

		package reflection;
		import java.io.Serializable;
		public class C implements Serializable{
		       private static final long serialVersionUID = 1L;
		       
		       public void hello()
		       {
		              System.out.println("hello from C");
		       }
		}


## java.lang.reflect

.Field;类型变量
.Constructor

## java.util.concurrent
实现并发的api包
.lock
.atomic

## java.util.Timer
.Timer