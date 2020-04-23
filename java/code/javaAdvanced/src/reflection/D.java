package reflection;

public class D {
	private String is = null;
	D(String is){
		this.is = is;
	}
	public void f1() {
		System.out.println("B.f1()...");
	}
	
	private String f2(String s) {
		System.out.println("B.f2()...");
		System.out.printf("is:%s", is);
		return s;
	}
}
