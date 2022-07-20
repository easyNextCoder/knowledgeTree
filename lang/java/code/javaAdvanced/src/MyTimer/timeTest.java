package MyTimer;

import java.util.*;

class timeTest {
	
	public static void main(String []s) {
		Timer mt = new Timer("display");
		TimerTask task = new myTask();
		mt.schedule(task, 1000, 1000);
	}

}



class myTask extends TimerTask{
	int n = 0;
	public void run() {
		n++;
		System.out.print(new Date());
		System.out.println("---"+n);
	}
}
