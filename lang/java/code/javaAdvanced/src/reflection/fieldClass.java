package reflection;

import java.lang.reflect.Field;

public class fieldClass {
	public static void main(String [] s) throws IllegalArgumentException, IllegalAccessException {
		A obj = new A(20, "Tom");
		Class c = obj.getClass();
		
		//��ȡ���༰�������е�public�ֶ�
		Field [] fs = c.getFields();
		System.out.println(fs[0].getName() + ":" + fs[0].get(obj));
		
		//��ñ��������������ֶ�
		Field[] fs2 = c.getDeclaredFields();
		for(Field f:fs2) {
			f.setAccessible(true);
			System.out.println(f.getName() + ":" + f.get(obj));
		}
	}
}
