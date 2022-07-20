package multiThread;
import java.lang.
public class multiThreadTest {
	public static void main(String [] s) {
		new Thread(()->{
			public void run() {
				for(int i = 0; i<1000; i++) {
					System.out.println("thread a.");
				}
			}
		}).start();	
	}
	
}
