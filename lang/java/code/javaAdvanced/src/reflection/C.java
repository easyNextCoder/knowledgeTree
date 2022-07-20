package reflection;

import java.io.Serializable;

public class C implements Serializable{
	private static final long serialVersionUID = 1L;
	
	public void hello()
	{
		System.out.println("hello from C");
	}
}
