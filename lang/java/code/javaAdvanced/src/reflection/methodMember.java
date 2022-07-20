package reflection;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class methodMember {
	
	private static final boolean Method = false;

	public static void main(String [] s) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException {

		D obj = new D();
		Class c = obj.getClass();
		
		//��ȡpublic������������͸��ӿ�
		Method[] ms = c.getMethods();
		
		for(Method m:ms) {
			if("f1".equals(m.getName())) {
				m.invoke(obj, null);
			}
		}
		
		//��ø�������з���
		Method[] ms2 = c.getDeclaredMethods();
		for(Method m:ms2) {
			if("f2".equals(m.getName())) {
				m.setAccessible(true);
				String result = (String) m.invoke(obj, "abc");
				System.out.println(result);
			}
			
		}
	
	}
}
