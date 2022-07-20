package week1;

import java.util.Scanner;



public class Main {



	public static void main(String[] args) {

		Scanner in = new Scanner(System.in);

		Fraction a = new Fraction(in.nextInt(), in.nextInt());

		Fraction b = new Fraction(in.nextInt(),in.nextInt());

		a.print();

		b.print();

		a.plus(b).print();

		//a.multiply(b).plus(new Fraction(5,6)).print();

		a.print();

		b.print();

		in.close();

	}


}

class Fraction{

	private int a;
	private int b;
	
	public Fraction(int nextInt, int nextInt2) {
		// TODO Auto-generated constructor stub
		a = nextInt;
		b = nextInt2;
	}

	public Fraction plus(Fraction b2) {
		// TODO Auto-generated method stub
		
		return null;
	}

	public void print() {
		// TODO Auto-generated method stub
		
	}
	
}