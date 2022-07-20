package reflection;

public class cloneB {
	public static void main(String [] s) throws CloneNotSupportedException {
		B obj2 = new B();
		obj2.hello();
		
		B obj3 = (B) obj2.clone();
		obj3.hello();
	}
}
