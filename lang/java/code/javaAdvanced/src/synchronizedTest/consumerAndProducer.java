package synchronizedTest;

abstract class  things<T>{
	abstract void put(T value);
	abstract T get();
}

class liquidThings extends things<Integer>{

	private Integer []  tube = null;
	private int size = 0;
	private int nowAt = 0;
	liquidThings(int n){
		this.size = n;
		tube = new Integer[size];
	}
	
	@Override
	synchronized void put(Integer value) {
		// TODO Auto-generated method stub
		if(nowAt<size-1) {
			tube[++nowAt] = value;
			this.notify();
		}else {
			try {
				this.wait();
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
	}

	@Override
	synchronized Integer get() {
		// TODO Auto-generated method stub
		if(nowAt >= 0) {
			Integer value = tube[nowAt--];
			this.notify();
			return value;
		}else {
			try {
				this.wait();
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
		return 0;
	}

	
	
}

class xproducer extends Thread{
	
	private things thing;
	xproducer(things t){
		this.thing = t;
	}
	public  static  int cnt = 0;
	public void product(Integer value) {
		
		thing.put(value);
		/*
		for(int i = 0; i<5000; i++) {
			System.out.println(this.cnt);
			synchronized(this){(new Thread(()-> {this.cnt++;System.out.println(this.cnt);})).start();};; 
		}
		*/	
	}
	public void run() {
		for(int i = 0; i<100; i++) {
			this.product(i);
		}
	}
}

class xconsumer extends Thread{
	private things thing;
	xconsumer(things t){
		this.thing = t;
	}
	public Integer xconsume() {
		return (Integer) thing.get();
	}
	public void run() {
		for(int i = 0; i<100; i++) {
			System.out.print(this.xconsume());
		}
	}
}

public class consumerAndProducer {
	
	public static void main(String []s) {
		liquidThings t = new liquidThings(5);
		xproducer proWork1 = new xproducer(t);
		xconsumer conWork1 = new xconsumer(t);
		proWork1.start();
		conWork1.start();
		//System.out.println(proWork1.cnt);
	}
}


