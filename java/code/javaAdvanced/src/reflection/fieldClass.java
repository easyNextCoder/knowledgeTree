package reflection;

import java.lang.reflect.Field;

public class fieldClass {
	public static void main(String [] s) throws IllegalArgumentException, IllegalAccessException {
		A obj = new A(20, "Tom");
		Class c = obj.getClass();
		
		//获取本类及父类所有的public字段
		Field [] fs = c.getFields();
		System.out.println(fs[0].getName() + ":" + fs[0].get(obj));
		
		//获得本类所有声明的字段
		Field[] fs2 = c.getDeclaredFields();
		for(Field f:fs2) {
			f.setAccessible(true);
			System.out.println(f.getName() + ":" + f.get(obj));
		}
	}
}
