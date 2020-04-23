package reflection;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;


public class NO {
	public static void main(String []s) throws InstantiationException, IllegalAccessException, ClassNotFoundException, NoSuchMethodException, SecurityException, IllegalArgumentException, InvocationTargetException {
		Object obj6 = Class.forName("reflection.A").newInstance();
		Method m = Class.forName("reflection.A").getMethod("hello");
		m.invoke(obj6);
		//A obj = new A();
		//obj.hello();
		//((A) obj6).hello();
	}
}


