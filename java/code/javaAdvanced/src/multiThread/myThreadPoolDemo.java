package multiThread;

import java.util.concurrent.*;
public class myThreadPoolDemo {
	public static void main(String [] s) {
		ExecutorService myPool = Executors.newCachedThreadPool();
		oakTask task1 = new oakTask(1);
		oakTask task2 = new oakTask(2);
		oakTask task3 = new oakTask(3);
		obkTask task4 = new obkTask(4);
		
		myPool.execute(task1);
		myPool.execute(task2);
		myPool.execute(task3);
		myPool.execute(task4);
		myPool.execute(new Thread(()-> {
				for(int i = 0; i< 100; i++) {
					try {
						Thread.sleep(500);
					} catch (InterruptedException e) {
						// TODO Auto-generated catch block
						e.printStackTrace();
					}
					System.out.print("n");
				}
			
		}));
		
		myPool.shutdown();
	}
}

class obkTask extends Thread{
	int n = 0;
	obkTask(int n){
		this.n = n;
	}
	public void run() {
		for(int i = 0; i<100; i++) {
			System.out.print(n);
		}
	}
}

class oakTask implements Runnable{
	int n = 0;
	oakTask(int n){
		this.n = n;
	}
	@Override
	public void run() {
		for(int i = 0; i<100; i++) {
			System.out.print(n);
		}
	}
}
