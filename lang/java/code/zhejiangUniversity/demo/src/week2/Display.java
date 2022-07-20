package week2;

public class Display {
	
	private int value = 0;
	private int limit = 0;
	
	Display(int limit){
		this.limit = limit;
	}
	
	public void increase() {
		value++;
		if(value>=limit) {
			value = 0;
		}
	}
	
	public int getValue() {
		return value;
	}
	
	
	@Override
	public String toString() {
		return "Display [value=" + value + ", limit=" + limit + "]";
	}

	public static void main(String []args) {
		Display display = new Display(24);
		for(;;) {
			display.increase();
			System.out.println(display);
		}
		
	}
}
