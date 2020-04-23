package reflection;

public class B implements Cloneable {
	public void hello()
	{
		System.out.println("hello from B");
	}
	protected Object clone() throws CloneNotSupportedException
	{
		return super.clone();
	}
}
