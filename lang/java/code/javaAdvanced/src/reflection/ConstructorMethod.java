package reflection;
import java.lang.reflect.Constructor;
public class Constructor {
	public static void main(String [] s) {
		D d = new D(null);
		
		Class c = d.getClass();
		
		Constructor cons = c.getConstructor();
		for(Constructor con:cons) {
			if(con.getParameterCount() > 0) {
				D obj = (D) con.newInstance("100");
				obj.f2();
			}else {
				;
			}
		}
	}
}
